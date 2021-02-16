BINARY_SERVER=server
BINARY_CLIENT=client
test:
	go test -v -cover -covermode=atomic ./...

client:
	go build -o ${BINARY_CLIENT} cmd/client/main.go

server:
	go build -o ${BINARY_SERVER} cmd/server/main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY_CLIENT} ] ; then rm ${BINARY_CLIENT} ; fi
	if [ -f ${BINARY_SERVER} ] ; then rm ${BINARY_SERVER} ; fi

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

remove:
	docker container rm postgres

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint