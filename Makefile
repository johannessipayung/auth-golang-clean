test:
	go test ./... -v

test-race:
	go test -race ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run

security:
	gosec ./...

vuln:
	govulncheck ./...