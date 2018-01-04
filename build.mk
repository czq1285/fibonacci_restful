#Common parameters to build
BINARY=$(shell basename $(dir $(realpath $(firstword $(MAKEFILE_LIST)))))
DEPS=
BUILDARG=
GOGETARG=-v
UTARG=-cover -v

#default target is to build the binary
build: deps
	go build -o ./${BINARY} ${BUILDARG}

install: build
	go install

#DEPS variable could be defined in sub-makefile
deps:
	@test -z "${DEPS}" || go get ${GOGETARG} ${DEPS}

ut:
	go test ${UTARG}

clean:
	go clean

.PHONY: build deps ut clean
