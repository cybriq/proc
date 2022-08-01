package proc

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "github.com/cybriq/proc"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "cb7e647d44929a9f28fb0af6dd70afaa33b71ab4"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-08-01T08:02:22+02:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.29"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/gitlab.com/cybriqsystems/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 29
)

// Version returns a pretty printed version information string
func Version() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tCommit: "+GitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tSemVer: "+SemVer+"\n",
	)
}
