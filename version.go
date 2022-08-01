package proc

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "github.com/cybriq/proc"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// ParentGitCommit is the commit hash of the parent HEAD
	ParentGitCommit = "d237c4f23975726f576c9f87442bb106434a727e"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-08-01T16:36:22+02:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.36"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/github.com/cybriq/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 36
)

// Version returns a pretty printed version information string
func Version() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tPacethGitCommit: "+ParentGitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tSemVer: "+SemVer+"\n",
	)
}
