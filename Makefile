BIN=./bin
FILE=./bin/main
SRC=./cmd/main/main.go

all: fmt clean build run
build:
	$(info ************ Building ************)
	go build -o $(FILE) $(SRC)
run:
	$(info ************ Running ************)
	$(FILE)
clean:
	$(info ************ Cleaning ************)
	rm -rf $(BIN)
fmt:
	$(info ************ Formatting ************)
	go fmt ./...
