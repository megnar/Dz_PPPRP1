#!/bin/bash

docker build -t your-docker-registry/microservice-a:latest -f Dockerfile.microservice-a .
docker build -t your-docker-registry/microservice-b:latest -f Dockerfile.microservice-b .

kubectl apply -f microservice-a.yaml
kubectl apply -f microservice-b.yaml

kubectl apply -f egress-gateway.yaml
