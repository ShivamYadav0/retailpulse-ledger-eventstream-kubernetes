   554  k3d image import smart-retail-dep-microservices-voice-ai-service:latest -c vani-ledger
  555  k3d image import smart-retail-dep-microservices-ledger-service:latest -c vani-ledger
  556  k3d image import smart-retail-dep-microservices-batch-worker:latest -c vani-ledger
  557  k3d image import smart-retail-dep-microservices-notification-service:latest -c vani-ledger
  558  kubectl rollout restart deployment ledger-service -n vani-ledger
  559  kubectl rollout status deployment ledger-service -n vani-ledger
  560  kubectl apply -f k8s
  561  cd deployments/
  562  kubectl apply -f k8s
  563  kubectl apply -R -f /k8s
  564  kubectl apply -R -f k8s/
  565    871  kubectl rollout restart deployment ledger-service -n vani-ledger
  566   kubectl rollout restart deployment ledger-service -n vani-ledger
  567  kubectl rollout restart deployment batch-worker -n vani-ledger
  568   kubectl rollout restart deployment voice-ai-service -n vani-ledger
  569  kubectl rollout restart deployment notification-service -n vani-ledger
  570  kubectl get pods -A
  571  cd client-vite/
  572  kubectl logs -f traefik-5d45fc8cc9-n5z8n -n vani-ledger
  573  kubectl logs -f traefik-5d45fc8cc9-n5z8n -n kube-system
  574  kubectl logs -f deployment/ledger-service -n vani-ledger
  575  kubectl logs -f deployment/voice-ai-service -n vani-ledger
  576  code .
  577  cd ..

  775  docker start c0aa32258adf
  776  docker logs --tail 100 -f  c0aa32258adf
  777  docker ps -a
  778  docker logs --tail 100 -f da6c9a5162b5
  779  docker logs --tail 100 -f  c0aa32258adf
  780  docker logs --tail 100 -f  73e943d9dfe5
  781  docker ps -a
  782  dockrt exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker
  783  docket exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker
  784  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group batch-worker
  785  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe 
  786  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --all-topics 
  787  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
  788  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh   --bootstrap-server localhost:9092   --group voice-ai-service   --describe
  789  kafka-consumer-groups.sh --bootstrap-server localhost:9092   --group voice-ai-service   --reset-offsets --to-earliest --execute --topic voice-requests
  790  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh   --bootstrap-server localhost:9092   --group voice-ai-service   --describedocker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh   --bootstrap-server kafka:9092   --create   --topic voice-results   --partitions 1   --replication-factor 1
  791  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh   --bootstrap-server kafka:9092   --create   --topic voice-results   --partitions 1   --replication-factor 1
  792  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-producer.sh   --bootstrap-server kafka:9092   --topic voice-requests
  793  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-results   --from-beginning
  794  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-request   --from-beginning
  795  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-requests   --from-beginning
  796  docker logs --tail 100 -f  73e943d9dfe5
  797  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-requests   --from-beginning
  798  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-producer.sh   --bootstrap-server kafka:9092   --topic voice-requests
  799  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-consumer-groups.sh   --bootstrap-server kafka:9092   --group voice-ai-service   --describe
  800  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-results   --from-beginning
  801  docker logs --tail 100 -f  73e943d9dfe5
  802  docker logs --tail 100 -f  c0aa32258adf
  803  docker compose up --build
  804  docker compose down
  805  docker compose up -d
  806  docler ps -a
  807  docker p -a
  808  docker ps -a
  809  docker logs 1eae30ed46ee
  810  docker logs c6e9131f9bfb
  811  docker logs 1eae30ed46ee
  812  docker logs c6e9131f9bfb
  813  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-results   --from-beginning
  814  docker ps -a
  815  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-results   --from-beginning
  816  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh   --bootstrap-server kafka:9092   --topic voice-requests   --from-beginning
  817  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --list
  818  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --create --topic voice-requests --partitions 1 --replication-factor 1
  819  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic voice-requests
  820  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic voice-requests
  821  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic voice-results
  822  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic voice-results
  823  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --create --topic voice-requests --partitions 1 --replication-factor 1
  824  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --create --topic voice-requests --partitions 1 --replication-factor 1
  825  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --create --topic voice-results --partitions 1 --replication-factor 1
  826  docker ps -a
  827  docker logs 1eae30ed46ee
  828  docker logs c6e9131f9bfb
  829  docker logs 1eae30ed46ee
  830  docker logs c6e9131f9bfb
  831  docker exec -it da6c9a5162b5 /opt/kafka/bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list kafka:9092 --topic voice-requests
  832  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list kafka:9092 --topic voice-requests
  833  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --all-groups --describe
  834  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic voice-requests
  835  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-requests --from-beginning --timeout-ms 5000
  836  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-requests --from-beginning
  837  docker exec -it smart-retail-dep-microservices-kafka-1 env | grep KAFKA
  838  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-producer.sh --bootstrap-server kafka:9092 --topic voice-requests
  839  docker compose down -v
  840  docker volume ls
  841  docker volume rm $(docker volume ls -q)
  842  docker volume ls
  843  docker compose up -d
  844  docker logs kafka -f
  845  docker ps -a
  846  docker logs smart-retail-dep-microservices-ledger-service
  847  docker logs 4b09685d8e41
  848  docker logs 85433665e366
  849  docker logs 4b09685d8e41
  850  docker logs 85433665e366
  851  docker exec -it kafka /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-requests --from-beginning
  852  docker ps -a
  853  docker exec -it d33d69455571 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-requests --from-beginning
  854  docker exec -it d33d69455571 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  855  ./del.sh 
  856  docker volume ls
  857  docker compose up -d
  858  ./del.sh 
  859  docker compose up --build
  860  go mod tidy
  861  docker compose up --build
  862  go mod tidy
  863  docker ps -a
  864  docker logs 50733b89eab7
  865  docker ps -a
  866  docker logs 435a6e59f171
  867  docker logs 50733b89eab7
  868  docker logs 435a6e59f171
  869  docker compose up --build
  870  ./del.sh 
  871  docker ps -a
  872  docker logs 08088b112fc2
  873  docker logs a75819e1d9fa
  874  docker logs 08088b112fc2
  875  docker ps -a
  876  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  877  docker ps -a
  878  docker exec -it 85dabdae7978 redis-cli
  879  docker logs -f smart-retail-dep-microservices-ledger-service-1
  880  docker compose up --build
  881  ./del.sh 
  882  docker ps -a
  883  docker logs -f smart-retail-dep-microservices-ledger-service-1
  884  docker ps -a
  885  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  886  docker exec -it e9d90f727699 redis-cli
  887  docker logs -f smart-retail-dep-microservices-ledger-service-1
  888  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  889  docker logs -f smart-retail-dep-microservices-voice-ai-service
  890  clear
  891  docker logs -f 692691787d26
  892  docker logs -f smart-retail-dep-microservices-voice-ai-service
  893  clear
  894  docker logs -f 692691787d26
  895  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  896  docker logs -f smart-retail-dep-microservices-ledger-service-1
  897  docker exec -it e9d90f727699 redis-cli
  898  docker compose up --build
  899  ./del.sh 
  900  docker compose up --build
  901  docker ps -a
  902  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  903  docker exec -it 0b883090ad7b redis-cli
  904  docker logs 126481f2c4d2
  905  ./del.sh 
  906  docker ps -a
  907  docker logs 81cf8e6c968b
  908  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  909  docker exec -it smart-retail-dep-microservices-redis-1 redis-cli
  910  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --list
  911  docker compose up --build
  912  ./del.sh 
  913  docker compose up --build
  914  docker ps -a
  915  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --list
  916  docker logs dc0ef247fcb0
  917  docker logs 0ecd3f3768cf
  918  docker logs dc0ef247fcb0
  919  docker logs 0ecd3f3768cf
  920  docker logs dc0ef247fcb0
  921  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  922  docker logs dc0ef247fcb0
  923  docker exec -it smart-retail-dep-microservices-redis-1 redis-cli
  924  docker logs dc0ef247fcb0
  925  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --describe --group ledger-voice-results-consumer
  926  docker exec -it smart-retail-dep-microservices-kafka-1 printenv | grep KAFKA
  927  ./del.sh 
  928  docker compose up --build
  929  docker exec -it smart-retail-dep-microservices-redis-1 redis-cli
  930  docker ps -a
  931  docker logs c1921889cb70
  932  docker logs 26342d163ecb
  933  docker logs 591167ac5c74
  934  docker exec -it smart-retail-dep-microservices-kafka-1 printenv | grep KAFKA
  935  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --describe --group ledger-voice-results-consumer
  936  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
  937  docker logs c1921889cb70
  938  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh   --bootstrap-server kafka:9092   --group ledger-voice-results-consumer   --reset-offsets --to-earliest --execute --topic voice-results
  939  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --group ledger-voice-results-consumer --reset-offsets --to-earliest --execute --topic voice-results
  940  docker start smart-retail-dep-microservices-ledger-service-1
  941  docker logs -f smart-retail-dep-microservices-ledger-service-1
  942  docker restart smart-retail-dep-microservices-ledger-service-1
  943  docker logs -f smart-retail-dep-microservices-ledger-service-1
  944  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --group ledger-voice-results-consumer --reset-offsets --to-earliest --execute --topic voice-results
  945  docker restart smart-retail-dep-microservices-ledger-service-1
  946  docker ps -a
  947  docker stop c1921889cb70
  948  docker start c1921889cb70
  949  history
  950  docker restart $(docker ps -aq)
  951  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server kafka:9092 --group ledger-voice-results-consumer --reset-offsets --to-earliest --execute --topic voice-results
  952  docker logs -f smart-retail-dep-microservices-ledger-service-1
  953  docker exec -it smart-retail-dep-microservices-redis-1 redis-cli
  954  docker stop c1921889cb70
  955  docker start c1921889cb70
  956  docker logs -f smart-retail-dep-microservices-ledger-service-1
  957  podman ps -a
  958  docker ps -a
  959  docker logs -f 591167ac5c74
  960  docker compose up -d
  961  docker restart $(docker ps -aq)
  962  docker ps -a
  963  docker logs -f smart-retail-dep-microservices-ledger-service-1
  964  docker ps -a
  965  docker logs b2396776023b
  966  docker ps -a
  967  docker logs 5f13ab082ea2
  968  npm run dev
  969  sudo systemctl stop zsaservice.service 
  970  docker compose down -v
  971  clear
  972  cd client-vite/
  973  npm run dev
  974  docker ps -a
  975  docker logs smart-retail-dep-microservices-ledger-service-1
  976  docker compose up --build
  977  docker ps -a
  978  docker logs smart-retail-dep-microservices-ledger-service-1
  979  docker restart $(docker ps -aq)
  980  docker logs smart-retail-dep-microservices-ledger-service-1
  981  cd client
  982  cd ../client-vite/
  983  docker compose down -v
  984  npm run dev
  985  k3d cluster delete --all
  986  docker rm -f $(docker ps -aq)
  987  docker volume rm $(docker volume ls -q)
  988  docker network rm $(docker network ls -q)
  989  docker volume ls
  990  docker images
  991  ./del.sh 
  992  clear
  993  docker compose up --build
  994  docker ps -a
  995  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
  996  docker logs --tail 100 -f smart-retail-dep-microservices-voice-ai-service-1
  997  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
  998  history
  999  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
 1000  docker ps -a
 1001  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1002  docker exec -it smart-retail-dep-microservices-redis-1 redic-cli
 1003  docker exec -it smart-retail-dep-microservices-redis-1 redis-cli
 1004  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1005  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
 1006  docker restart $(docker ps -aq)
 1007  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1008  docker logs --tail 100 -f smart-retail-dep-microservices-voice-ai-service-1
 1009  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1010  clear
 1011  docker ps -a
 1012  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1013  docker logs --tail 100 -f smart-retail-dep-microservices-voice-ai-service-1
 1014  docker logs --tail 100 -f smart-retail-dep-microservices-kafka-1
 1015  docker logs --tail 100 -f smart-retail-dep-microservices-ledger-service-1
 1016  docker logs --tail 100 -f smart-retail-dep-microservices-voice-ai-service-1
 1017  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic voice-results --from-beginning
 1018  docker ps -a
 1019  docker logs --tail 100 -f smart-retail-dep-microservices-notification-service-1
 1020  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic --list
 1021  history
 1022  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --describe --all-topics
 1023  docker exec -it smart-retail-dep-microservices-kafka-1 /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka:9092 --list
 1024  history


