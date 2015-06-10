default: run

run:
	@go run forum.go start

testOnWeb:
	@cd $(GOPATH)/src/github.com/golang-vietnam/forum/tests;$(GOPATH)/bin/goconvey

test:
	@echo "Begin test on console!"
	@cd $(GOPATH)/src/github.com/golang-vietnam/forum/tests;go test -v
	@echo "End test!"

runOnTest:
	@go run forum.go test

install:
	@echo "Installing..."
	@go get github.com/kr/godep
	@export PATH=$(PATH):$(GOPATH)/bin;godep restore
	@echo "Install dependencies successful!"

save:
	@echo "Saving..."
	@go get github.com/kr/godep
	@export PATH=$(PATH):$(GOPATH)/bin;godep save
	@echo "Save new dependencies successful!"

	