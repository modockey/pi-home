swagger-generate:
	swagger generate server -t gen -f ../swagger.yaml

run:
	go run ./gen/cmd/photo-award-server/main.go --port 8000

build:
	go build -o ./target/photo-award-server.go ./gen/cmd/photo-award-server/main.go 