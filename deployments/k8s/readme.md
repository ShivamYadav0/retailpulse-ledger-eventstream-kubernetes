helm repo add traefik https://traefik.github.io/charts

helm repo update

helm install traefik traefik/traefik \
  --namespace traefik \
  --create-namespace

kubectl apply -f namespace.yaml

kubectl apply -f secrets.yaml

kubectl apply -f redis/

kubectl apply -f postgres/

kubectl apply -f pgbouncer/

kubectl apply -f kafka/

kubectl apply -f ledger-service/

kubectl apply -f voice-ai-service/

kubectl apply -f batch-worker/

kubectl apply -f notification-service/

kubectl apply -f traefik/


sudo chmod 666 /var/run/docker.sock
Verify Everything
Pods
kubectl get pods -n vani-ledger
Services
kubectl get svc -n vani-ledger
Ingress
kubectl get ingressroute -n vani-ledger
Logs
kubectl logs -f deployment/ledger-service -n vani-ledger

https://copilot.microsoft.com/shares/w57LdbvREtXuz6ENqunNt
