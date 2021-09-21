package run

import (
	log "github.com/sirupsen/logrus"
)

type Run struct{}

func NewRun() *Run {
	return &Run{}
}

func (r *Run) Runner() {
	log.Debugf("Run.")
}
