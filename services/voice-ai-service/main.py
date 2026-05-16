from datetime import datetime
import base64
import json
import logging
import os
import tempfile
import threading
import time

import openai
import google.generativeai as genai
from fastapi import FastAPI
from confluent_kafka import Producer, Consumer, KafkaException
from pydantic import BaseModel

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI(title="Voice AI Service (Kafka-Driven)")

KAFKA_BROKERS = os.getenv("KAFKA_BROKERS", "kafka:9092")
VOICE_REQUEST_TOPIC = os.getenv("VOICE_REQUEST_TOPIC", "voice-requests")
VOICE_RESPONSE_TOPIC = os.getenv("VOICE_RESPONSE_TOPIC", "voice-results")
VOICE_CONSUMER_GROUP = os.getenv("VOICE_CONSUMER_GROUP", "voice-ai-service")

GEMINI_MODEL = os.getenv("GEMINI_MODEL", "gemini-2.5-flash")
TRANSCRIPTION_MODEL = os.getenv("TRANSCRIPTION_MODEL", "whisper-1")

CIRCUIT_FAILURE_THRESHOLD = int(
    os.getenv("VOICE_CIRCUIT_FAILURE_THRESHOLD", "5")
)
CIRCUIT_RESET_SECONDS = int(
    os.getenv("VOICE_CIRCUIT_RESET_SECONDS", "15")
)

openai_api_key = os.getenv("OPENAI_API_KEY")
google_api_key = os.getenv("GOOGLE_API_KEY")

producer: Producer | None = None
consumer: Consumer | None = None


class CircuitBreaker:
    def __init__(self, failure_threshold: int, reset_timeout: int):
        self.failure_threshold = failure_threshold
        self.reset_timeout = reset_timeout
        self.state = "closed"
        self.failures = 0
        self.last_failure = 0.0

    def allow(self) -> bool:
        now = time.time()

        if (
            self.state == "open"
            and now - self.last_failure >= self.reset_timeout
        ):
            self.state = "half_open"
            logger.info("Circuit breaker transitioning to half_open")
            return True

        return self.state != "open"

    def success(self) -> None:
        self.state = "closed"
        self.failures = 0
        self.last_failure = 0.0
        logger.info("Circuit breaker closed")

    def failure(self) -> None:
        self.failures += 1
        self.last_failure = time.time()

        logger.warning(
            "Circuit breaker failure count increased",
            extra={
                "failures": self.failures,
                "threshold": self.failure_threshold,
            },
        )

        if self.failures >= self.failure_threshold:
            self.state = "open"
            logger.warning("Circuit breaker opened")


circuit_breaker = CircuitBreaker(
    CIRCUIT_FAILURE_THRESHOLD,
    CIRCUIT_RESET_SECONDS,
)


class VoiceResultData(BaseModel):
    amount: float = 0.0
    description: str = ""


class VoiceKafkaRequest(BaseModel):
    request_id: str
    user_id: str
    text: str | None = None
    audio_base64: str | None = None
    created_at: str


class VoiceKafkaResponse(BaseModel):
    request_id: str
    user_id: str
    transcript: str
    intent: str
    data: VoiceResultData
    status: str
    processed_at: str
    error: str | None = None


def configure_apis() -> None:
    if openai_api_key:
        openai.api_key = openai_api_key
        logger.info("Configured OpenAI for audio transcription")
    else:
        logger.warning(
            "OPENAI_API_KEY not set; audio transcription unavailable"
        )

    if google_api_key:
        genai.configure(api_key=google_api_key)
        logger.info("Configured Gemini API")
    else:
        logger.warning(
            "GOOGLE_API_KEY not set; Gemini parsing unavailable"
        )


