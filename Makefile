bin:
	CGO_ENABLE=0 go build -o bin/fw cmd/fw/fw.go
images:
	docker build -f build/Dockerfile -t fw/test . 
vet:
	go vet ./...
