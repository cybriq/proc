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
	GitCommit = "cb0c69008b86954d7b1cb1f816bf1be8e10e36ae"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-07-09T17:26:44+03:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.16"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/gitlab.com/cybriqsystems/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 16
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
