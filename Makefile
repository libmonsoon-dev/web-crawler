GO := go
CONTAINER_NAME := web-crawler
DOCKER_TAG := web-crawler
LIB_TOR_VERSION := v1.0.380
LIB_TOR_PATH := $(shell $(GO) env GOPATH)/pkg/mod/github.com/ipsn/go-libtor@$(LIB_TOR_VERSION)

run: image
	docker run -d --name $(CONTAINER_NAME) $(DOCKER_TAG)

stop:
	docker kill --signal=SIGTSTP $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

image:
	docker build . -t $(DOCKER_TAG)

logs:
	docker logs -f $(CONTAINER_NAME)

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
