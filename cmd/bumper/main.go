package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/cybriqsystems/proc"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

var (
	URL                 string
	GitRef              string
	GitCommit           string
	BuildTime           string
	SemVer              string
	Major, Minor, Patch int
	Meta                string
	PathBase            string
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(
			os.Stderr,
			"arguments required in order to bump and push this repo",
		)
		os.Exit(1)
	}
	proc.App = "bumper"
	BuildTime = time.Now().Format(time.RFC3339)
	var cwd string
	var e error
	if cwd, e = os.Getwd(); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	var repo *git.Repository
	if repo, e = git.PlainOpen(cwd); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	var rr []*git.Remote
	if rr, e = repo.Remotes(); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	for i := range rr {
		rs := rr[i].String()
		if strings.HasPrefix(rs, "origin") {
			rss := strings.Split(rs, "git@")
			if len(rss) > 1 {
				rsss := strings.Split(rss[1], ".git")
				URL = strings.ReplaceAll(rsss[0], ":", "/")
				break
			}
			rss = strings.Split(rs, "https://")
			if len(rss) > 1 {
				rsss := strings.Split(rss[1], ".git")
				URL = rsss[0]
				break
			}

		}
	}
	var tr *git.Worktree
	if tr, e = repo.Worktree(); log.E.Chk(e) {
		fmt.Println(e)
	}
	var rh *plumbing.Reference
	if rh, e = repo.Head(); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	rhs := rh.Strings()
	GitRef = rhs[0]
	GitCommit = rhs[1]
	var rt storer.ReferenceIter
	if rt, e = repo.Tags(); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	var maxVersion int
	if e = rt.ForEach(
		func(pr *plumbing.Reference) (e error) {
			s := strings.Split(pr.String(), "/")
			prs := s[2]
			if strings.HasPrefix(prs, "v") {
				var va [3]int
				_, _ = fmt.Sscanf(
					prs,
					"v%d.%d.%d",
					&va[0],
					&va[1],
					&va[2],
				)
				vn := va[0]*1000000 + va[1]*1000 + va[2]
				if maxVersion < vn {
					maxVersion = vn
					Major = va[0]
					Minor = va[1]
					Patch = va[2]
				}
			}
			return
		},
	); log.E.Chk(e) {
		fmt.Println(e)
		return
	}
	// Bump to next patch version every time
	Patch++
	// Update SemVer
	SemVer = fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
	PathBase = tr.Filesystem.Root() + "/"
	versionFile := `package proc

import (
	"fmt"
)

var (
	// URL is the git URL for the repository
	URL = "%s"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "%s"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "%s"
	// BuildTime stores the time when the current binary was built
	BuildTime = "%s"
	// SemVer lists the (latest) git tag on the build
	SemVer = "%s"
	// PathBase is the path base returned from runtime caller
	PathBase = "%s"
	// Major is the major number from the tag
	Major = %d
	// Minor is the minor number from the tag
	Minor = %d
	// Patch is the patch version number from the tag
	Patch = %d
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
`
	versionFileOut := fmt.Sprintf(
		versionFile,
		URL,
		GitRef,
		GitCommit,
		BuildTime,
		SemVer,
		PathBase,
		Major,
		Minor,
		Patch,
	)
	path := filepath.Join(PathBase, "version.go")
	if e = ioutil.WriteFile(path, []byte(versionFileOut), 0666); log.E.Chk(e) {
		fmt.Println(e)
	}
	log.I.Ln(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tCommit: "+GitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tSemVer: "+SemVer+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
	)
	e = runCmd("git", "add", ".")
	if log.E.Chk(e) {
		panic(e)
	}
	commitString := strings.Join(os.Args[1:], " ")

	commitString = strings.ReplaceAll(commitString, " -- ", "\n\n")

	e = runCmd("git", "commit", "-m"+commitString)
	if log.E.Chk(e) {
		panic(e)
	}
	e = runCmd("git", "tag", SemVer)
	if log.E.Chk(e) {
		panic(e)
	}
	gr := strings.Split(GitRef, "/")
	branch := gr[2]
	e = runCmd("git", "push", "origin", branch)
	if log.E.Chk(e) {
		panic(e)
	}
	e = runCmd("git", "push", "origin", SemVer)
	if log.E.Chk(e) {
		panic(e)
	}
	return
}

func runCmd(cmd ...string) (err error) {

	c := exec.Command(cmd[0], cmd[1:]...)
	var output []byte
	output, err = c.CombinedOutput()
	if err == nil && string(output) != "" {
		log.I.Ln("\n", string(output))
	}
	return
}