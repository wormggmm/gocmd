package shell

import (
	"strings"

	"github.com/wormggmm/gocmd/common"
)

type Shell struct {
	*BaseShell
	receiver common.ICmdReceiver
	tableKey bool
}

func NewShell() *Shell {
	sh := &Shell{}
	sh.BaseShell = NewBaseShell(sh, sh, sh)
	sh.BaseShell.InputChar('>')
	return sh
}
func (s *Shell) SetReceiver(receiver common.ICmdReceiver) {
	s.receiver = receiver
}
func (s *Shell) KeyBeforeBackspace() (isBreak bool) {
	if len(s.CurrentLine()) == 1 && s.CurrentLine()[0] == '>' {
		return true
	}
	return false
}
func (s *Shell) KeyAfterEnter() {
	s.InputChar('>')
}
func (s *Shell) KeyEnter() {
	cmd := s.getCurrentLine()
	if s.receiver != nil {
		s.receiver.Cmd(cmd)
	}
}
func (s *Shell) SetCurrentLine(content string) {
	s.BaseShell.SetCurrentLine(">" + content)
}
func (s *Shell) getCurrentLine() string {
	cmd := s.CurrentLine()
	cmd = strings.TrimLeft(cmd, ">")
	return cmd
}
func (s *Shell) CurrentLineChange() {
	s.tableKey = false
}
func (s *Shell) KeyTable() {
	if s.tableKey && s.receiver != nil {
		currentCmd := s.getCurrentLine()
		s.InputChar('\n')
		cmdList := s.receiver.CmdList()
		for _, cmd := range cmdList {
			if strings.Index(cmd, currentCmd) == 0 {
				// s.content += cmd + "\t"
				// s.currentLine += cmd + "\t"
				s.BaseShell.SetCurrentLine(s.getCurrentLine() + cmd + "\t")
			}
		}
		// s.InputChar('\n')
		// s.content += currentCmd
		s.EnterLine()
		s.SetCurrentLine(currentCmd)
	}
	s.tableKey = true
}
