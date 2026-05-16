Here’s a consolidated list of the **Kafka `docker exec` commands** we discussed, along with what each one does. This way you have a clear cheat‑sheet for working with your broker and topics:

---

### 🔎 Topic Management
- **List all topics**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server kafka:9092 --list
  ```
  ➝ Shows all topics currently in the broker (e.g. `transactions`, `voice-requests`, `voice-results`).

- **Create a topic**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh \
    --bootstrap-server kafka:9092 \
    --create --topic voice-results --partitions 1 --replication-factor 1
  ```
  ➝ Creates the `voice-results` topic (or any topic you specify).

---

### 📥 Consuming Messages
- **Consume from a topic (from beginning)**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --topic voice-results \
    --from-beginning
  ```
  ➝ Streams all messages in `voice-results` starting at the earliest offset.

- **Consume requests**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --topic voice-requests \
    --from-beginning
  ```
  ➝ Streams all messages in `voice-requests`.

---

### 📤 Producing Messages
- **Produce into a topic**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-producer.sh \
    --bootstrap-server kafka:9092 \
    --topic voice-requests
  ```
  ➝ Opens a prompt where you can paste JSON payloads (e.g. a voice request). Each line you enter is sent as a Kafka message.

---

### 👥 Consumer Group Management
- **Describe a consumer group**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh \
    --bootstrap-server kafka:9092 \
    --group ledger-voice-results-consumer \
    --describe
  ```
  ➝ Shows partitions, current offsets, log end offsets, and lag for the group `ledger-voice-results-consumer`.

- **Reset offsets (requires group inactive)**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh \
    --bootstrap-server kafka:9092 \
    --group ledger-voice-results-consumer \
    --reset-offsets --to-earliest --execute --topic voice-results
  ```
  ➝ Resets the group’s offsets to the earliest position. Only works if the group is inactive (no consumers connected).

---

### 🛠️ Debugging
- **Check broker environment variables**
  ```bash
  docker exec -it smart-retail-dep-microservices-kafka-1 printenv | grep KAFKA
  ```
  ➝ Shows Kafka configuration (listeners, advertised listeners, auto‑create flag, etc.).

- **View broker logs**
  ```bash
  docker logs smart-retail-dep-microservices-kafka-1
  ```
  ➝ Displays Kafka broker logs for debugging startup, listener, or coordinator issues.

---

This set of commands covers **topic creation/listing**, **producing/consuming messages**, and **consumer group management**.  

Would you like me to also add a **ready‑to‑use JSON payload example** for `voice-requests` that you can paste directly into the producer console, so you can immediately test the full pipeline end‑to‑end?