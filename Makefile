

test:
	go test ./...

deploy:
	cp main /Users/bigmac/Bin/hello_world

# installs mocktail into $GOPATH/bin
mock_install:
	go install github.com/traefik/mocktail@latest

