package proc

import (
	"fmt"
)

var (
	// URL is the git URL for the repository.
	URL = "github.com/cybriq/proc"
	// GitRef is the gitref, as in refs/heads/branchname.
	GitRef = "refs/heads/master"
	// ParentGitCommit is the commit hash of the parent HEAD.
	ParentGitCommit = "c8ef73b41ff34b0cdf9ac970681c7e4d2e252b2a"
	// BuildTime stores the time when the current binary was built.
	BuildTime = "2022-09-22T14:15:49+02:00"
	// SemVer lists the (latest) git tag on the build.
	SemVer = "v0.1.36"
	// PathBase is the path base returned from runtime caller.
	PathBase = "/home/loki/src/github.com/cybriq/proc/"
	// Major is the major number from the tag.
	Major = 0
	// Minor is the minor number from the tag.
	Minor = 1
	// Patch is the patch version number from the tag.
	Patch = 36
)

// Version returns a pretty printed version information string.
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
