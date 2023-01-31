package show

import (
	"time"

	"github.com/wormggmm/gocmd/common"
)

type Manager struct {
	blocks     []IDrawable
	exitCh     chan interface{}
	controller *common.ScreenController
}

func NewManager() *Manager {
	return &Manager{
		exitCh:     make(chan interface{}),
		controller: &common.ScreenController{},
	}
}
func (s *Manager) Start() {
	go s.update()
}
func (s *Manager) Stop() {
	close(s.exitCh)
}
func (s *Manager) AddBlock(block IDrawable) {
	s.blocks = append(s.blocks, block)
}

func (s *Manager) update() {
	exit := false
	for !exit {
		select {
		case <-s.exitCh:
			exit = true
		default:
			s.draw()
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
func (s *Manager) draw() {
	s.controller.SaveCursorPos()
	for _, b := range s.blocks {
		b.Draw()
	}
	s.controller.LoadCursorPos()
}
