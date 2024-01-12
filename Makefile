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
	GOARCH=amd64 \
	GOOS=linux \
	go build -ldflags "-X main.Version=${VERSION}" -a -o ${BUILD_DIR}/${BINARY}-${VERSION} main.go

.PHONY: build-darwin
build-darwin:
	GOARCH=amd64 \
	GOOS=darwin \
	go build -ldflags "-X main.Version=${VERSION}" -a -o ${BUILD_DIR}/${BINARY}-${VERSION} main.go

.PHONY: build-windows
build-windows:
	set CGO_ENABLED=0 && \
	${GO_LOCATION}\bin\go.exe build -ldflags "-s -w -X main.Version=${VERSION}" -a -o .\build\${WINDOWS} .\main.go

.PHONY: deps
deps:
	dep ensure;

.PHONY: test
test:
	go test -v ./tests > report.txt
