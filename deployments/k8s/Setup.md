# Local Kubernetes Setup with k3d + k3s + kubectl + Docker

This guide explains:

* How Kubernetes works locally using k3d
* How to install everything on RHEL8/Linux
* How to deploy services
* How Docker and Kubernetes interact
* How to rebuild applications after code changes
* How Kubernetes networking works
* How to debug pod failures
* How Ingress + Traefik routing works
* How to use Secrets
* Common errors and fixes
* Important kubectl commands

---

# Architecture Overview

Your setup:

```text
Browser / curl
       ↓
Traefik Ingress
       ↓
Kubernetes Service
       ↓
Pods
       ↓
Applications
```

Infrastructure:

```text
k3d
 └── k3s Kubernetes Cluster
      ├── Traefik
      ├── PostgreSQL
      ├── PgBouncer
      ├── Redis
      ├── Kafka
      ├── ledger-service
      ├── voice-ai-service
      ├── batch-worker
      └── notification-service
```

---

# What is k3d?

k3d runs Kubernetes inside Docker containers.

Instead of:

```text
Real VMs
```

it creates:

```text
Docker containers acting as Kubernetes nodes
```

Very lightweight.
Very fast.
Perfect for local development.

---

# Components Explanation

## Docker

Builds container images.

Example:

```bash
docker build -t ledger-service:latest .
```

---

## k3d

Creates local Kubernetes cluster.

Example:

```bash
k3d cluster create vani-ledger
```

---

## kubectl

CLI used to communicate with Kubernetes.

Example:

```bash
kubectl get pods
```

---

## k3s

Lightweight Kubernetes distribution.

k3d internally runs k3s.

---

# Install Docker

## RHEL8

```bash
sudo dnf install -y dnf-plugins-core

sudo dnf config-manager \
--add-repo \
https://download.docker.com/linux/centos/docker-ce.repo

sudo dnf install -y \
docker-ce docker-ce-cli containerd.io
```

---

# Start Docker

```bash
sudo systemctl enable docker
sudo systemctl start docker
```

Verify:

```bash
docker ps
```

---

# Add User To Docker Group

```bash
sudo usermod -aG docker $USER
```

IMPORTANT:

Logout and login again.

OR:

```bash
newgrp docker
```

Verify:

```bash
docker ps
```

Should work WITHOUT sudo.

---

# Install kubectl

```bash
curl -LO "https://dl.k8s.io/release/stable.txt"
```

Then:

```bash
VERSION=$(curl -L -s https://dl.k8s.io/release/stable.txt)
```

Download:

```bash
curl -LO "https://dl.k8s.io/release/${VERSION}/bin/linux/amd64/kubectl"
```

Install:

```bash
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```

Verify:

```bash
kubectl version --client
```

---

# Install k3d

```bash
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

Verify:

```bash
k3d version
```

---

# Create Kubernetes Cluster

```bash
k3d cluster create vani-ledger -p "80:80@loadbalancer"
```

Explanation:

```text
80:80
```

maps local machine port 80 to Traefik load balancer.

---

# Verify Cluster

```bash
kubectl cluster-info
```

```bash
kubectl get nodes
```

Expected:

```text
NAME                       STATUS
k3d-vani-ledger-server-0  Ready
```

---

# Why Namespace?

Namespace separates applications logically.

Example:

```text
vani-ledger
```

contains ONLY your application resources.

Avoids conflicts.

---

# Create Namespace

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: vani-ledger
```

Apply:

```bash
kubectl apply -f namespace.yaml
```

Verify:

```bash
kubectl get namespaces
```

---

# Kubernetes Folder Structure

```text
k8s/
 ├── namespace.yaml
 │
 ├── postgres/
 ├── redis/
 ├── kafka/
 ├── pgbouncer/
 │
 ├── ledger-service/
 ├── voice-ai-service/
 ├── batch-worker/
 ├── notification-service/
 │
 ├── ingress/
 │
 └── secrets/
```

---

# Apply Kubernetes Resources

Apply everything:

```bash
kubectl apply -R -f k8s/
```

Explanation:

```text
-R
```

means recursive.

Kubernetes reads all YAML files.

---

# Verify Resources

## Pods

```bash
kubectl get pods -n vani-ledger
```

---

## Services

```bash
kubectl get svc -n vani-ledger
```

---

## StatefulSets

```bash
kubectl get statefulsets -n vani-ledger
```

---

## Deployments

```bash
kubectl get deployments -n vani-ledger
```

---

# Important Kubernetes Objects

## Deployment

Used for stateless apps.

Examples:

* ledger-service
* voice-ai-service
* notification-service

---

## StatefulSet

Used for stateful systems.

Examples:

* PostgreSQL
* Kafka

Provides:

* stable hostname
* persistent storage

