GO ?= go

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	$(GO) vet ./...

.PHONY: test
test: test.proto
	$(GO) test -race ./...

.PHONY: bench
bench: test.proto
	$(GO) test -bench . ./... -benchmem

.PHONY: test.proto
test.proto: test/bin/protoc-gen-go-log-valuer
	protoc --go_out=test/ --plugin=test/bin/protoc-gen-go-log-valuer --go-log-valuer_out=test/ test/*.proto
	mv test/github.com/kei2100/protoc-gen-go-log-valuer/test/*.go test/
	rm -r test/github.com

test/bin/protoc-gen-go-log-valuer: $(shell find plugin -name '*.go') go.sum
	$(GO) build -ldflags="-s -w" -trimpath -o $@ ./plugin/protoc-gen-go-log-valuer
