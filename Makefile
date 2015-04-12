default: run

run:
	go run main.go

install:
	go get github.com/kr/godep
	export PATH=$(PATH):$(GOPATH)/bin;godep restore
	@echo "Install dependencies successful"

	