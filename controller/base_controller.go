package controller

import (
	"github.com/eiannone/keyboard"
	"github.com/wormggmm/gocmd/common"
	"github.com/wormggmm/gocmd/content"
)

type BaseController struct {
	*content.TextContent
	hooker   common.IKeyHooker
	listener common.IDataListener
}

func NewBaseController(hooker common.IKeyHooker) *BaseController {
	return &BaseController{
		TextContent: &content.TextContent{},
		hooker:      hooker,
	}
}
func (s *BaseController) Listener(listener common.IDataListener) {
	s.listener = listener
}
func (s *BaseController) Input(content string) {
	for _, char := range content {
		s.InputChar(char)
	}
}
func (s *BaseController) InputChar(char rune) {
	s.TextContent.InputChar(char)
	if s.hooker != nil {
		s.hooker.CurrentLineChange()
	}
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
func (s *BaseController) InputKey(key keyboard.Key) {
	switch key {
	case keyboard.KeySpace:
		if s.hooker != nil {
			s.hooker.CurrentLineChange()
		}
		s.InputChar(' ')
	case keyboard.KeyEnter:
		if s.hooker != nil {
			s.hooker.KeyEnter()
			s.hooker.CurrentLineChange()
		}
		s.EnterLine()
		if s.hooker != nil {
			s.hooker.KeyAfterEnter()
		}
	case keyboard.KeyBackspace2:
		if s.CurrentLineLen() > 0 {
			if s.hooker.KeyBeforeBackspace() {
				break
			}
			s.BackSpace()
			if s.hooker != nil {
				s.hooker.CurrentLineChange()
			}
		}
	case keyboard.KeyTab:
		if s.hooker != nil {
			s.hooker.KeyTable()
		}
	}
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
