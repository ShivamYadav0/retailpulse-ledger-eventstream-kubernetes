# Smart Retail — Detailed Architecture & Runbook

This file contains an in-depth description of the architecture, service responsibilities, deployment instructions (Docker Compose and k3d/Kubernetes), diagrams, operational notes, and developer pointers.

## Overview

Smart Retail is an event-driven microservices platform focused on reliable transaction processing, voice-driven input parsing, durable ledger persistence, and asynchronous notifications. The design centers on a single transaction gateway and Kafka-backed messaging for decoupling.

## High-level diagram

```mermaid
flowchart LR
  Client[Frontend client] -->|POST /v1/transaction| Gateway[Ledger Gateway]
  Gateway -->|Kafka: voice-requests| VoiceAI[voice-ai-service]
  VoiceAI -->|Kafka: voice-results| Gateway
  Gateway -->|Kafka: transactions| Batch[batch-worker]
  Gateway -->|Kafka: transactions| Notify[notification-service]
  Batch -->|Postgres| PostgresDB[(Postgres)]
  Notify -->|Push/SMS/Email| Users[Users]
```

## Components

- `services/ledger-service` — Transaction gateway. Validates requests, publishes `voice-requests` when voice input is present, consumes `voice-results`, and emits `transactions` events.
- `services/voice-ai-service` — Kafka consumer that runs voice parsing (Gemini + transcription). Protected by a circuit breaker to avoid external AI downtime cascading.
- `services/batch-worker` — Kafka consumer that persists transactions to Postgres.
- `services/notification-service` — Kafka consumer that delivers notifications asynchronously.
- `pkg/` — Shared libraries: logging, Kafka helpers, resilience utilities.
- `deployments/k8s/` — Kubernetes manifests for cluster deployment.

## Gateway flow (detailed)

1. Client posts to `/v1/transaction` with either direct transaction fields or audio/text.
2. Gateway validates and, if audio/text present, publishes `voice-requests` (correlation id included).
3. `voice-ai-service` consumes `voice-requests` and publishes `voice-results` with the correlation id.
4. Gateway consumes the `voice-results` (or receives a timeout/error), validates the parsed intent, and publishes the final `transactions` event.
5. Downstream consumers (`batch-worker`, `notification-service`) process the `transactions` topic.

## Kubernetes (k3d) quick runbook

Prerequisites: `k3d`, `kubectl`, `docker`, and `make`.

1. Create k3d cluster:

```bash
k3d cluster create vani-ledger --api-port 6550 -p "80:80@loadbalancer"
```

2. Build images and import into k3d (local build):

```bash
make build-images
k3d image import smart-retail-dep-* -c vani-ledger
```

3. Apply manifests:

```bash
kubectl apply -f deployments/k8s/namespace.yml
kubectl apply -f deployments/k8s/configMap/init-scripts.yaml
kubectl apply -f deployments/k8s/traefik/ingressroute.yml
kubectl apply -R -f deployments/k8s
```

4. Verify:

```bash
kubectl get pods -n vani-ledger
kubectl get svc -n vani-ledger
```

5. If you need to import images built locally into k3d, use:

```bash
k3d image import <image-name>:<tag> -c vani-ledger
```

## Running the included k6 loadtest inside k3d

If you want to run performance tests inside the cluster (so network visibility matches production), create a `ConfigMap` with the script and run a `Job` that uses the `grafana/k6` image.

Example (from root):

```bash
kubectl -n default create configmap k6-script --from-file=scripts/loadtest.js
cat <<'EOF' | kubectl apply -f -
apiVersion: batch/v1
kind: Job
metadata:
  name: k6-loadtest
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: k6
        image: grafana/k6:latest
        command: ["k6", "run", "/scripts/loadtest.js"]
        volumeMounts:
        - name: script
          mountPath: /scripts
      volumes:
      - name: script
        configMap:
          name: k6-script
EOF

kubectl logs -f job/k6-loadtest
kubectl delete job/k6-loadtest configmap/k6-script
```

## Environment variables

Create a `.env` with secure values before running locally or in k3d:

```bash
OPENAI_API_KEY=...
GOOGLE_API_KEY=...
POSTGRES_URL=postgres://user:pass@postgres:5432/ledger
KAFKA_BROKER=kafka:9092
```

## Observability and recommendations

- Add Prometheus metrics and Grafana dashboards for request rates and latency.
- Export Kafka consumer lag metrics to monitor processing health.
- Configure rolling updates and readiness/liveness probes for all services.

## Troubleshooting

- If services can't reach Kafka in k3d, check `kubectl get svc` and ensure DNS names match the manifests.
- Use `kubectl logs -n vani-ledger <pod>` to inspect service logs.

## Links

- Gateway details: [gateway.md](gateway.md)
- Docker Compose wiring: [docker-compose.yml](docker-compose.yml)
- Scripts: [scripts/](scripts)


## Contribution

Please open issues or PRs when improving the voice parsing logic, adding CI/CD, or hardening production configuration.
