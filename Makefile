proto:
	Remove-Item -Path .\pb\* -Recurse
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative `
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative `
	proto\*.proto

server_local:
	go run main.go

image:
	docker image rm chord:latest
	docker build -t chord:latest .
	
.PHONY: proto server_local protowin