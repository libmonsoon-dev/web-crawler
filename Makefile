GO := go
CONTAINER_NAME := web-crawler
DOCKER_TAG := web-crawler
LIB_TOR_VERSION := v1.0.380
LIB_TOR_PATH := $(shell $(GO) env GOPATH)/pkg/mod/github.com/ipsn/go-libtor@$(LIB_TOR_VERSION)

PG_HOST := localhost
PG_PORT := 5432
PG_USER := postgres
PG_PASS := devpass
PG_DB := postgres

MIGRATION_NAME := ""


run: image
	docker run -d --name $(CONTAINER_NAME) $(DOCKER_TAG)

stop:
	docker kill --signal=SIGTSTP $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

image:
	docker build . -t $(DOCKER_TAG)

logs:
	docker logs -f $(CONTAINER_NAME)

test:
	TEST_PG_CONNECTION="postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable" \
	  $(GO) test -v -race ./...

dev-db:
	docker run \
		--detach \
		--name web-crawler-dev \
		--env POSTGRES_DB=$(PG_DB) \
		--env POSTGRES_USER=$(PG_USER) \
		--env POSTGRES_PASSWORD=$(PG_PASS) \
		--publish $(PG_PORT):5432 \
		--volume /var/lib/postgresql/web-crawler/data:/var/lib/postgresql/data \
		postgres:13.3 || echo

migration:
	if [ -z $(MIGRATION_NAME) ]; then echo "Usage: make migration MIGRATION_NAME=..."; exit 1; fi
	migrate create -ext sql -dir ./storage/sql/postgresql/migration -seq $(MIGRATION_NAME)

jet-pg:
	jet -source=PostgreSQL \
		-host=$(PG_HOST) \
		-port=$(PG_PORT) \
		-user=$(PG_USER) \
		-password=$(PG_PASS) \
		-dbname=$(PG_DB) \
		-path=./storage/sql/query
	rm -rf storage/sql/query/postgres/public/model

dependency:
	$(GO) mod tidy
	$(GO) mod vendor
	make vendor/github.com/ipsn/go-libtor/build \
		vendor/github.com/ipsn/go-libtor/config \
		vendor/github.com/ipsn/go-libtor/libevent \
		vendor/github.com/ipsn/go-libtor/libevent_config \
		vendor/github.com/ipsn/go-libtor/openssl \
		vendor/github.com/ipsn/go-libtor/openssl_config \
		vendor/github.com/ipsn/go-libtor/tor \
		vendor/github.com/ipsn/go-libtor/tor_config \
		vendor/github.com/ipsn/go-libtor/zlib

vendor/github.com/ipsn/go-libtor/build:
	cp -r $(LIB_TOR_PATH)/build vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/config:
	cp -r $(LIB_TOR_PATH)/config vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/libevent:
	cp -r $(LIB_TOR_PATH)/libevent vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/libevent_config:
	cp -r $(LIB_TOR_PATH)/libevent_config vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/openssl:
	cp -r $(LIB_TOR_PATH)/openssl vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/openssl_config:
	cp -r $(LIB_TOR_PATH)/openssl_config vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/tor:
	cp -r $(LIB_TOR_PATH)/tor vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/tor_config:
	cp -r $(LIB_TOR_PATH)/tor_config vendor/github.com/ipsn/go-libtor

vendor/github.com/ipsn/go-libtor/zlib:
	cp -r $(LIB_TOR_PATH)/zlib vendor/github.com/ipsn/go-libtor
