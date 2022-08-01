package main

import (
	"errors"
	"time"

	"github.com/cybriq/proc"
)

var log = proc.GetLogger(proc.PathBase)

func main() {
	proc.App = "logtest"
	log.I.C(proc.Version)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.W.Ln("testing")
	log.E.Chk(errors.New("testing"))
	log.F.Ln("testing")
	log.I.S(proc.AllSubsystems)
	log.I.Ln("setting timestamp format to RFC822Z")
	proc.SetTimeStampFormat(time.RFC822Z)
	log.I.Ln("setting log level to info and printing from all levels")
	proc.SetLogLevel(proc.Info)
	log.T.Ln("testing")
	log.D.Ln("testing")
	log.I.Ln("testing")
	log.W.Ln("testing")
	log.E.Chk(errors.New("testing"))
	log.F.Ln("testing")
	log.T.Ln("testing")

}