def transcribe_audio_base64(audio_base64: str) -> str:
    """
    Decode base64 audio and transcribe using OpenAI Whisper.
    Uses unique temp files to avoid concurrency issues.
    """

    if not openai_api_key:
        return "[audio transcription disabled]"

    temp_path = None

    try:
        decoded = base64.b64decode(audio_base64)

        with tempfile.NamedTemporaryFile(
            suffix=".wav",
            delete=False,
        ) as tmp:
            tmp.write(decoded)
            temp_path = tmp.name

        with open(temp_path, "rb") as audio_file:
            transcription = openai.Audio.transcribe(
                TRANSCRIPTION_MODEL,
                audio_file,
            )

        transcript = getattr(transcription, "text", "").strip()

        logger.info(
            "Audio transcription completed",
            extra={"transcript": transcript[:100]},
        )

        return transcript

    except Exception as exc:
        logger.error(
            "Audio transcription failed",
            exc_info=exc,
        )
        return "[audio transcription failed]"

    finally:
        if temp_path and os.path.exists(temp_path):
            try:
                os.remove(temp_path)
            except Exception:
                pass


def parse_voice_text(text: str) -> dict:
    """
    Parse voice command into structured transaction data using Gemini.
    """

    if not google_api_key:
        return {
            "amount": 0.0,
            "description": text[:120],
            "is_transaction": False,
        }

    prompt = f"""
Extract transaction information from this voice command.

Voice command:
{text}

Return ONLY valid JSON in this exact schema:

{{
  "amount": <monetary amount as float, 0.0 if not found>,
  "description": "<brief transaction description>",
  "is_transaction": <true if this is a valid transaction request, false otherwise>
}}
"""

    text_output = ""

    try:
        model = genai.GenerativeModel(GEMINI_MODEL)

        response = model.generate_content(
            prompt,
            generation_config={
                "temperature": 0,
                "max_output_tokens": 256,
                "response_mime_type": "application/json",
            },
        )

        text_output = response.text.strip()

        logger.info(
            "Gemini raw response received",
            extra={"response": text_output},
        )

        parsed = json.loads(text_output)

        normalized = {
            "amount": float(parsed.get("amount", 0.0)),
            "description": str(parsed.get("description", "")).strip(),
            "is_transaction": bool(
                parsed.get("is_transaction", False)
            ),
        }

        logger.info(
            "Gemini parsing successful",
            extra={"parsed": normalized},
        )

        return normalized

    except json.JSONDecodeError as exc:
        logger.error(
            "Invalid JSON returned from Gemini",
            extra={"raw_output": text_output},
            exc_info=exc,
        )

    except Exception as exc:
        logger.error(
            "Gemini parsing failed",
            exc_info=exc,
        )

    circuit_breaker.failure()

    return {
        "amount": 0.0,
        "description": text[:120],
        "is_transaction": False,
    }


def build_voice_response(
    request: VoiceKafkaRequest,
    status: str,
    data: dict,
    error: str | None = None,
) -> VoiceKafkaResponse:

    return VoiceKafkaResponse(
        request_id=request.request_id,
        user_id=request.user_id,
        transcript=data.get("transcript", ""),
        intent=data.get("intent", "unknown"),
        data=VoiceResultData(
            amount=float(data.get("amount", 0.0)),
            description=str(data.get("description", "")),
        ),
        status=status,
        processed_at=datetime.utcnow().isoformat() + "Z",
        error=error,
    )


def process_voice_request(payload: dict) -> VoiceKafkaResponse:
    request = VoiceKafkaRequest(**payload)

    transcript = request.text or ""

    if request.audio_base64:
        transcript = (
            transcribe_audio_base64(request.audio_base64)
            or transcript
        )

    if not transcript:
        transcript = "[empty voice input]"

    if not circuit_breaker.allow():
        logger.warning(
            "Voice processing blocked by circuit breaker",
            extra={"request_id": request.request_id},
        )

        return build_voice_response(
            request,
            "failed",
            {
                "transcript": transcript,
                "intent": "failed",
            },
            "circuit breaker open",
        )

    parsed = parse_voice_text(transcript)

    if parsed.get("is_transaction"):
        circuit_breaker.success()
        intent = "transaction.create"
        status = "completed"
        error = None
    else:
        intent = "unknown"
        status = "failed"
        error = "voice command not recognized as transaction"

    result = {
        "transcript": transcript,
        "intent": intent,
        "amount": parsed.get("amount", 0.0),
        "description": parsed.get(
            "description",
            transcript,
        ),
    }

    return build_voice_response(
        request=request,
        status=status,
        data=result,
        error=error,
    )


