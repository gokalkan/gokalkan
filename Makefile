.PHONY: lint
lint:
	@command -v golangci-lint > /dev/null 2>&1 || (cd $${TMPDIR} && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2)
	golangci-lint run --config .golangci.yaml

.PHONY: fmtimport
fmtimport:
	golangci-lint run -E goimports --fix

.PHONY: test
test:
	@go clean -testcache
	@go test -v -run . ./internal/testdata/test/

.PHONY: wtest
wtest:
	@go clean -testcache
	@go test -v -run . ./internal/testdata/wtest/

.PHONY: testprod
testprod:
	@go clean -testcache
	@go test -v -run . ./internal/testdata/prod/
