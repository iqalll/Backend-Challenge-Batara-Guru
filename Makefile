.PHONY: build clean validate api run

swagger: validate clean build

validate:
	swagger validate ./api/swagger.yml

build:
	swagger -q generate server -A batara -f api/swagger.yml -s api -m models
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v ./cmd/batara-server

clean:
	rm -rf batara-server
	go clean -i .

api: validate clean build
	./batara-server --port=8080 --host=0.0.0.0

run: swagger
	./batara-server --port=8080 --host=0.0.0.0