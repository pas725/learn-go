
# Sample run
.PHONY: run
run:
	go run pkg/06_pointers/pointers.go

# Sample build
build:
	go build -o ./cmd/ pkg/06_pointers/pointers.go

# Creates a sample module
mod-init:
	go mod init sample-mod