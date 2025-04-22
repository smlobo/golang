
TARGET = ssa-build

build:
	go build -o bin/${TARGET} main.go

module:
	rm -f go.mod go.sum
	go mod init ${TARGET}
	go mod tidy

clean:
	rm -rf bin go.mod go.sum

