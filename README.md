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
1. Run the bash script below to spin up the k8s cluster locally, and deploy service to it:  
```
run.sh
```
> There is no CI step as the Docker image is getting built and published by this job https://github.com/Swif2er/go-store-things/actions/workflows/ci.yml and the package is available here: https://github.com/Swif2er/go-store-things/pkgs/container/go-store-things
> Script will also verify the service and Redis are up and ready.

2. Get the CLI service pod name
```
CLI_POD_NAME=$(kubectl get pod -l app.kubernetes.io/instance=go-store-things -o jsonpath="{.items[0].metadata.name}")
```
3. Run CLI commands:
```
kubectl exec -it $CLI_POD_NAME -- ./go-store-things ping
kubectl exec -it $CLI_POD_NAME -- ./go-store-things set -k one -v one
kubectl exec -it $CLI_POD_NAME -- ./go-store-things get -k one
```

4. Destroy Kubernetes cluster:
```
k3d cluster delete test-cluster
```
