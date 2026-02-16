.PHONY: openapi

openapi:
	@swag init -g internal/core/http/http.go --output api --outputTypes json,yaml --v3.1
ifneq ($(OS),Windows_NT)
	@mv api/swagger.json api/openapi.json
	@mv api/swagger.yaml api/openapi.yml
else
	@powershell -Command "Move-Item api/swagger.json api/openapi.json -Force"
	@powershell -Command "Move-Item api/swagger.yaml api/openapi.yml -Force"
endif
