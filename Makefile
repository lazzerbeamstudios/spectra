# docker
docker-login:
	docker login ghcr.io

docker-logout:
	docker logout ghcr.io

docker-network:
	docker network create docker-network --driver=bridge --attachable

docker-postgres:
	docker compose -f docker-postgres.yaml up -d

docker-valkey:
	docker compose -f docker-valkey.yaml up -d

docker-api:
	docker compose -f docker-api.yaml up

docker-run:
	$(MAKE) docker-network
	$(MAKE) docker-postgres
	$(MAKE) docker-valkey
	$(MAKE) docker-api

docker-tag:
	sh docker_tag.sh

docker-buildx:
	sh docker_buildx.sh

# helm
helm-install-dev:
	helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/dev.yaml

helm-upgrade-dev:
	helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/dev.yaml

helm-install-stag:
	helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/stag.yaml

helm-upgrade-stag:
	helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/stag.yaml

helm-install-prod:
	helm install api helm-v1 --values helm-v1/values.yaml -f helm-v1/prod.yaml

helm-upgrade-prod:
	helm upgrade api helm-v1 --values helm-v1/values.yaml -f helm-v1/prod.yaml

helm-uninstall:
	helm uninstall api

# gcp
gcp-auth-login:
	gcloud auth login

gcp-auth-revoke:
	gcloud auth revoke

gcp-auth-login-terraform:
	gcloud auth application-default login

gcp-auth-revoke-terraform:
	gcloud auth application-default revoke

gcp-create-vpc:
	gcloud compute networks create vpc-v1 --subnet-mode=auto --project [project]

gcp-delete-vpc:
	gcloud compute networks delete vpc-v1 --quiet --project [project]

# argocd
argocd-create:
	kubectl create namespace api-stag
	kubectl create namespace api-prod
	kubectl create namespace argocd
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

argocd-port:
	kubectl port-forward svc/argocd-server -n argocd 8080:443

argocd-password:
	argocd admin initial-password -n argocd
