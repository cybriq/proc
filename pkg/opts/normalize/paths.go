package normalize

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func ResolvePath(input string, abs bool) (cleaned string, e error) {
	if strings.HasPrefix(input, "~") {
		homeDir := getHomeDir()
		input = strings.Replace(input, "~", homeDir, 1)
	}
	cleaned = filepath.Clean(input)
	if abs {
		if cleaned, e = filepath.Abs(cleaned); log.E.Chk(e) {
			return
		}
	}
	return
}

func getHomeDir() (homeDir string) {
	var usr *user.User
	var e error
	if usr, e = user.Current(); e == nil {
		homeDir = usr.HomeDir
	}
	// Fall back to standard HOME environment variable that
	// works for most POSIX OSes if the directory from the
	// Go standard lib failed.
	if e != nil || homeDir == "" {
		homeDir = os.Getenv("HOME")
	}
	return homeDir
}
