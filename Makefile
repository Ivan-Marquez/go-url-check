build:
	export GO111MODULE=on
	go build -o bin/go-url-check main.go fetch.go
clean:
	rm -rf ./bin