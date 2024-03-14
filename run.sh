#!/bin/bash

# create k8s cluster
k3d cluster create --config ./k3d.yaml

# Update helm dependencies and deploy chart to a cluster
helm dependency update helm
helm install go-store-things helm

# wait for the pods to start
kubectl wait pods -l app.kubernetes.io/instance=go-store-things --for=condition=Ready
