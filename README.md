# go-release

Exposes information about the build of projects.

## Usage

Write your code
```
package main

import (
    "fmt"
    "github.com/snebel29/go-release/pkg/release"
)

func main() {
    fmt.Println(release.Get().String())
}
```


Then you have to set release variables in this package using linker flags, for example in your Makefile:
```
BRANCH      = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT      = $(shell git show -s --format=%h)
COMMIT_TIME = $(shell git show -s --format=%cI)
BUILD_TIME  = $(shell date --iso-8601='seconds')
RELEASE_PKG = github.com/snebel29/go-release/pkg/release

ifeq ($(VERSION),)
VERSION := $(git tag --points-at HEAD)
endif
ifeq ($(VERSION),)
VERSION := unknown
endif

LDFLAGS = "-X ${RELEASE_PKG}.version=${VERSION} \
           -X ${RELEASE_PKG}.branch=${BRANCH} \
           -X ${RELEASE_PKG}.branch=${COMMIT} \
           -X ${RELEASE_PKG}.branch=${COMMIT_TIME} \
           -X ${RELEASE_PKG}.branch=${BUILD_TIME}"
```

Then build

```
$ go build -ldflags ${LDFLAGS} main.go
```

