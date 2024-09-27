APP_NAME=card-validator
MAIN=main.go
BIN_DIR=bin

run:
	go run $(MAIN)

build:
	mkdir $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) $(MAIN)

install:
	go install

clean:
	rm -rf bin