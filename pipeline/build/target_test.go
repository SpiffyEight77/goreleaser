package build

import (
	"fmt"
	"testing"

	"github.com/goreleaser/goreleaser/config"
	"github.com/goreleaser/goreleaser/internal/buildtarget"
	"github.com/stretchr/testify/assert"
)

func TestAllBuildTargets(t *testing.T) {
	var assert = assert.New(t)
	var build = config.Build{
		Goos: []string{
			"linux",
			"darwin",
			"freebsd",
		},
		Goarch: []string{
			"386",
			"amd64",
			"arm",
			"arm64",
		},
		Goarm: []string{
			"6",
			"7",
		},
		Ignore: []config.IgnoredBuild{
			{
				Goos:   "darwin",
				Goarch: "386",
			}, {
				Goos:   "linux",
				Goarch: "arm",
				Goarm:  "7",
			},
		},
	}
	assert.Equal([]buildtarget.Target{
		buildtarget.New("linux", "386", ""),
		buildtarget.New("linux", "amd64", ""),
		buildtarget.New("linux", "arm", "6"),
		buildtarget.New("linux", "arm64", ""),
		buildtarget.New("darwin", "amd64", ""),
		buildtarget.New("freebsd", "386", ""),
		buildtarget.New("freebsd", "amd64", ""),
		buildtarget.New("freebsd", "arm", "6"),
		buildtarget.New("freebsd", "arm", "7"),
	}, buildTargets(build))
}

func TestGoosGoarchCombos(t *testing.T) {
	var platforms = []struct {
		os    string
		arch  string
		valid bool
	}{
		// valid targets:
		{"android", "arm", true},
		{"darwin", "386", true},
		{"darwin", "amd64", true},
		{"dragonfly", "amd64", true},
		{"freebsd", "386", true},
		{"freebsd", "amd64", true},
		{"freebsd", "arm", true},
		{"linux", "386", true},
		{"linux", "amd64", true},
		{"linux", "arm", true},
		{"linux", "arm64", true},
		{"linux", "mips", true},
		{"linux", "mipsle", true},
		{"linux", "mips64", true},
		{"linux", "mips64le", true},
		{"linux", "ppc64", true},
		{"linux", "ppc64le", true},
		{"netbsd", "386", true},
		{"netbsd", "amd64", true},
		{"netbsd", "arm", true},
		{"openbsd", "386", true},
		{"openbsd", "amd64", true},
		{"openbsd", "arm", true},
		{"plan9", "386", true},
		{"plan9", "amd64", true},
		{"solaris", "amd64", true},
		{"windows", "386", true},
		{"windows", "amd64", true},
		// invalid targets
		{"darwin", "arm", false},
		{"darwin", "arm64", false},
		{"windows", "arm", false},
		{"windows", "arm64", false},
	}
	for _, p := range platforms {
		t.Run(fmt.Sprintf("%v %v valid=%v", p.os, p.arch, p.valid), func(t *testing.T) {
			assert.Equal(t, p.valid, valid(buildtarget.New(p.os, p.arch, "")))
		})
	}
}
