# Kubernetes setup (Ubuntu + k3s)

This document describes how to deploy a local Kubernetes cluster on Ubuntu
using **k3s**, and how to run **Kafka (Strimzi Operator)** and **Kafka UI**
for development purposes.

The setup was tested on a single-node cluster.

---

## Environment

- OS: Ubuntu 22.04 LTS
- Kubernetes: k3s
- Container runtime: containerd
- Kafka operator: Strimzi
- Kafka UI: provectuslabs/kafka-ui

---

## 1. Server access (SSH)

All commands below can be executed directly on the server
or via SSH from a local machine.

Example:
```bash
ssh user@SERVER_IP
```
## 2. Install k3s

Install k3s using the official installation script:
```bash
curl -sfL https://get.k3s.io | sh -
```
Check that k3s is running:
```bash
sudo systemctl status k3s
```
## 3. Configure kubectl
k3s installs kubectl automatically.
To use kubectl without sudo:
```bash
sudo chmod 644 /etc/rancher/k3s/k3s.yaml
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
```
Verify cluster status:
```bash
kubectl get nodes
```
Expected result:

1 node

roles: control-plane,master

## 4. Create namespace for Kafka
Kafka and related components are deployed into a dedicated namespace:
```bash
kubectl create namespace kafka-dev
```
## 5. Install Strimzi Operator
Apply Strimzi operator manifests:
```bash
kubectl apply -n kafka-dev -f https://strimzi.io/install/latest?namespace=kafka-dev
```
Verify operator is running:
```bash
kubectl get pods -n kafka-dev
```
Expected:
```bash
strimzi-cluster-operator pod in Running state
```
## 7. Expose Kafka externally (NodePort)
Kafka is exposed via NodePort for local development.

Check services:
```bash
kubectl get svc -n kafka-dev
```
Look for:
```bash
my-cluster-kafka-external-bootstrap
```
Example:
```bash
30092:32473/TCP
```
Kafka bootstrap address:
```bash
SERVER_IP:32473
```
## 8. Deploy Kafka UI
Apply Kafka UI manifest:
```bash
kubectl apply -f kafka-ui-dev.yaml
```
Check service:
```bash
kubectl get svc -n kafka-dev
```
Example:
```bash
8080:30080/TCP
```
Kafka UI URL:
```bash
http://SERVER_IP:30080
```
## 9. Access via SSH tunnel (optional)
If ports are not directly accessible, use SSH port forwarding:
```bash
ssh -L 30080:localhost:30080 user@SERVER_IP
```
Then open:
```bash
http://localhost:30080
```

## 10. Application configuration
The application supports switching Kafka brokers via environment variables.

Example .env:
```bash
APP_ENV=k8s
KAFKA_BROKERS_K8S=SERVER_IP:32473
KAFKA_BROKERS_MINIKUBE=localhost:9094
```
Kafka brokers are selected based on APP_ENV.

# Notes

This setup is intended for development and testing

Single-node k3s cluster

No TLS or authentication enabled for Kafka

