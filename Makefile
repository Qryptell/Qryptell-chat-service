build:
	go build -o ./bin/chat ./cmd/

run: build
	./bin/chat
