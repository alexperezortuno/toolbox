EXECUTABLE=toolbox
VERSION=0.0.1
BUILD_DIR=./build
WINDOWS=$(EXECUTABLE)_windows_amd64_$(VERSION).exe
LINUX=$(EXECUTABLE)_linux_amd64_$(VERSION)
DARWIN=$(EXECUTABLE)_darwin_amd64_$(VERSION)

.PHONY: default
default: build

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: build-version
build-version:
	go build -a -o ${BUILD_DIR}/${BINARY}-${VERSION} main.go

.PHONY: build-linux
build-linux:
	set CGO_ENABLED=0 && \
	${GOROOT}/bin/go build -ldflags "-s -w -X main.Version=${VERSION}" -a -o build/${LINUX} main.go

.PHONY: build-darwin
build-darwin:
	set CGO_ENABLED=0 && \
	${GOROOT}/bin/go build -ldflags "-s -w -X main.Version=${VERSION}" -a -o build/${DARWIN} main.go

.PHONY: build-windows
build-windows:
	set CGO_ENABLED=0 && \
	${GOROOT}\bin\go.exe build -ldflags "-s -w -X main.Version=${VERSION}" -a -o .\build\${WINDOWS} .\main.go

.PHONY: deps
deps:
	dep ensure;

.PHONY: test
test:
	set CGO_ENABLED=0 && \
	${GOROOT}/bin/go test -v ./tests > report.txt
