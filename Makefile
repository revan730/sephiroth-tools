ifneq ($(shell git rev-parse --abbrev-ref HEAD), master)
  VERSION = $(shell git rev-parse --abbrev-ref HEAD):$(shell git rev-parse --short HEAD)
else
  VERSION = $(shell git tag)
endif
LDFLAGS = -ldflags "-X github.com/revan730/sephiroth-tools/cmd.version=${VERSION}"
BUILD_DIR = $(shell pwd)/dist
BINARY = sephiroth

build:
	vgo build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}
	cp -R src/templates ${BUILD_DIR}/

clean:
	rm -r ${BUILD_DIR}/
