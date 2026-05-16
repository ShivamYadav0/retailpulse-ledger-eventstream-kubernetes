import { useEffect, useMemo, useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';

type VoiceResponse = {
  request_id: string;
  status: string;
  intent: string;
  error?: string;
  transcript?: string;
  data?: {
    user_id: string;
    amount: number;
    description: string;
  };
  transaction?: Transaction;
};

type Transaction = {
  id: string;
  user_id: string;
  amount: number;
  description: string;
  created_at: string;
};

type QueueStatus = {
  queue_length: number;
  total_processed: number;
};

const DEFAULT_USER = '00000000-0000-0000-0000-000000000001';
const PAGE_SIZE = 6;
const POLL_INTERVAL = 1500;
const MAX_POLL_ATTEMPTS = 30;

function App() {
  const { t, i18n } = useTranslation();

  const [userId, setUserId] = useState(DEFAULT_USER);
  const [text, setText] = useState(t('promptExample'));

  const [ledger, setLedger] = useState<Transaction[]>([]);
  const [voiceResponse, setVoiceResponse] = useState<VoiceResponse | null>(null);

  const [status, setStatus] = useState(t('ready'));
  const [errorMessage, setErrorMessage] = useState('');

  const [recording, setRecording] = useState(false);
  const [recordingSeconds, setRecordingSeconds] = useState(0);

  const [loading, setLoading] = useState(false);
  const [polling, setPolling] = useState(false);

  const [queueStatus, setQueueStatus] = useState<QueueStatus>({
    queue_length: 0,
    total_processed: 0,
  });

  const [page, setPage] = useState(1);

  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const chunksRef = useRef<Blob[]>([]);
  const timerRef = useRef<number | null>(null);

  const totalPages = Math.max(1, Math.ceil(ledger.length / PAGE_SIZE));

  const visibleTransactions = useMemo(
    () => ledger.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE),
    [ledger, page],
  );

    useEffect(() => {
    fetchLedger(DEFAULT_USER);
   
  }, []);


  async function fetchLedger(id: string) {
    try {
      const res = await fetch(`/v1/ledger/${id}`);

      const data = await res.json();

      if (!res.ok) {
        setErrorMessage(data.error || 'Failed to load ledger');
        return;
      }

      setLedger(data.data ?? []);
      setPage(1);

    } catch (error) {
      setErrorMessage(String(error));
    }
  }

  function blobToBase64(blob: Blob): Promise<string> {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();

      reader.onloadend = () => {
        const base64 = reader.result?.toString().split(',')[1] ?? '';
        resolve(base64);
      };

      reader.onerror = reject;

      reader.readAsDataURL(blob);
    });
  }

  function startTimer() {
    setRecordingSeconds(0);

    timerRef.current = window.setInterval(() => {
      setRecordingSeconds((prev) => prev + 1);
    }, 1000);
  }

  function stopTimer() {
    if (timerRef.current) {
      clearInterval(timerRef.current);
      timerRef.current = null;
    }
  }

  async function startRecording() {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({
        audio: true,
      });

      const recorder = new MediaRecorder(stream);

      chunksRef.current = [];

      recorder.ondataavailable = (event) => {
        chunksRef.current.push(event.data);
      };

      recorder.onstop = async () => {
        stopTimer();

        setRecording(false);

        const blob = new Blob(chunksRef.current, {
          type: 'audio/webm',
        });

        const base64 = await blobToBase64(blob);

        await submitRequest({
          user_id: userId,
          audio_base64: base64,
        });
      };

      mediaRecorderRef.current = recorder;

      recorder.start();

      startTimer();

      setRecording(true);
      setStatus(t('recording'));
      setErrorMessage('');

    } catch (error) {
      setErrorMessage(String(error));
    }
  }

  function stopRecording() {
    const recorder = mediaRecorderRef.current;

    if (!recorder) return;

    recorder.stop();

    mediaRecorderRef.current = null;
  }

  async function submitRequest(payload: {
    user_id: string;
    text?: string;
    audio_base64?: string;
  }) {

    try {

      setLoading(true);
      setErrorMessage('');
      setVoiceResponse(null);

      setStatus('Sending request...');

      const res = await fetch('/v1/transaction', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      const data = await res.json();

      if (!res.ok) {
        setErrorMessage(data.error || 'Request failed');
        setLoading(false);
        return;
      }

      const requestId = data.data?.request_id;

      if (!requestId) {
        setErrorMessage('Missing request id');
        setLoading(false);
        return;
      }

      setStatus('Queued for processing');

      await pollVoiceStatus(requestId);

    } catch (error) {

      setErrorMessage(String(error));

    } finally {

      setLoading(false);
    }
  }

  async function pollVoiceStatus(requestId: string) {

    setPolling(true);

    for (let attempt = 1; attempt <= MAX_POLL_ATTEMPTS; attempt++) {

      try {

        setStatus(`Processing request (${attempt}/${MAX_POLL_ATTEMPTS})`);

        const res = await fetch(`/v1/voice-status/${requestId}`);
console.log('Polling voice status, attempt', attempt, res.status, res);
        if (!res.ok) {

          await sleep(POLL_INTERVAL);
          continue;
        }

        const data = await res.json();
console.log('Voice status response', data);
        const response: VoiceResponse = data.data;
console.log('Voice status response', response);
        if (!response) {

          await sleep(POLL_INTERVAL);
          continue;
        }


        setVoiceResponse(data);

        if (data.status === 'processing') {

          await sleep(POLL_INTERVAL);
          continue;
        }

        if (data.status === 'failed') {

          setStatus('Voice processing failed');

          setErrorMessage(
            data.error || 'Voice processing failed',
          );

          setPolling(false);

          return;
        }

        if (data.status === 'completed') {

          setStatus('Transaction created successfully');

          if (data.transaction) {

            setLedger((prev) => [
              data.transaction!,
              ...prev,
            ]);
          } else {
            await fetchLedger(userId);
          }

          setPolling(false);

          return;
        }

      } catch (error) {

        setErrorMessage(String(error));
      }

      await sleep(POLL_INTERVAL);
    }

    setPolling(false);

    setStatus('Timed out waiting for AI response');
  }

  function sleep(ms: number) {
    return new Promise((resolve) => {
      setTimeout(resolve, ms);
    });
  }

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault();

    await submitRequest({
      user_id: userId,
      text,
    });
  }

  function switchLanguage(code: 'en' | 'hi') {
    i18n.changeLanguage(code);
  }

  function goPreviousPage() {
    setPage((prev) => Math.max(1, prev - 1));
  }

  function goNextPage() {
    setPage((prev) => Math.min(totalPages, prev + 1));
  }

  return (
    <div className="app-shell">

      <div className="card">

        <div className="top-bar">

          <div>
            <h1>{t('appTitle')}</h1>
            <p>{t('appDescription')}</p>
          </div>

          <div className="language-switcher">
            <button
              type="button"
              onClick={() => switchLanguage('en')}
              className={i18n.language === 'en' ? 'active' : ''}
            >
              EN
            </button>

            <button
              type="button"
              onClick={() => switchLanguage('hi')}
              className={i18n.language === 'hi' ? 'active' : ''}
            >
              हिंदी
            </button>
          </div>
        </div>

        <form onSubmit={handleSubmit}>

          <div className="field-group">
            <label>{t('userIdLabel')}</label>

            <input
              value={userId}
              onChange={(e) => setUserId(e.target.value)}
            />
          </div>

          <div className="field-group">

            <label>{t('promptLabel')}</label>

            <textarea
              value={text}
              rows={4}
              onChange={(e) => setText(e.target.value)}
            />
          </div>

          <div className="button-row">

            <button
              type="submit"
              disabled={loading || polling}
            >
              {loading || polling
                ? 'Processing...'
                : t('sendText')}
            </button>

            <button
              type="button"
              onClick={
                recording
                  ? stopRecording
                  : startRecording
              }
              className={recording ? 'danger' : ''}
              disabled={loading || polling}
            >
              {recording
                ? `Stop (${recordingSeconds}s)`
                : t('recordVoice')}
            </button>
          </div>
        </form>

        <div className="status-panel">

          <div>
            <strong>Status:</strong> {status}
          </div>

          <div className="status-meta">
            Queue: {queueStatus.queue_length}
            {' · '}
            Processed: {queueStatus.total_processed}
          </div>

          {errorMessage && (
            <div className="error-message">
              {errorMessage}
            </div>
          )}
        </div>

        {voiceResponse && (

          <div className="box">

            <h2>Voice Processing Result</h2>

            <div className="result-grid">

              <div>
                <strong>Request ID</strong>
                <div>{voiceResponse.request_id}</div>
              </div>

              <div>
                <strong>Status</strong>
                <div>{voiceResponse.status}</div>
              </div>

              <div>
                <strong>Intent</strong>
                <div>{voiceResponse.intent}</div>
              </div>

              <div>
                <strong>Amount</strong>
                <div>
                  {voiceResponse.data?.amount ?? 0}
                </div>
              </div>

              <div>
                <strong>Description</strong>
                <div>
                  {voiceResponse.data?.description}
                </div>
              </div>
            </div>

            {voiceResponse.transcript && (
              <>
                <strong>Transcript</strong>

                <pre>
                  {voiceResponse.transcript}
                </pre>
              </>
            )}
          </div>
        )}

        <div className="box">

          <div className="section-header">

            <h2>{t('ledgerEntries')}</h2>

            <span>
              Page {page} / {totalPages}
            </span>
          </div>

          {ledger.length === 0 ? (

            <p>No ledger entries found</p>

          ) : (

            <>
              <table>

                <thead>
                  <tr>
                    <th>User</th>
                    <th>Amount</th>
                    <th>Description</th>
                    <th>Created</th>
                  </tr>
                </thead>

                <tbody>

                  {visibleTransactions.map((entry) => (

                    <tr key={entry.id}>

                      <td>
                        {entry.user_id.slice(0, 8)}
                      </td>

                      <td>
                        {entry.amount}
                      </td>

                      <td>
                        {entry.description}
                      </td>

                      <td>
                        {new Date(
                          entry.created_at,
                        ).toLocaleString()}
                      </td>

                    </tr>
                  ))}
                </tbody>
              </table>

              <div className="pagination">

                <button
                  type="button"
                  onClick={goPreviousPage}
                  disabled={page === 1}
                >
                  Previous
                </button>

                <span>
                  {page} / {totalPages}
                </span>

                <button
                  type="button"
                  onClick={goNextPage}
                  disabled={page === totalPages}
                >
                  Next
                </button>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
export default App;