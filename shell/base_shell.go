package shell

import (
	"github.com/eiannone/keyboard"
	"github.com/wormggmm/gocmd/common"
)

type BaseShell struct {
	hooker      IShellHooker
	listener    common.IDataListener
	content     string
	currentLine string
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
	s.content += string(char)
	s.currentLine += string(char)
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
func (s *BaseShell) InputKey(key keyboard.Key) {
	switch key {
	case keyboard.KeyEnter:
		if s.hooker != nil {
			s.hooker.KeyEnter()
		}
		s.content += "\n"
		s.currentLine = ""
	case keyboard.KeyBackspace2:
		if len(s.currentLine) > 0 && len(s.content) > 0 {
			s.content = s.content[0 : len(s.content)-1]
			s.currentLine = s.currentLine[0 : len(s.currentLine)-1]
		}
	}
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