def delivery_report(err, msg):
    if err is not None:
        logger.error(f"Kafka delivery failed: {err}")
    else:
        logger.info(
            f"Message delivered to {msg.topic()} [{msg.partition()}]"
        )


def produce_response(response: VoiceKafkaResponse) -> None:
    if producer is None:
        logger.error("Kafka producer is not initialized")
        return

    try:
        payload = json.dumps(
            response.dict()
        ).encode("utf-8")

        producer.produce(
            VOICE_RESPONSE_TOPIC,
            key=response.request_id,
            value=payload,
            callback=delivery_report,
        )

        producer.flush()

    except Exception as exc:
        logger.error(
            "Failed to publish response",
            exc_info=exc,
        )


def consume_requests() -> None:
    global consumer

    while True:
        try:
            msg = consumer.poll(1.0)

            if msg is None:
                continue

            if msg.error():
                logger.error(
                    f"Kafka consumer error: {msg.error()}"
                )
                continue

            payload = json.loads(
                msg.value().decode("utf-8")
            )

            logger.info(
                "Voice request received",
                extra={
                    "request_id": payload.get("request_id"),
                },
            )

            response = process_voice_request(payload)

            produce_response(response)

        except Exception as exc:
            logger.error(
                "Failed to process voice request",
                exc_info=exc,
            )


@app.on_event("startup")
async def startup_event() -> None:
    global producer, consumer

    configure_apis()

    max_retries = 10
    retry_delay = 2  # initial delay in seconds
    connected = False

    for attempt in range(1, max_retries + 1):
        try:
            logger.info(f"Attempting to connect to Kafka (Attempt {attempt}/{max_retries})...")
            
            # Initialize Producer
            producer = Producer({
                "bootstrap.servers": KAFKA_BROKERS,
            })

            # Initialize Consumer
            consumer = Consumer({
                "bootstrap.servers": KAFKA_BROKERS,
                "group.id": VOICE_CONSUMER_GROUP,
                "auto.offset.reset": "earliest",
            })

            # Trigger a metadata call to verify if the broker is actually accessible.
            # confluent-kafka initializes lazily; this forces a real connection check.
            producer.list_topics(timeout=3.0)

            # Subscribe to the target topic
            consumer.subscribe([VOICE_REQUEST_TOPIC])

            logger.info(
                "Voice AI service successfully connected to Kafka",
                extra={
                    "request_topic": VOICE_REQUEST_TOPIC,
                    "response_topic": VOICE_RESPONSE_TOPIC,
                },
            )
            connected = True
            break  # Connection successful, break out of retry loop

        except (KafkaException, Exception) as exc:
            logger.warning(
                f"Kafka connection attempt {attempt} failed. Retrying in {retry_delay}s...",
                extra={"error": str(exc)}
            )
            time.sleep(retry_delay)
            retry_delay = min(retry_delay * 2, 30)  # Exponential backoff up to 30 seconds

    if not connected:
        logger.error("Critical: Failed to initialize Kafka clients after maximum retries. Shutting down.")
        raise RuntimeError("Kafka cluster unreachable during application startup.")

    # Safe to spin up your worker thread now that clients are verified
    threading.Thread(
        target=consume_requests,
        daemon=True,
    ).start()


@app.get("/health")
async def health_check() -> dict:
    return {
        "status": "ok",
        "service": "voice-ai-service",
        "kafka_brokers": KAFKA_BROKERS,
        "voice_request_topic": VOICE_REQUEST_TOPIC,
        "voice_response_topic": VOICE_RESPONSE_TOPIC,
        "circuit_state": circuit_breaker.state,
        "circuit_failures": circuit_breaker.failures,
    }
@app.get("/voice/status")
async def voice_status():
    return {"status": "ok", "service": "voice-ai-service"}
