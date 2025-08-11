

build:
	go build -o mcp-server cmd/subsystems/main.go

test:
	go test ./...

deploy:
	cp mcp-server /Users/bigmac/Bin/mcp-server

# installs mocktail into $GOPATH/bin
mock_install:
	go install github.com/traefik/mocktail@latest

