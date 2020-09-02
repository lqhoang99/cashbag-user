# cashbag-me-mini

# How to run docker compose up
docker-compose up --build

# Check test coverage
go test ./...  -coverprofile=coverage.out
#
protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto