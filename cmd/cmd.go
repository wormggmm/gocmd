package cmd

import (
	"sync"

	"github.com/wormggmm/gocmd/show"
)

var cmd *Cmd
var onceCmd sync.Once

func CMD() *Cmd {
	onceCmd.Do(func() {
		cmd = &Cmd{}
		cmd.init()
	})
	return cmd
}

type Cmd struct {
	showManager *show.Manager
}

func (s *Cmd) init() {
	s.showManager = show.NewManager()
}

func (s *Cmd) Start() *Cmd {
	s.showManager.Start()
	return s
}
func (s *Cmd) Stop() *Cmd {
	s.showManager.Stop()
	return s
}
