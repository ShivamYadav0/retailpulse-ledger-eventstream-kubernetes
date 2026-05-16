#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PYTHON_VENV_DIR="$ROOT_DIR/.venv"
FRONTEND_DIR="$ROOT_DIR/client-vite"
PYTHON_SERVICE_DIR="$ROOT_DIR/services/voice-ai-service"
BIN_DIR="$ROOT_DIR/bin"

usage() {
  cat <<EOF
Usage: $0 <command>

Commands:
  install      Install dependencies for Go, Python voice service, and frontend
  build        Build Go binaries, frontend assets, and Python bytecode
  build-images Build all service Docker images
  clean        Remove generated binaries, frontend build artifacts, and Python venv
EOF
}

ensure_command() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "ERROR: required command '$1' is not installed or not in PATH." >&2
    exit 1
  fi
}

install_go() {
  echo "==> Installing Go module dependencies"
  cd "$ROOT_DIR"
  ensure_command go
  go mod download
}

install_python() {
  echo "==> Installing Python dependencies for voice-ai-service"
  ensure_command python3
  ensure_command pip
  python3 -m venv "$PYTHON_VENV_DIR"
  # shellcheck disable=SC1091
  source "$PYTHON_VENV_DIR/bin/activate"
  python -m pip install --upgrade pip
  pip install --no-cache-dir -r "$PYTHON_SERVICE_DIR/requirements.txt"
  deactivate
}

install_frontend() {
  echo "==> Installing frontend dependencies"
  ensure_command npm
  cd "$FRONTEND_DIR"
  npm install
}

build_go() {
  echo "==> Building Go services"
  ensure_command go
  mkdir -p "$BIN_DIR"

  cd "$ROOT_DIR"
  go build -o "$BIN_DIR/ledger-service" ./services/ledger-service/cmd/api
  go build -o "$BIN_DIR/batch-worker" ./services/batch-worker/main.go
  go build -o "$BIN_DIR/notification-service" ./services/notification-service/main.go
  echo "Go binaries built into $BIN_DIR"
}

build_python() {
  echo "==> Building Python service artifacts"
  ensure_command python3
  if [[ ! -d "$PYTHON_VENV_DIR" ]]; then
    echo "Python virtualenv not found. Run '$0 install' first." >&2
    exit 1
  fi
  # shellcheck disable=SC1091
  source "$PYTHON_VENV_DIR/bin/activate"
  python -m compileall "$PYTHON_SERVICE_DIR"
  deactivate
}

build_frontend() {
  echo "==> Building frontend assets"
  ensure_command npm
  cd "$FRONTEND_DIR"
  npm run build
}

build_images() {
  echo "==> Building Docker images"
  ensure_command docker
  cd "$ROOT_DIR"

  docker build -t smart-retail-dep-microservices-ledger-service:latest -f services/ledger-service/Dockerfile .
  docker build -t smart-retail-dep-microservices-batch-worker:latest -f services/batch-worker/Dockerfile .
  docker build -t smart-retail-dep-microservices-notification-service:latest -f services/notification-service/Dockerfile .
  docker build -t smart-retail-dep-microservices-voice-ai-service:latest -f services/voice-ai-service/Dockerfile .
}

clean() {
  echo "==> Cleaning generated artifacts"
  rm -rf "$BIN_DIR"
  rm -rf "$PYTHON_VENV_DIR"
  cd "$FRONTEND_DIR"
  rm -rf node_modules dist
  echo "Clean complete"
}

main() {
  if [[ $# -lt 1 ]]; then
    usage
    exit 1
  fi

  case "$1" in
    install)
      install_go
      install_python
      install_frontend
      ;;
    build)
      build_go
      build_python
      build_frontend
      ;;
    build-images)
      build_images
      ;;
    clean)
      clean
      ;;
    *)
      usage
      exit 1
      ;;
  esac
}

main "$@"
