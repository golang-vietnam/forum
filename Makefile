default: autoReload

autoReload:
	@export PATH=$(PATH):$(GOPATH)/bin;gin -a 8080 run

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
	@go get github.com/codegangsta/gin
	@go get github.com/smartystreets/goconvey
	@export PATH=$(PATH):$(GOPATH)/bin;godep restore
	@echo "Install dependencies successful!"

save:
	@echo "Saving..."
	@go get
	@go get github.com/kr/godep	
	@export PATH=$(PATH):$(GOPATH)/bin;godep save
	@echo "Save new dependencies successful!"

	