
.PHONY: get-mockery
get-mockery:
	@go install github.com/vektra/mockery/v2@v2.38.0

.PHONY: mocks
mocks:
	@mockery --all --keeptree
