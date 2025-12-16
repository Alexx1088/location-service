#!/bin/bash

set -e

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
KAFKA_NAMESPACE="kafka"

echo "📦 Запускаю PostgreSQL через Docker Compose..."
cd "$PROJECT_DIR"
docker compose up -d db

echo "⏳ Жду, пока PostgreSQL запустится..."
until docker exec location-db pg_isready -U user -d location; do
  sleep 2
done
echo "✅ PostgreSQL готова!"

echo "🚀 Запускаю Minikube..."
minikube start --memory=8192 --cpus=4 --disk-size=50g

echo "📦 Устанавливаю Strimzi Kafka Operator (если ещё не установлен)..."
kubectl create namespace "$KAFKA_NAMESPACE" 2>/dev/null || true
kubectl apply -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka

echo "🛠️  Применяю манифесты Kafka и UI..."
kubectl apply -f "$PROJECT_DIR/kafka-cluster.yaml" -n kafka
kubectl apply -f "$PROJECT_DIR/kafka-ui.yaml" -n kafka

echo "⏳ Жду, пока все поды Kafka запустятся..."
kubectl wait --for=condition=Ready pod -l strimzi.io/cluster=my-cluster -n kafka --timeout=300s
kubectl wait --for=condition=Ready pod -l app=kafka-ui -n kafka --timeout=120s

echo "🔗 Пробрасываю порты в фоне..."

# Kafka (для Go-бэкенда)
kubectl port-forward -n kafka svc/my-cluster-kafka-external-bootstrap 9094:9094 >/dev/null 2>&1 &
KAFKA_PORT_FORWARD_PID=$!

# Kafka UI (для браузера)
kubectl port-forward -n kafka svc/kafka-ui 8081:8080 >/dev/null 2>&1 &
UI_PORT_FORWARD_PID=$!

echo "✅ Порты проброшены:"
echo "   - Kafka:  localhost:9094"
echo "   - Kafka UI: http://localhost:8081"

echo "🚀 Запускаю Go-бэкенд..."
cd "$PROJECT_DIR"
go run cmd/location-service/main.go

# При завершении остановим фоновые процессы
trap '{
  echo "🛑 Останавливаю всё...";
  kill $KAFKA_PORT_FORWARD_PID $UI_PORT_FORWARD_PID 2>/dev/null
  docker compose down
}' EXIT