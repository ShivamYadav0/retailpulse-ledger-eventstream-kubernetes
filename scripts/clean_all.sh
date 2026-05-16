#!/bin/bash
echo "Stopping and removing all containers..."

echo "Removing all k3d clusters..."
k3d cluster delete --all

echo "Stopping and removing all containers..."
docker rm -f $(docker ps -aq)

echo "Removing all Docker images..."
docker rmi -f $(docker images -aq)

echo "Removing all Docker volumes..."
docker volume rm $(docker volume ls -q)

echo "Pruning all Docker networks..."
docker network prune -f

echo "Removing all k3d clusters..."
k3d cluster delete --all

echo "Full cleanup complete!"
