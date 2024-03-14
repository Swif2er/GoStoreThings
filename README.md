# go-store-things

A simple Golang CLI tool that reads and writes key-value pairs to a Redis storage. 

There is a CI build using GitHub actions to generate a Docker image and publish it to a publicly available repo in GHCR. This git repo also includes a Helm chart to deploy the CLI tool and a single instance of Redis (configured as Helm chart dependency and uses the official Bitnami helm chart). There is also a k3d configuration file that can be used to spin up the local Kubernetes cluster.

## Prerequisites:
- Docker 25.0.3
- k3d v5.6.0
- Kubectl v1.29.1
- Helm v3.14.2
> Older versions should work too, but it was tested using these versions.

## Deployment 
1. Build the app, spin up the k8s cluster, and deploy to it:  
```
run.sh
```