---

## Service

Internal load balancer.

Example:

```text
ledger-service:8080
```

Pods communicate through services.

---

## Ingress

Routes external traffic.

Example:

```text
/v1 → ledger-service
/voice → voice-ai-service
```

---

# How Networking Works

## Internal Kubernetes DNS

Every service automatically gets DNS.

Example:

```text
redis
kafka
postgres
ledger-service
```

Pods can connect using service names.

Example:

```text
redis:6379
kafka:9092
```

---

# Why Docker Compose And Kubernetes Are Separate

Docker Compose:

```text
Direct Docker containers
```

Kubernetes:

```text
Pods managed by k3d/k3s
```

These are DIFFERENT environments.

Avoid running both simultaneously.

---

# Stop Docker Compose

```bash
docker compose down
```

Verify:

```bash
docker ps
```

Should show only:

```text
k3d containers
```

---

# Building Images

Example:

```bash
docker build \
-t smart-retail-dep-microservices-ledger-service:latest \
-f services/ledger-service/Dockerfile .
```

---

# Why ImagePullBackOff Happens

Kubernetes cannot access local Docker images automatically.

Need:

```bash
k3d image import
```

---

# Import Images Into k3d

## Ledger Service

```bash
k3d image import \
smart-retail-dep-microservices-ledger-service:latest \
-c vani-ledger
```

---

## Voice AI Service

```bash
k3d image import \
smart-retail-dep-microservices-voice-ai-service:latest \
-c vani-ledger
```

---

## Batch Worker

```bash
k3d image import \
smart-retail-dep-microservices-batch-worker:latest \
-c vani-ledger
```

---

## Notification Service

```bash
k3d image import \
smart-retail-dep-microservices-notification-service:latest \
-c vani-ledger
```

---

# Restart Deployments

After importing image:

```bash
kubectl rollout restart deployment ledger-service \
-n vani-ledger
```

---

# Watch Deployment Progress

```bash
kubectl rollout status deployment ledger-service \
-n vani-ledger
```

---

# View Logs

## Deployment Logs

```bash
kubectl logs -f deployment/ledger-service \
-n vani-ledger
```

---

## Specific Pod Logs

```bash
kubectl logs -f POD_NAME -n vani-ledger
```

---

# Execute Shell Inside Pod

```bash
kubectl exec -it deployment/ledger-service \
-n vani-ledger -- sh
```

Useful for:

* checking env vars
* checking ports
* testing connectivity

---

# Verify Listening Ports

Inside pod:

```bash
netstat -tulpn
```

OR:

```bash
ss -tulpn
```

Need:

```text
0.0.0.0:8080
```

NOT:

```text
127.0.0.1:8080
```

---

# Readiness Probe

Readiness means:

```text
Can pod receive traffic?
```

Example:

```yaml
readinessProbe:
  httpGet:
    path: /health
    port: 8080
```

---

# Why READY 0/1 Happens

Common reasons:

* wrong readiness path
* app crashing
* wrong port
* app listening on localhost only

---

# Verify Endpoints

```bash
kubectl get endpoints -n vani-ledger
```

Bad:

```text
ledger-service <none>
```

Good:

```text
ledger-service 10.x.x.x:8080
```

---

# Why Traefik Returns Service Unavailable

Flow:

```text
Traefik
 ↓
Service
 ↓
NO READY PODS
```

So traffic has nowhere to go.

---

# Common Debugging Commands

## Describe Pod

```bash
kubectl describe pod POD_NAME -n vani-ledger
```

Shows:

* events
* image pull errors
* readiness failures
* crash reasons

---

## Get All Resources

```bash
kubectl get all -n vani-ledger
```

---

## Watch Pods Live

```bash
kubectl get pods -n vani-ledger -w
```

---

## Delete Pod

Kubernetes recreates automatically.

```bash
kubectl delete pod POD_NAME -n vani-ledger
```

---

# Access Application

## Through Ingress

```bash
curl http://localhost/v1/health
```

```bash
curl http://localhost/voice/health
```

---

# Port Forwarding

Useful for debugging.

## Ledger Service

```bash
kubectl port-forward svc/ledger-service \
8080:8080 -n vani-ledger
```

Then:

```bash
curl http://localhost:8080/health
```

---

# Kafka Setup

Kafka must have:

```yaml
ports:
  - name: kafka
    port: 9092

  - name: controller
    port: 9093
```

Port names are required.

---

# Kafka Common Error

## Not Coordinator For Group

Usually occurs during startup.

Kafka consumer retries automatically.

Often temporary.

---

# Secrets

## Secret YAML

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
  namespace: vani-ledger

type: Opaque

stringData:
  OPENAI_API_KEY: sk-xxxxx