kubectl exec -it kafka-0 -n vani-ledger -- \
/opt/kafka/bin/kafka-topics.sh \
--bootstrap-server localhost:9092 \
--list





 1271  ./del.sh 
 1272  make build-images 
 1273  make k8s-deploy 
 1274  kubectl get pods -A
 1275  kubectl logs statefulset/postgres -n vani-ledger --tail=20
 1276  kubectl get pods -A
 1277  kubectl exec -it statefulset/postgres -n vani-ledger -- cat /etc/postgresql/pg_hba.conf
 1278  kubectl rollout restart deployment pgbouncer -n vani-ledger
 1279  kubectl logs -f deployment/pgbouncer -n vani-ledger
 1280  kubectl rollout restart deployment batch-worker -n vani-ledger
 1281  kubectl logs -f deployment/batch-worker -n vani-ledger

 1284  kubectl logs -f deployment/batch-worker -n vani-ledger
 1285  kubectl logs -f deployment/pgbouncer -n vani-ledger
 1286  kubectl exec -it statefulset/postgres -n vani-ledger -- cat /etc/postgresql/pg_hba.conf
 1287  kubectl exec -it statefulset/postgres -n vani-ledger -- psql -U ledger -d ledger -c "SELECT pg_reload_conf();"
 1288  kubectl exec -it statefulset/postgres -n vani-ledger -- cp /etc/postgresql/pg_hba.conf /var/lib/postgresql/data/pg_hba.conf
 1289  kubectl exec -it statefulset/postgres -n vani-ledger -- psql -U ledger -d ledger -c "SELECT pg_reload_conf();"
 1290  kubectl logs -f deployment/batch-worker -n vani-ledger
 1291  kubectl get pods -A
 1292  kubectl rollout restart deployment pgbouncer -n vani-ledger
 1293  kubectl get pods -A
 1294  kubectl rollout restart deployment batch-worker -n vani-ledger
 1295  kubectl logs -f deployment/batch-worker -n vani-ledger
 1296  cd /home/shivamy/Videos/smart-retail-dep-Microservices/deployments/k8s/postgres
 1297  kubectl apply -f .
 1298  kubectl exec -it statefulset/postgres -n vani-ledger -- cat /etc/postgresql/pg_hba.conf
 1299  kubectl logs statefulset/postgres -n vani-ledger --tail=20
 1300  history



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


k3d cluster delete --all
docker rm -f $(docker ps -aq)
docker volume rm $(docker volume ls -q)
docker network rm $(docker network ls -q)

make build-images
    make k8s-deploy
   kubectl get pods -n vani-ledger
   kubectl get svc -n vani-ledger
    kubectl logs -f deployment/voice-ai-service -n vani-ledger
 kubectl logs -f deployment/ledger-service -n vani-ledger
