NOW=`date +%Y%m%d%H%M%S`
OS=`uname -n -m`
AFTER_COMMIT=`git rev-parse HEAD`
GOPATH_DIR=`go env GOPATH`

install:
	go install -ldflags "-X 'main.BuildTime=$(NOW)' -X 'main.BuildOSUname=$(OS)' -X 'main.BuildCommit=$(AFTER_COMMIT)'" .

check:
	@echo "running tests..."
	@go test -count 1 -coverprofile=coverage.txt -covermode=atomic -race -v ./src/...
	@echo "everything is OK"

.PHONY: check
