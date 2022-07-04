package main

import (
	"errors"

	l "gitlab.com/cybriqsystems/proc"
	"gitlab.com/cybriqsystems/proc/version"
)

var log = l.Get(version.PathBase)

func main() {
	l.App = "logtest"
	// log.I.C(version.Get)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.W.Ln("testing")
	log.E.Chk(errors.New("testing"))
	log.F.Ln("testing")
}
