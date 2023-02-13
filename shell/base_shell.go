package shell

import (
	"github.com/eiannone/keyboard"
	"github.com/wormggmm/gocmd/common"
)

type BaseShell struct {
	common.TextContent
	hooker   IShellHooker
	listener common.IDataListener
}

func NewBaseShell(hooker IShellHooker) *BaseShell {
	return &BaseShell{
		hooker: hooker,
	}
}
func (s *BaseShell) Listener(listener common.IDataListener) {
	s.listener = listener
}
func (s *BaseShell) InputChar(char rune) {
	s.TextContent.InputChar(char)
	if s.hooker != nil {
		s.hooker.CurrentLineChange()
	}
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
func (s *BaseShell) InputKey(key keyboard.Key) {
	switch key {
	case keyboard.KeyEnter:
		if s.hooker != nil {
			s.hooker.KeyEnter()
			s.hooker.CurrentLineChange()
		}
		s.TextContent.EnterLine()
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
