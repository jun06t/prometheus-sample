.PHONY: prepare build

REGISTRY_DOMAIN := jun06t

build:
	docker build -f Dockerfile-gateway -t ${REGISTRY_DOMAIN}/grpc-gateway-prometheus:latest .
	docker build -f Dockerfile-grpc -t ${REGISTRY_DOMAIN}/grpc-backend-prometheus:latest .

push:
	docker push ${REGISTRY_DOMAIN}/grpc-gateway-prometheus
	docker push ${REGISTRY_DOMAIN}/grpc-backend-prometheus
