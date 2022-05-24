lint:
	golangci-lint run
test:
	cp .env internal/api/delivery/
	cp .env internal/middleware/
	go test -coverpkg=./... -coverprofile cover.out.tmp ./...
	cat cover.out.tmp grep -v "monitoring" | grep -v "easyjson" | grep -v "mock_*" | grep -v ".pb.go" | grep -v ".pb" | grep -v "middleware.go" | grep -v "/cmd*"> cover.out
	go tool cover -func cover.out