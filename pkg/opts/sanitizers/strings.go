package sanitizers

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	Password  = "password"
	FilePath  = "filepath"
	Directory = "directory"
)

func StringType(typ, input string, defaultPort int) (cleaned string, e error) {
	switch typ {
	case Password:
		// password type is mainly here for the input method of the app using this config library
	case FilePath:
		if strings.HasPrefix(input, "~") {
			var homeDir string
			var usr *user.User
			var e error
			if usr, e = user.Current(); e == nil {
				homeDir = usr.HomeDir
			}
			// Fall back to standard HOME environment variable that works for most POSIX OSes if the directory from the Go
			// standard lib failed.
			if e != nil || homeDir == "" {
				homeDir = os.Getenv("HOME")
			}

			input = strings.Replace(input, "~", homeDir, 1)
		}
		if cleaned, e = filepath.Abs(filepath.Clean(input)); log.E.Chk(e) {
		}
	case Directory:
		if strings.HasPrefix(input, "~") {
			var homeDir string
			var usr *user.User
			var e error
			if usr, e = user.Current(); e == nil {
				homeDir = usr.HomeDir
			}
			// Fall back to standard HOME environment variable that works for most POSIX OSes if the directory from the Go
			// standard lib failed.
			if e != nil || homeDir == "" {
				homeDir = os.Getenv("HOME")
			}

			input = strings.Replace(input, "~", homeDir, 1)
		}
		if cleaned, e = filepath.Abs(filepath.Clean(input)); log.E.Chk(e) {
		}
	default:
		cleaned = input
	}
	return
}
