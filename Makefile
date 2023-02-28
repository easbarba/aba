.DEFAULT_GOAL := tests

NAME := aba
MAIN := ./main.go
DEST := ${HOME}/.local/bin

deps:
	go mod download
	go mod verify

fmt:
	go build -o ${HOME}/.local/bin/${NAME} ./cmd/${NAME}/main.go

lint:
	golint ./...

vet:
	go vet ./...

test:
	go test ./...

imports:
	goimports -l -w .

clean:
	rm ${NAME}
	go mod tidy

watch:
	CompileDaemon --build="go build -o ./${NAME} ${MAIN}" --command="./${NAME}"

img:
	podman build --file ./Dockerfile --tag ${USER}/${NAME}:$(shell cat .version)

.PHONY: img clean imports test vet lint fmt deps watch
