package show

import (
	"time"

	"github.com/wormggmm/gocmd/common"
)

type Manager struct {
	blocks       map[int64]IDrawable
	exitCh       chan interface{}
	drawCh       chan IDrawable
	updateTicker *time.Ticker
	controller   *common.ScreenController
}

func NewManager() *Manager {
	return &Manager{
		blocks:     map[int64]IDrawable{},
		exitCh:     make(chan interface{}),
		drawCh:     make(chan IDrawable, 10),
		controller: &common.ScreenController{},
	}
}
func (s *Manager) Start() {
	s.updateTicker = time.NewTicker(1 * time.Second)
	go s.update()
}
func (s *Manager) Stop() {
	close(s.exitCh)
}
func (s *Manager) AddBlock(block IDrawable) {
	block.SetManager(s)
	s.blocks[block.ID()] = block
}

func (s *Manager) FlushDraw(block IDrawable) {
	s.drawCh <- block
}

func (s *Manager) SetCursorPos(row, col int) {
	s.controller.SetCursorPos(row, col)
}
func (s *Manager) update() {
	exit := false
	for !exit {
		select {
		case <-s.exitCh:
			exit = true
		case drawBlock := <-s.drawCh:
			s.controller.SaveCursorPos()
			drawBlock.Draw()
			s.controller.LoadCursorPos()
		case <-s.updateTicker.C:
			s.draw()
		}
	}
}
func (s *Manager) draw() {
	s.controller.SaveCursorPos()
	for _, b := range s.blocks {
		b.Draw()
	}
	s.controller.LoadCursorPos()
}
