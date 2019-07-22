.PHONY: destroy apply apply-deployments apply-services

apply: apply-deployments apply-services

apply-deployments:
	kubectl apply -f deployment-backend.yml
	kubectl apply -f deployment-gateway.yml
	kubectl apply -f deployment-grafana.yml

apply-services:
	kubectl apply -f service-backend.yml
	kubectl apply -f service-gateway.yml
	kubectl apply -f service-grafana.yml

destroy:
	kubectl delete deployment gateway-deployment
	kubectl delete deployment backend-deployment
	kubectl delete deployment grafana-deployment
	kubectl delete service gateway-svc
	kubectl delete service backend-svc
	kubectl delete service grafana-svc

