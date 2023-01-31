package show

import (
	"fmt"
	"time"
)

type Manager struct {
	blocks []*Block
	exitCh chan interface{}
}

func NewManager() *Manager {
	return &Manager{
		exitCh: make(chan interface{}),
	}
}
func (s *Manager) Start() {
	go s.update()
}
func (s *Manager) Stop() {
	close(s.exitCh)
}
func (s *Manager) AddBlock(block *Block) {
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
	fmt.Println("Manager draw")
}
