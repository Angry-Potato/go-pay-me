# Shamelessly based on https://sohlich.github.io/post/go_makefile/
GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEPENSURE=$(DEPCMD) ensure
IMAGE_NAME=go-pay-me
BINARY_NAME=app
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_ARGS=-v ./...
COMPOSE_PARANOID_ARGS=--force-recreate --always-recreate-deps --build -V --remove-orphans

all: deps test build
build: deps
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -tags=unit $(TEST_ARGS)
test-integration: 
	$(GOTEST) -tags=integration $(TEST_ARGS)
test-acceptance: 
	$(GOTEST) -tags=acceptance $(TEST_ARGS)
test-async:
	$(GOTEST) -tags="unit integration" $(TEST_ARGS)
test-sync: test-acceptance
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
	docker-compose up $(COMPOSE_PARANOID_ARGS)

docker-compose-test:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up $(COMPOSE_PARANOID_ARGS) --exit-code-from $(IMAGE_NAME)-tests