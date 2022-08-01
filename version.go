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
	ParentGitCommit = "300ece0f404dd87c62e14360eb3aa4bcd1c858e2"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-08-01T16:24:26+02:00"
	// SemVer lists the (latest) git tag on the build
	SemVer = "v0.0.35"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/github.com/cybriq/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 35
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
