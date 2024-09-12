DC = docker-compose
EXEC = docker exec -it
LOGS = docker logs

APP_FILE = docker-compose.yml
APP_CONTAINER = gowhisper

BINARY_NAME = gowhisper
BUILD_DIR = ./dist


.PHONY: app app-down app-shell app-logs app-reset
app:
	${DC} -f ${APP_FILE} up --build -d 

app-down:
	${DC} -f ${APP_FILE} down

app-shell:
	${EXEC} ${APP_CONTAINER} bash

app-logs:
	${LOGS} ${APP_CONTAINER} -f

app-reset:
	${DC} -f ${APP_FILE} down
	git reset --hard HEAD~1
	${DC} -f ${APP_FILE} up --build -d


.PHONY: build-all
build-all: build-linux build-mac build-windows build-arm

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY_NAME}-linux ./cmd/gowhisper

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY_NAME}-mac ./cmd/gowhisper

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe ./cmd/gowhisper

build-arm:
	GOOS=linux GOARCH=arm64 go build -o ${BUILD_DIR}/${BINARY_NAME}-arm ./cmd/gowhisper
