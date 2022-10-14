SHELL=/bin/bash
IMAGE=localhost/kevydotvinu/curl-paste
ENGINE?=podman
NAME=curl-paste

.PHONY: all
all: build kill run logs

.PHONY: build
build: curl-paste.go
	${ENGINE} build . --tag ${IMAGE}

.PHONY: run
run: curl-paste.go
	${ENGINE} run --detach --name ${NAME} --rm --net host ${IMAGE}

.PHONY: logs
logs:
	${ENGINE} logs --follow ${NAME}

.PHONY: kill
kill:
	-${ENGINE} kill ${NAME}
