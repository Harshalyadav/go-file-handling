package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Fatal("jello")
	log.SetLevel(log.Debug)
	log.SetLevel(log.Trace)
	log.Debug("jello")
	log.Trace("hello")
}
