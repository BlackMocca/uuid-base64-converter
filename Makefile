VERSION=1.0.0

build-os:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-linux-amd64-$(VERSION)  ./cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-linux-arm64-$(VERSION)  ./cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-darwin-amd64-$(VERSION)  ./cmd/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-darwin-arm64-$(VERSION)  ./cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-windows-386-$(VERSION)  ./cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-X main.Version=$(VERSION)' -o build/ub64c-windows-amd64-$(VERSION)  ./cmd/main.go