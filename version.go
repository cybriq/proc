package proc

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "gitlab.com/cybriqsystems/proc"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "e98186012bb0c4279fb5768e73d6c5685bd31e3a"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-08-01T07:21:06+02:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.28"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/gitlab.com/cybriqsystems/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 28
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
