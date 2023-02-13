package shell

import (
	"github.com/eiannone/keyboard"
	"github.com/wormggmm/gocmd/common"
)

type BaseShell struct {
	hooker      IShellHooker
	listener    common.IDataListener
	lines       []string
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
	// s.content += string(char)
	if char == '\n' {
		s.enterLine()
	} else {
		s.currentLine += string(char)
	}
	if s.hooker != nil {
		s.hooker.CurrentLineChange()
	}
	if s.listener != nil {
		s.listener.DataChanged()
	}
}
func (s *BaseShell) enterLine() {
	s.lines = append(s.lines, s.currentLine)
	s.currentLine = ""
}
func (s *BaseShell) InputKey(key keyboard.Key) {
	switch key {
	case keyboard.KeyEnter:
		if s.hooker != nil {
			s.hooker.KeyEnter()
			s.hooker.CurrentLineChange()
		}
		// s.content += "\n"
		s.enterLine()
		if s.hooker != nil {
			s.hooker.KeyAfterEnter()
		}
	case keyboard.KeyBackspace2:
		if len(s.currentLine) > 0 {
			if s.hooker.KeyBeforeBackspace() {
				break
			}
			// s.content = s.content[0 : len(s.content)-1]
			s.currentLine = s.currentLine[0 : len(s.currentLine)-1]
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
