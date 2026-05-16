
docker run --rm \
  -v "$PWD/scripts":/scripts \
  -w /scripts \
  --network host \
  grafana/k6:latest \
  run loadtest.js