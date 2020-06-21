deploy-cluster:
	kubectl apply -f ./.k8s/ingress-deployment.yaml
	kubectl apply -f ./.k8s/ingress-service.yaml
	$(MAKE) update-configmap
	kubectl apply -f ./.k8s/deployment.yml
	kubectl apply -f ./.k8s/ingress-resources.yaml

update-configmap:
	kubectl create configmap --dry-run -o yaml --from-env-file=.env app-settings | kubectl apply -f -

build-image:
ifdef tag
	docker build -t crfalcao.azurecr.io/money-tracker:$(tag) -t crfalcao.azurecr.io/money-tracker:latest .
else
	docker build -t crfalcao.azurecr.io/money-tracker:latest .
endif

get-next-version:
	$(eval lastVersionString=$(shell sh -c "docker images \"crfalcao.azurecr.io/money-tracker:v*\" --format \"{{ .Tag }}\" | head -1 | cut -c 2-"))
	$(eval lastVersionString=$(shell echo $$(($(lastVersionString)+1))))
	@echo $(lastVersionString)

rollout-update: 
ifdef version
build-image tag=$(version)
kubectl set image deployments/money-tracker-deployment money-tracker-app=crfalcao.azurecr.io/money-tracker:$(version)
endif

rollout-update-v2: 
	$(eval lastVersionString=$(shell sh -c "docker images \"crfalcao.azurecr.io/money-tracker:v*\" --format \"{{ .Tag }}\" | head -1 | cut -c 2-"))
	$(eval version=$(shell echo $$(($(lastVersionString)+1))))
	$(eval version=$(shell echo v$(version)))

	@echo "Updating image to image crfalcao.azurecr.io/money-tracker:$(version)"
	$(MAKE) build-image tag=$(version)
	kubectl set image deployments/money-tracker-deployment money-tracker-app=crfalcao.azurecr.io/money-tracker:$(version)
	