proto:
	rm -rf pb/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

server_local:
	go run main.go

server_docker:
	docker image rm chord:latest
	docker build -t chord:latest .
	docker run --name chord1 -p 9090:9090 -e CHORD_ID=1 -e SUCCESSOR_PORT=localhost:9091 -e SERVER_ADDRESS=localhost:9090 chord:latest
	docker run --name chord2 -p 9091:9091 -e CHORD_ID=2 -e SUCCESSOR_PORT=localhost:9092 -e SERVER_ADDRESS=localhost:9091 chord:latest
	docker run --name chord3 -p 9092:9092 -e CHORD_ID=3 -e SUCCESSOR_PORT=localhost:9090 -e SERVER_ADDRESS=localhost:9092 chord:latest
.PHONY: proto server_local