 832  chmod +x kubectl
  833  sudo mv kubectl /usr/local/bin/
  834  kubectl version --client
  835  curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
  836  k3d version
  837  sudo su
  838  k3d cluster create vani-ledger
  839  sudo usermod -aG docker $USER
  840  docker ps
  841  sudo systemctl restart sssd
  842  systemctl status sssd
  843  sudo rm -f /var/lib/sss/db/*
  844  sudo systemctl restart sssd
  845  systemctl status sssd
  846  sudo usermod -aG docker $USER
  847  sudo chmod 666 /var/run/docker.sock
  848  docker ps
  849  sudo rm -f /usr/local/bin/kubectl
  850  curl -LO "https://dl.k8s.io/release/v1.31.0/bin/linux/amd64/kubectl"
  851  chmod +x kubectl
  852  sudo mv kubectl /usr/local/bin/
  853  kubectl version --client
  854  k3d cluster list
  855  kubectl get nodes
  856  docker images
  857  image: smart-retail-dep-microservices-ledger-service:latest
  858  kubectl get namespaces
  859  docker exec -it k3d-vani-ledger-server-0 crictl images
  860  kubectl get pods -n vani-ledger
  861  kubectl get deployments -n vani-ledger
  862  kubectl get svc -n vani-ledger
  863  kubectl get statefulsets -n vani-ledger
  864  kubectl logs -f deployment/ledger-service -n vani-ledger
  865  kubectl get pods -n vani-ledger
  866  kubectl logs -f deployment/ledger-service -n vani-ledger
  867  kubectl logs -f kafka-0 -n vani-ledger
  868  kubectl logs -f deployment/ledger-service -n vani-ledger
  869  kubectl get pods -n vani-ledger
  870  kubectl logs -f deployment/ledger-service -n vani-ledger
  871  kubectl rollout restart deployment ledger-service -n vani-ledger
  872  kubectl logs -f deployment/ledger-service -n vani-ledger
  873  kubectl rollout restart deployment ledger-service -n vani-ledger
  874  kubectl logs -f deployment/ledger-service -n vani-ledger
  875  kubectl get pods -n vani-ledger
  876  kubectl get svc -n vani-ledger
  877  kubectl rollout restart deployment batch-worker -n vani-ledger
  878  kubectl get pods -n vani-ledger
  879  kubectl rollout restart deployment batch-worker -n vani-ledger
  880  kubectl get pods -n vani-ledger
  881  kubectl delete deployment batch-worker -n vani-ledger
  882  kubectl delete deployment notification-service -n vani-ledger
  883  kubectl delete deployment voice-ai-service -n vani-ledger
  884  kubectl get pods -n vani-ledger
  885  kubectl logs -f deployment/ledger-service -n vani-ledger
  886  kubectl get pods -n vani-ledger
  887  kubectl logs -f deployment/ledger-service -n vani-ledger
  888  kubectl logs -f deployment/batch-worker -n vani-ledger
  889  kubectl logs -f deployment/voice-ai-service -n vani-ledger
  890  k3d image import smart-retail-dep-microservices-batch-worker:latest -c vani-ledger
  891  k3d image import smart-retail-dep-microservices-notification-service:latest -c vani-ledger
  892  k3d image import smart-retail-dep-microservices-voice-ai-service:latest -c vani-ledger
  893  docker exec -it k3d-vani-ledger-server-0 crictl images
  894  kubectl rollout restart deployment batch-worker -n vani-ledger
  895  kubectl rollout restart deployment voice-ai-service -n vani-ledger
  896  kubectl rollout restart deployment notification-service -n vani-ledger
  897  kubectl get pods -n vani-ledger
  898  kubectl logs -f deployment/batch-worker -n vani-ledger
  899  kubectl logs -f deployment/voice-ai-service -n vani-ledger
  900  kubectl logs -f deployment/ledger-service -n vani-ledger
  901  kubectl logs -f deployment/batch-worker -n vani-ledger
  902  kubectl get pods -n vani-ledger
  903  kubectl logs -f deployment/batch-worker -n vani-ledger
  904  kubectl get pods -n vani-ledger
  905  kubectl logs -f deployment/voice-ai-service -n vani-ledger
  906  kubectl get pods -n vani-ledger
  907  kubectl logs -f deployment/batch-worker -n vani-ledger
  908  kubectl logs -f deployment/ledger-service -n vani-ledger
  909  kubectl get svc -n vani-ledger
  910  kubectl get ingressroute -n vani-ledger
  911  kubectl describe ingressroute ledger-ingress -n vani-ledger
  912  k3d cluster create vani-ledger -p "80:80@loadbalancer"
  913  curl http://localhost/v1
  914  http://localhost/voice
  915  curl http://localhost/voice
  916  curl http://localhost/v1
  917  curl http://localhost/voice
  918  kubectl get pods -A | grep traefik
  919  kubectl get svc -n vani-ledger
  920  kubectl get pods -n vani-ledger
  921  docker ps
  922  kubectl get pods -n vani-ledger
  923  kubectl get svc -n vani-ledger
  924  curl http://localhost/voice
  925  curl http://localhost/v1
  926  k3d cluster delete vani-ledger
  927  k3d cluster create vani-ledger -p "80:80@loadbalancer"
  928  kubectl get pods -n vani-ledger
  929  kubectl get svc -n vani-ledger
  930  k3d image import smart-retail-dep-microservices-ledger-service:latest -c vani-ledger
  931  k3d image import smart-retail-dep-microservices-batch-worker:latest -c vani-ledger
  932  k3d image import smart-retail-dep-microservices-notification-service:latest -c vani-ledger
  933  k3d image import smart-retail-dep-microservices-voice-ai-service:latest -c vani-ledger
  934  curl http://localhost/v1
  935  curl http://localhost/voice
  936  curl http://localhost/v1
  937  kubectl get pods -n vani-ledger
  938  kubectl get endpoints -n vani-ledger
  939  kubectl logs deployment/voice-ai-service -n vani-ledger
  940  kubectl exec -it deployment/voice-ai-service -n vani-ledger -- sh
  941  kubectl get endpoints -n vani-ledger
  942  kubectl exec -it deployment/voice-ai-service -n vani-ledger -- sh
  943  kubectl get endpoints -n vani-ledger
  944  kubectl logs deployment/voice-ai-service -n vani-ledger
  945  kubectl get endpoints -n vani-ledger
  946  kubectl logs deployment/ledger-service -n vani-ledger
  947  kubectl logs deployment/voice-ai-service -n vani-ledger
  948  curl http://localhost/v1
  949  curl http://localhost/voice
  950  docker images
  951  curl http://localhost/v1
  952  kubectl logs deployment/voice-ai-service -n vani-ledger
  953  kubectl get endpoints -n vani-ledger
  954  kubectl logs deployment/voice-ai-service -n vani-ledger
  955  apiVersion: v1
  956  kubectl get secrets -n vani-ledger
  957  kubectl rollout restart deployment voice-ai-service -n vani-ledger
  958  kubectl logs -f deployment/voice-ai-service -n vani-ledger
  959  history
   66  kubectl get pods -A
   67  kubectl exec -it kafka-tools -n vani-ledger -- bash
   68  kubectl run kafka-tools -n vani-ledger --rm -it   --image=bitnami/kafka:latest --restart=Never -- bash
   69  kubectl logs -f deployment/batch-worker -n vani-ledger
   70  kubectl logs -f deployment/ledger-service -n vani-ledger
   71  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
   72  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --create --topic transactions --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
   73  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
   74  kubectl get pods -A
   75  kubectl logs -f deployment/ledger-service -n vani-ledger
   76  kubectl logs -f deployment/batch-worker -n vani-ledger
   77  kubectl logs -f deployment/ledger-service -n vani-ledger
   78  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
   79  kubectl logs -f deployment/ledger-service -n vani-ledger
   80  kubectl logs -f deployment/batch-worker -n vani-ledger
   81  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --create --topic transactions --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
   82  kubectl logs -f deployment/ledger-service -n vani-ledger
   83  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
   84  kubectl logs -f deployment/batch-worker -n vani-ledger
   85  [shivamy@shivamy ~]$ kubectl logs -f deployment/ledger-service -n vani-ledger
   86  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic transactions
   87  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic transactions --from-beginning
   88  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
   89  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker-group
   90  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker
   91  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic transactions
   92  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker
   93  watch -n 2 '
kubectl exec -it kafka-0 -n vani-ledger -- \
/opt/kafka/bin/kafka-consumer-groups.sh \
--bootstrap-server localhost:9092 \
--describe \
--group batch-worker-group
   94  watch -n 2 '
kubectl exec -it kafka-0 -n vani-ledger -- \
/opt/kafka/bin/kafka-consumer-groups.sh \
--bootstrap-server localhost:9092 \
--describe \
--group batch-worker-group
'
   95  watch -n 2 '
kubectl exec -it kafka-0 -n vani-ledger -- \
/opt/kafka/bin/kafka-consumer-groups.sh \
--bootstrap-server localhost:9092 \
--describe \
--group batch-worker
'
   96  kubectl exec -it kafka-0 -n vani-ledger -- /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic transactions
   97  kubectl logs -f deployment/voice-ai-service -n vani-ledger
   98  history
