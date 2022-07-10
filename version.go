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
	GitCommit = "42589fe4feec0762f1679b2ca168705c8ec84e51"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-07-10T11:45:44+03:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.25"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/gitlab.com/cybriqsystems/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 25
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
