package main

import (
	"errors"

	l "gitlab.com/cybriqsystems/proc/pkg/log"
	"gitlab.com/cybriqsystems/proc/version"
)

var log = l.Get(version.PathBase)

func main() {
	// log.I.C(version.Get)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.I.Chk(errors.New("error check"))
	log.W.Ln("testing")
	log.E.Ln("testing")
	log.F.Ln("testing")
}
