// Package indra is the root level package for Indranet, a low latency, 
// Lightning Network monetised distributed VPN protocol designed for providing
// strong anonymity to valuable internet traffic.
package indra

import (
	"fmt"
)

var (
	// URL is the git URL for the repository.
	URL = "github.com/cybriq/proc"
	// GitRef is the gitref, as in refs/heads/branchname.
	GitRef = "refs/heads/master"
	// ParentGitCommit is the commit hash of the parent HEAD.
	ParentGitCommit = "e2a2931897cff030b29c19e33a54b649a939ea94"
	// BuildTime stores the time when the current binary was built.
	BuildTime = "2022-12-23T18:23:27Z"
	// SemVer lists the (latest) git tag on the build.
	SemVer = "v0.20.4"
	// PathBase is the path base returned from runtime caller.
	PathBase = "/home/loki/src/github.com/cybriq/proc/"
	// Major is the major number from the tag.
	Major = 0
	// Minor is the minor number from the tag.
	Minor = 20
	// Patch is the patch version number from the tag.
	Patch = 4
)

// Version returns a pretty printed version information string.
func Version() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tParentGitCommit: "+ParentGitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tSemVer: "+SemVer+"\n",
	)
}
