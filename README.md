# go-release

Exposes information about the build of projects.

You have to set release variables in this package using linker flags, for example in your Makefile:

```
PKG         = github.com/snebel29/go-release/pkg/release
BRANCH      = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT      = $(shell git show -s --format=%h)
COMMIT_TIME = $(shell git show -s --format=%cI)
BUILD_TIME  = $(shell date --iso-8601='seconds')

ifeq ($(VERSION),)
VERSION := $(git tag --points-at HEAD)
endif
ifeq ($(VERSION),)
VERSION := unknown
endif

LDFLAGS = "-X ${PKG}.version=${VERSION} \
           -X ${PKG}.branch=${BRANCH} \
           -X ${PKG}.branch=${COMMIT} \
           -X ${PKG}.branch=${COMMIT_TIME} \
           -X ${PKG}.branch=${BUILD_TIME}"
```

Then build

```
$ go build -ldflags ${LDFLAGS} main.go
```
