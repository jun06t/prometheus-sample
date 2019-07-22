.PHONY: destroy apply apply-deployments apply-services apply-configmap

apply: apply-configmap apply-deployments apply-services

apply-configmap:
	kubectl apply -f configmap-prometheus.yml

apply-deployments:
	kubectl apply -f deployment-backend.yml
	kubectl apply -f deployment-gateway.yml
	kubectl apply -f deployment-grafana.yml
	kubectl apply -f deployment-prometheus.yml

apply-services:
	kubectl apply -f service-backend.yml
	kubectl apply -f service-gateway.yml
	kubectl apply -f service-grafana.yml
	kubectl apply -f service-prometheus.yml

destroy:
	kubectl delete deployment gateway-deployment
	kubectl delete deployment backend-deployment
	kubectl delete deployment grafana-deployment
	kubectl delete deployment prometheus-deployment
	kubectl delete service gateway-svc
	kubectl delete service backend-svc
	kubectl delete service grafana-svc
	kubectl delete service prometheus-svc
	kubectl delete service prometheus-proxy
	kubectl delete configmap prometheus-config

