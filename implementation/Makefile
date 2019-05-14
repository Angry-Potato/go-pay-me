# Shamelessly taken from https://sohlich.github.io/post/go_makefile/
GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEPENSURE=$(DEPCMD) ensure
IMAGE_NAME=go-pay-me
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix
		
all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v -tags=unit ./...
test-integration: 
	$(GOTEST) -v -tags=integration ./... 
test-async:
	$(GOTEST) -v -tags="unit integration" ./... 
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run: build
	./$(BINARY_NAME)
deps:
	$(DEPENSURE)

# Docker
build-docker-image:
	docker build --force-rm --pull --tag $(IMAGE_NAME):latest .

run-docker-image: build-docker-image
	docker run --rm $(IMAGE_NAME):latest

docker-compose-up:
	docker-compose up --force-recreate --always-recreate-deps --build -V --remove-orphans

docker-compose-test:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --force-recreate --always-recreate-deps --build -V --remove-orphans --exit-code-from $(IMAGE_NAME)-tests