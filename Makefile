default: run

run:
	@go run forum.go start

test:
	@go run forum.go test

install:
	@echo "Installing ..."
	@go get github.com/kr/godep
	@export PATH=$(PATH):$(GOPATH)/bin;godep restore
	@echo "Install dependencies successful !"

save:
	@echo "Saving ..."
	@go get github.com/kr/godep
	@export PATH=$(PATH):$(GOPATH)/bin;godep save
	@echo "Save new dependencies successful !"

	