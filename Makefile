GO:=go
LDFLAGS:=-s -w
ORDER_BIN:=order-server
ORDER_BIN_UNIX:=order-server-linux
WRITER_BIN:=writer
WRITER_BIN_UNIX:=writer-linux
OUT_PATH:=./.out
LINUX_ENV:=GOARCH=amd64 GOOS=linux CGO_ENABLED=0

.PHONY: migrate clean test up down

all: order order-linux writer writer-linux

order: 
	@$(GO) build -ldflags '$(LDFLAGS)' -o $(OUT_PATH)/$(ORDER_BIN) ./cmd/order

order-linux: 
	@$(LINUX_ENV) $(GO) build -ldflags '$(LDFLAGS)' -o $(OUT_PATH)/$(ORDER_BIN_UNIX) ./cmd/order

writer:
	@$(GO) build -ldflags '$(LDFLAGS)' -o $(OUT_PATH)/$(WRITER_BIN) ./cmd/writer

writer-linux: 
	@$(LINUX_ENV) $(GO) build -ldflags '$(LDFLAGS)' -o $(OUT_PATH)/$(WRITER_BIN_UNIX) ./cmd/writer

clean: 
	@rm -rf .out

generate: 
	@$(GO) generate ./internal/...
	@$(GO) generate ./migrate

generate-ent:
	@$(GO) generate ./ent

migrate:
	@$(GO) run ./cmd/migrate

test:
	@$(GO) test -v ./...

up:
	@docker compose up -d

down:
	@docker compose down