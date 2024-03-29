// Package release exposes information about the build of projects.
// You have to set release variables in this package using linker flags, for example in your Makefile:
//
// BRANCH      = $(shell git rev-parse --abbrev-ref HEAD)
// COMMIT      = $(shell git show -s --format=%h)
// COMMIT_TIME = $(shell git show -s --format=%cI)
// BUILD_TIME  = $(shell date --iso-8601='seconds')
// RELEASE_PKG = github.com/snebel29/go-release/pkg/release
//
// ifeq ($(VERSION),)
// VERSION := $(git tag --points-at HEAD)
// endif
// ifeq ($(VERSION),)
// VERSION := unknown
// endif
// LDFLAGS = "-X ${RELEASE_PKG}.version=${VERSION} \
//            -X ${RELEASE_PKG}.branch=${BRANCH} \
//            -X ${RELEASE_PKG}.commit=${COMMIT} \
//            -X ${RELEASE_PKG}.commitTime=${COMMIT_TIME} \
//            -X ${RELEASE_PKG}.buildTime=${BUILD_TIME}"
//
// Then build:
// go build -ldflags ${LDFLAGS} main.go
package release

import (
	"fmt"
	"runtime"
)

var (
	version    string
	branch     string
	commit     string
	commitTime string
	buildTime  string
)

// Release defines the release attributes.
type Release struct {
	Version    string
	Branch     string
	Commit     string
	CommitTime string
	BuildTime  string
	Go         string
	Platform   string
}

// Get return the release object.
func Get() Release {
	return Release{
		Version:    version,
		Branch:     branch,
		Commit:     commit,
		CommitTime: commitTime,
		BuildTime:  buildTime,
		Go:         runtime.Compiler + "/" + runtime.Version(),
		Platform:   runtime.GOOS + "/" + runtime.GOARCH,
	}
}

// String return a string represenation of the release while implementing fmt.Stringer interface.
func (r Release) String() string {
	return fmt.Sprintf("version=%s go=%s platform=%s branch=%s commit=%s commit-time=%s build-time=%s",
		r.Version, r.Go, r.Platform, r.Branch, r.Commit, r.CommitTime, r.BuildTime)
}
