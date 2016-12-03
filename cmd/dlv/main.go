package main

import (
	"github.com/landonb/delve/cmd/dlv/cmds"
	"github.com/landonb/delve/version"
)

// Build is the git sha of this binaries build.
var Build string

func main() {
	version.DelveVersion.Build = Build
	cmds.New().Execute()
}
