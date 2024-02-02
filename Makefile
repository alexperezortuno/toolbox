EXECUTABLE=toolbox
VERSION=$(git describe --tags --always --long --dirty)
BUILD_DIR=./build
WINDOWS=$(EXECUTABLE)_windows_amd64_$(VERSION).exe
LINUX=$(EXECUTABLE)_linux_amd64_$(VERSION)
DARWIN=$(EXECUTABLE)_darwin_amd64_$(VERSION)

.PHONY: default
default: build

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: build
build: clean build-linux build-darwin build-windows

.PHONY: build-version
build-version:
	go build -a -o ${BUILD_DIR}/${BINARY}-${VERSION} ./main.go

.PHONY: build-linux
build-linux:
	env CGO_ENABLED=0 && \
	env GOOS=linux && \
	env GOARCH=amd64 && \
	${GOROOT}/bin/go build -ldflags "-s -w -X main.Version=${VERSION}" -a -o ./build/${LINUX} main.go

.PHONY: build-darwin
build-darwin:
	env CGO_ENABLED=0 && \
	env GOOS=darwin GOARCH=amd64 ${GOROOT}/bin/go build -ldflags "-s -w -X main.Version=${VERSION}" -a -o ./build/${DARWIN} main.go

.PHONY: build-windows
build-windows:
	set CGO_ENABLED=0 && \
	set GOOS=windows && \
	set GOARCH=amd64 && \
	 ${GOROOT}\bin\go.exe build -ldflags "-s -w -X main.Version=${VERSION}" -a -o .\build\${WINDOWS} .\main.go

.PHONY: deps
deps:
	dep ensure;

.PHONY: test
test:
	env CGO_ENABLED=0 && \
	${GOROOT}/bin/go test -v ./tests > report.txt
