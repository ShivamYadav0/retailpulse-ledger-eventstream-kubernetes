#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
K3D_CLUSTER_NAME="${K3D_CLUSTER_NAME:-vani-ledger}"
K3D_AGENT_COUNT="${K3D_AGENT_COUNT:-1}"
K3D_LB_PORT="${K3D_LB_PORT:-80}"
K3D_API_PORT="${K3D_API_PORT:-6550}"

usage() {
  cat <<EOF
Usage: $0 <command>

Commands:
  create-cluster   Create a k3d cluster named ${K3D_CLUSTER_NAME}
  apply            Apply Kubernetes manifests with kubectl
  import-images    Import locally built Docker images into k3d
  deploy           Create cluster, import images, and apply manifests
  delete-cluster   Delete the k3d cluster
  help             Show this help message
EOF
}

ensure_command() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "ERROR: required command '$1' is not installed." >&2
    exit 1
  fi
}

create_cluster() {
  ensure_command k3d
  ensure_command kubectl

  if k3d cluster list | grep -q "^${K3D_CLUSTER_NAME}\b"; then
    echo "Cluster '${K3D_CLUSTER_NAME}' already exists. Skipping creation."
    return
  fi

  echo "==> Creating k3d cluster '${K3D_CLUSTER_NAME}'"
  k3d cluster create "${K3D_CLUSTER_NAME}" \
    --agents "${K3D_AGENT_COUNT}" \
    --api-port "${K3D_API_PORT}" \
    -p "${K3D_LB_PORT}:80@loadbalancer"

  echo "==> Setting kubectl context to k3d-${K3D_CLUSTER_NAME}"
  kubectl config use-context "k3d-${K3D_CLUSTER_NAME}"
}

apply_manifests() {
  ensure_command kubectl

  echo "==> Applying Kubernetes manifests"
  cd "$ROOT_DIR/deployments/k8s"
  kubectl apply -f namespace.yml
  
  kubectl apply -R -f .
}

import_images() {
  ensure_command k3d
  ensure_command docker

  echo "==> Importing service images into k3d cluster '${K3D_CLUSTER_NAME}'"
  k3d image import smart-retail-dep-microservices-ledger-service:latest -c "${K3D_CLUSTER_NAME}"
  k3d image import smart-retail-dep-microservices-voice-ai-service:latest -c "${K3D_CLUSTER_NAME}"
  k3d image import smart-retail-dep-microservices-batch-worker:latest -c "${K3D_CLUSTER_NAME}"
  k3d image import smart-retail-dep-microservices-notification-service:latest -c "${K3D_CLUSTER_NAME}"
}

deploy() {
  create_cluster
  import_images
  apply_manifests
}

delete_cluster() {
  ensure_command k3d
  echo "==> Deleting k3d cluster '${K3D_CLUSTER_NAME}'"
  k3d cluster delete "${K3D_CLUSTER_NAME}"
}

main() {
  if [[ $# -lt 1 ]]; then
    usage
    exit 1
  fi

  case "$1" in
    create-cluster)
      create_cluster
      ;;
    apply)
      apply_manifests
      ;;
    import-images)
      import_images
      ;;
    deploy)
      deploy
      ;;
    delete-cluster)
      delete_cluster
      ;;
    help|--help|-h)
      usage
      ;;
    *)
      echo "ERROR: Unknown command: $1" >&2
      usage
      exit 1
      ;;
  esac
}

main "$@"
