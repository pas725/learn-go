
# Sample run
.PHONY: run
run:
	go run pkg/06_pointers/pointers.go

# Sample build
build:
	go build -o ./cmd/ pkg/06_pointers/pointers.go