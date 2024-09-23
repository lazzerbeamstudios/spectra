# go
go-run-dev:
	go run api.go --env=dev --server=api

# ent
ent-generate:
	go get entgo.io/ent/cmd/internal/printer@v0.14.0
	go generate ./ent
	go mod tidy

# atlas
atlas-diff-dev:
	atlas migrate diff --env dev

atlas-apply-dev:
	atlas migrate apply --env dev

atlas-dev:
	$(MAKE) atlas-diff-dev
	$(MAKE) atlas-apply-dev

atlas-diff-stag:
	atlas migrate diff --env stag

atlas-apply-stag:
	atlas migrate apply --env stag

atlas-stag:
	$(MAKE) atlas-diff-stag
	$(MAKE) atlas-apply-stag

atlas-diff-prod:
	atlas migrate diff --env prod

atlas-apply-prod:
	atlas migrate apply --env prod

atlas-prod:
	$(MAKE) atlas-diff-prod
	$(MAKE) atlas-apply-prod
