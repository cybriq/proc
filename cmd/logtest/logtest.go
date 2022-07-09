package main

import (
	"errors"
	"time"

	"gitlab.com/cybriqsystems/proc"
)

var log = proc.GetLogger(proc.PathBase)

func main() {
	proc.App = "logtest"
	log.I.C(proc.Get)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.W.Ln("testing")
	log.E.Chk(errors.New("testing"))
	log.F.Ln("testing")
	log.I.S(proc.AllSubsystems)

	proc.SetTimeStampFormat(time.RFC822Z)
	proc.SetLogLevel(proc.Info)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.W.Ln("testing")
	log.E.Chk(errors.New("testing"))
	log.F.Ln("testing")
	log.T.Ln("testing")

}
