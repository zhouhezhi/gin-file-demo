## build: Build source code for host platform
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

.PHONY: run	
run:
	go run main.go

