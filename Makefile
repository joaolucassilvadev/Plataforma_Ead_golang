run-grpc-server:
	go run cmd/grpc-server/main.go
run-service-core-tests:
   cd service-core && go test ./...