#!/bin/bash

set -e

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
KAFKA_NAMESPACE="kafka"

echo "📦 Running PostgreSQL via Docker Compose..."
cd "$PROJECT_DIR"
docker compose up -d db

echo "⏳ Waiting for PostgreSQL to start..."
until docker exec location-db pg_isready -U user -d location; do
  sleep 2
done
echo "✅ PostgreSQL is ready!"

echo "🚀 Launching Minikube..."
minikube start --memory=8192 --cpus=4 --disk-size=50g

echo "📦 Install the Strimzi Kafka Operator (if it isn't already installed)..."
kubectl create namespace "$KAFKA_NAMESPACE" 2>/dev/null || true
kubectl apply -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka

echo "🛠️  Apply manifests Kafka and UI..."
kubectl apply -f "$PROJECT_DIR/kafka-cluster.yaml" -n kafka
kubectl apply -f "$PROJECT_DIR/kafka-ui.yaml" -n kafka

echo "⏳ Waiting for all Kafka pods to start..."
kubectl wait --for=condition=Ready pod -l strimzi.io/cluster=my-cluster -n kafka --timeout=300s
kubectl wait --for=condition=Ready pod -l app=kafka-ui -n kafka --timeout=120s

echo "🔗 Forwarding ports in the background..."

# Kafka (for Go-backend)
kubectl port-forward -n kafka svc/my-cluster-kafka-external-bootstrap 9094:9094 >/dev/null 2>&1 &
KAFKA_PORT_FORWARD_PID=$!

# Kafka UI (for browser)
kubectl port-forward -n kafka svc/kafka-ui 8081:8080 >/dev/null 2>&1 &
UI_PORT_FORWARD_PID=$!

echo "✅ Ports are forwarded:"
echo "   - Kafka:  localhost:9094"
echo "   - Kafka UI: http://localhost:8081"

echo "🚀 Launching the Go backend..."
cd "$PROJECT_DIR"
go run cmd/location-service/main.go

# When finished, stop background processes
trap '{
  echo "🛑 I stop everything...";
  kill $KAFKA_PORT_FORWARD_PID $UI_PORT_FORWARD_PID 2>/dev/null
  docker compose down
}' EXIT