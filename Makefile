CONTAINER_NAME := web-crawler
DOCKER_TAG := web-crawler

run: image
	docker run -d --name $(CONTAINER_NAME) $(DOCKER_TAG)

stop:
	docker kill --signal=SIGTSTP $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

image:
	docker build . -t $(DOCKER_TAG)

logs:
	docker logs -f $(CONTAINER_NAME)