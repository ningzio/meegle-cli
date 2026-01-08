run:
	go run ./cmd/meegle-tui

lint:
	golangci-lint run ./...
