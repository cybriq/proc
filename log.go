package proc

import "github.com/gookit/color"

type LogLevel int32

const (
	Off LogLevel = iota
	Fatal
	Error
	Check
	Warn
	Info
	Debug
	Trace
)

var LvlStr = map[LogLevel]string{
	Off:   "off",
	Fatal: "fatal",
	Error: "error",
	Warn:  "warn",
	Info:  "info",
	Check: "check",
	Debug: "debug",
	Trace: "trace",
}

type (
	// Println prints lists of interfaces with spaces in between
	Println func(a ...interface{})

	// Printf prints like fmt.Println surrounded by log details
	Printf func(format string, a ...interface{})

	// Prints  prints a spew.Sdump for an interface slice
	Prints func(a ...interface{})

	// Printc accepts a function so that the extra computation can be avoided if it is
	// not being viewed
	Printc func(closure func() string)

	// Chk is a shortcut for printing if there is an error, or returning true
	Chk func(e error) bool

	// LevelPrinter defines a set of terminal printing primitives that output with
	// extra data, time, log logLevelList, and code location
	LevelPrinter struct {
		Ln Println
		// F prints like fmt.Println surrounded by log details
		F Printf
		// S uses spew.dump to show the content of a variable
		S Prints
		// C accepts a function so that the extra computation can be avoided if it is
		// not being viewed
		C Printc
		// Chk is a shortcut for printing if there is an error, or returning true
		Chk Chk
	}

	LevelSpec struct {
		Name      string
		Colorizer func(format string, a ...interface{}) string
	}
)

func gLS(lvl LogLevel, r, g, b byte) LevelSpec {
	return LevelSpec{
		Name:      LvlStr[lvl],
		Colorizer: color.Bit24(r, g, b, false).Sprintf,
	}
}

// LevelSpecs specifies the id, string name and color-printing function
var LevelSpecs = map[LogLevel]LevelSpec{
	Off:   gLS(Off, 0, 0, 0),
	Fatal: gLS(Fatal, 128, 0, 0),
	Error: gLS(Error, 255, 0, 0),
	Check: gLS(Check, 255, 255, 0),
	Warn:  gLS(Warn, 0, 255, 0),
	Info:  gLS(Info, 255, 255, 0),
	Debug: gLS(Debug, 0, 128, 255),
	Trace: gLS(Trace, 128, 0, 255),
}

func logPrint(level LogLevel, subsystem string, printFunc func()) func() {
	return func() {

	}
}

type Logger struct {
}