```

---

# Apply Secret

```bash
kubectl apply -f k8s/secrets/
```

---

# Use Secret In Deployment

```yaml
env:
  - name: OPENAI_API_KEY
    valueFrom:
      secretKeyRef:
        name: app-secret
        key: OPENAI_API_KEY
```

---

# Restart After Secret Change

```bash
kubectl rollout restart deployment voice-ai-service \
-n vani-ledger
```

Pods do NOT automatically reload env vars.

---

# Verify Secrets Loaded

```bash
kubectl exec -it deployment/voice-ai-service \
-n vani-ledger -- sh
```

Inside pod:

```bash
env | grep OPENAI
```

---

# Clean Everything

## Delete Namespace

```bash
kubectl delete namespace vani-ledger
```

Deletes all resources inside namespace.

---

# Delete Entire Cluster

```bash
k3d cluster delete vani-ledger
```

---

# Full Development Workflow

```text
Code change
 ↓
Docker build
 ↓
k3d image import
 ↓
kubectl rollout restart
 ↓
kubectl logs
 ↓
Verify readiness
 ↓
Test APIs
```

---

# Recommended Scripts

## deploy-ledger.sh

```bash
#!/bin/bash

set -e

docker build \
-t smart-retail-dep-microservices-ledger-service:latest \
-f services/ledger-service/Dockerfile .

k3d image import \
smart-retail-dep-microservices-ledger-service:latest \
-c vani-ledger

kubectl rollout restart deployment ledger-service \
-n vani-ledger

kubectl rollout status deployment ledger-service \
-n vani-ledger
```

---

# Important Concepts Summary

## Docker

Builds images.

---

## Kubernetes

Runs containers as pods.

---

## Pod

Running container instance.

---

## Deployment

Manages stateless pods.

---

## StatefulSet

Manages stateful systems.

---

## Service

Internal networking/load balancing.

---

## Ingress

External HTTP routing.

---

## Traefik

Ingress controller.

---

## Namespace

Logical isolation.

---

## Readiness Probe

Determines whether pod receives traffic.

---

## k3d

Local Kubernetes using Docker.

---

# Final Verification Checklist

## Cluster

```bash
kubectl get nodes
```

---

## Pods

```bash
kubectl get pods -n vani-ledger
```

Need:

```text
1/1 Running
```

---

## Services

```bash
kubectl get svc -n vani-ledger
```

---

## Endpoints

```bash
kubectl get endpoints -n vani-ledger
```

Should NOT be empty.

---

## Logs

```bash
kubectl logs -f deployment/ledger-service \
-n vani-ledger
```

---

## Ingress Access

```bash
curl http://localhost/v1/health
```

```bash
curl http://localhost/voice/health
```

---

# Production Improvements Later

For production you would add:

* Helm
* ArgoCD
* CI/CD pipelines
* TLS certificates
* Horizontal Pod Autoscaler
* Monitoring
* Prometheus
* Grafana
* Loki
* Distributed tracing
* Kafka replication
* PostgreSQL HA
* Redis persistence
* Network policies
* RBAC
* Resource quotas
* Service mesh

* Docker
* kubectl
* k3d
* Kubernetes concepts
* namespaces
* deployments
* services
* ingress
* Traefik
* Kafka
* Redis
* PostgreSQL
* StatefulSets
* image rebuild flow
* k3d image import
* rollout restart
* readiness probes
* debugging commands
* logs
* endpoints
* secrets
* networking
* port-forwarding
* troubleshooting
* full development workflow
* cleanup commands
* production improvements


---

## 🔍 Check Existing Clusters
List all clusters:
```bash
k3d cluster list
```
This shows cluster names, servers, and agents.

---

## 🗑 Delete All Clusters
To delete everything:
```bash
k3d cluster delete --all
```
That removes all clusters created with k3d.

---

## 🧹 Clean Up Namespaces (inside a cluster)
If you only want to clear resources inside a running cluster (not delete the cluster itself):

1. List namespaces:
   ```bash
   kubectl get namespaces
   ```

2. Delete everything in a namespace (example: `default`):
   ```bash
   kubectl delete all --all -n default
   ```

3. To delete the namespace itself:
   ```bash
   kubectl delete namespace <namespace-name>
   ```

---

## 🔄 Fresh Install Workflow
1. Delete all clusters:
   ```bash
   k3d cluster delete --all
   ```
2. Create a new cluster:
   ```bash
   k3d cluster create mycluster
   ```
3. Re‑install your workloads (operators, deployments, etc.).

---

## ⚠️ Notes
- `k3d cluster delete --all` is the nuclear option — it wipes everything.  
- If you only want to reset workloads but keep the cluster, use `kubectl delete all --all -n <namespace>`.  
- Always double‑check which cluster/context you’re connected to:
  ```bash
  kubectl config current-context
  ```
