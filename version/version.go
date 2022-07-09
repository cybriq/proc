package version

//go:generate go run ./update/.

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "gitlab.com/cybriqsystems/proc"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "7bbe20b4197e58e42919d9d09a7feec1cbccbf57"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2022-07-09T07:53:53+03:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.0.6"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/davidvennik/src/gitlab.com/cybriqsystems/proc/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 0
	// Patch is the patch version number from the tag
	Patch = 6
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tCommit: "+GitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tTag: "+Tag+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
	)
}
