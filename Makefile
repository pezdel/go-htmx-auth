.PHONY: run js

run:
	templ generate
	go run cmd/api/main.go

js:
	cd js && npm run build-prod
