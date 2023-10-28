# Go parameters
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
CMD_MAIN = ./cmd/main.go
START = $(GO) run $(CMD_MAIN)
INFRA_CONTRACTS = ./internal/infra/contracts
EXTERNAL_CONTRACTS = ./external
MOCKS = ./external/mocks

# Output binary name
BINARY_NAME=myapp

#
start:
	echo "ok"

# Target for building the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Target for running tests
test:
	$(GOTEST) ./...

# Target for cleaning up the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Target for installing project dependencies
get:
	$(GOGET)

generate-mocks:
	mockgen -source=$(INFRA_CONTRACTS)/transaction.contract.go -destination=$(MOCKS)/transaction.repository.go -package=mocks
	mockgen -source=$(INFRA_CONTRACTS)/user.contract.go -destination=$(MOCKS)/user.repository.go -package=mocks
	mockgen -source=$(INFRA_CONTRACTS)/operation.contract.go -destination=$(MOCKS)/operation.repository.go -package=mocks
	mockgen -source=$(INFRA_CONTRACTS)/notificationMailer.contract.go -destination=$(MOCKS)/notificationMailer.repository.go -package=mocks
	mockgen -source=$(EXTERNAL_CONTRACTS)/configs/queue/producer.queue.go -destination=$(MOCKS)/producer.queue.go -package=mocks




# Default target when you just run 'make'
default: build

