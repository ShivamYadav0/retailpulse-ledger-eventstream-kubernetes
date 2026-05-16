# Vani Ledger Frontend

A simple Vite + React frontend for the ledger platform.

## Run locally

1. Start your backend services and gateway stack so that:
   - `http://localhost/v1` reaches the ledger service
   - `http://localhost/voice` reaches the voice AI service

2. Start the frontend:
   ```bash
   cd client-vite
   npm install
   npm run dev
   ```

3. Open the local Vite URL shown in the terminal.

## Usage

- Enter a `User ID`.
- Type a text prompt or record a voice request.
- Press `Send Text` or `Record Voice`.
- The app sends the data through the gateway and creates a ledger transaction.
- The ledger history is fetched and displayed after the request completes.

## Notes

- The app proxies `/voice` and `/v1` through Vite to the local gateway.
- The voice request is sent as base64 audio to `/voice/process`.
- The ledger transaction is sent to `/v1/transaction`.
- Ledger entries are loaded from `/v1/ledger/:userId`.
