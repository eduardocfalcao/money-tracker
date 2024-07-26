run:
	go run cmd/api/main.go

sqlc:
	sqlc generate -f ./database/sqlc.yaml

new-migration: 
	migrate create -ext sql -dir ./database/migrations -seq ${name}

db_url ?= "postgres://postgres:12345678a@localhost:5433/money-tracker?sslmode=disable"
run-migrations:
	migrate -verbose -database ${db_url} -path ./database/migrations up

start-postgres:
	docker start money-tracker

create-postgres-container:
	docker run --name money-tracker -e POSTGRES_PASSWORD=12345678a -e POSTGRES_DB=money-tracker -p 5433:5432 -v ./.pgdata:/var/lib/postgresql/data -d postgres:15

drop-database:
	migrate -database ${db_url} -path ./database/migrations drop

down-database:
	migrate -database ${db_url} -path ./database/migrations down ${n}

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
	