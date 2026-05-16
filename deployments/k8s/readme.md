
# 🚀 Vani Ledger Deployment Guide

This guide walks through deploying the **Vani Ledger ecosystem** (ledger-service, voice-ai-service, batch-worker, notification-service, and supporting infrastructure) on Kubernetes using Helm and Traefik.

---

## 📦 Prerequisites

- Kubernetes cluster up and running (local via Minikube, Docker Desktop, or cloud provider).
- `kubectl` installed and configured.
- `helm` installed (v3+).
- Docker installed (with access to `/var/run/docker.sock`).
- Basic knowledge of Kubernetes manifests and Helm charts.

---

## ⚙️ Step 1: Install Traefik via Helm

Traefik will act as the ingress controller.

```bash
helm repo add traefik https://traefik.github.io/charts
helm repo update

helm install traefik traefik/traefik \
  --namespace traefik \
  --create-namespace
```

This creates the Traefik ingress controller in its own namespace.

---

## ⚙️ Step 2: Apply Namespaces and Secrets

Create the application namespace and apply secrets:

```bash
kubectl apply -f namespace.yaml
kubectl apply -f secrets.yaml
```

- `namespace.yaml` → defines the `vani-ledger` namespace.
- `secrets.yaml` → stores sensitive credentials (DB passwords, API keys, etc.).

---

## ⚙️ Step 3: Deploy Infrastructure Services

Apply manifests for supporting services:

```bash
kubectl apply -f redis/
kubectl apply -f postgres/
kubectl apply -f pgbouncer/
kubectl apply -f kafka/
```

- **Redis** → caching and pub/sub.
- **Postgres** → primary database.
- **PgBouncer** → lightweight connection pooler for Postgres.
- **Kafka** → event streaming backbone.

---

## ⚙️ Step 4: Deploy Application Services

```bash
kubectl apply -f ledger-service/
kubectl apply -f voice-ai-service/
kubectl apply -f batch-worker/
kubectl apply -f notification-service/
kubectl apply -f traefik/
```

- **Ledger Service** → core financial transaction logic.
- **Voice AI Service** → speech-to-text and AI-driven voice features.
- **Batch Worker** → background jobs and scheduled tasks.
- **Notification Service** → email/SMS/push notifications.
- **Traefik configs** → ingress routes for external access.

---

## ⚙️ Step 5: Docker Socket Permissions

Ensure Docker socket is accessible:

```bash
sudo chmod 666 /var/run/docker.sock
```

---

## ✅ Step 6: Verification

### Pods
Check all pods are running:

```bash
kubectl get pods -n vani-ledger
```

Expected: All pods should be in `Running` or `Completed` state.

---

### Services
List services:

```bash
kubectl get svc -n vani-ledger
```

Expected: Services for Redis, Postgres, Kafka, and each microservice.

---

### Ingress
Verify Traefik ingress routes:

```bash
kubectl get ingressroute -n vani-ledger
```

Expected: Routes for ledger-service, voice-ai-service, etc.

---

### Logs
Check logs for ledger-service:

```bash
kubectl logs -f deployment/ledger-service -n vani-ledger
```

Use this to debug startup issues.

---

## 🔍 Troubleshooting

- **Pods not starting** → Run `kubectl describe pod <pod-name> -n vani-ledger` to check events.
- **CrashLoopBackOff** → Inspect logs with `kubectl logs`.
- **Ingress not working** → Ensure Traefik is running and ingress routes are applied.
- **Database connection issues** → Verify secrets and PgBouncer configuration.

#helm reference
https://copilot.microsoft.com/shares/w57LdbvREtXuz6ENqunNt
