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
	sh.BaseShell = NewBaseShell(sh)
	return sh
}
func (s *Shell) SetReceiver(receiver common.ICmdReceiver) {
	s.receiver = receiver
}
func (s *Shell) KeyAfterEnter() {
	s.InputChar('>')
}
func (s *Shell) KeyEnter() {
	cmd := s.currentLine
	if s.receiver != nil {
		s.receiver.Cmd(cmd)
	}
}
func (s *Shell) CurrentLineChange() {
	s.tableKey = false
}
func (s *Shell) KeyTable() {
	if s.tableKey && s.receiver != nil {
		currentCmd := s.currentLine
		s.InputChar('\n')
		cmdList := s.receiver.CmdList()
		for _, cmd := range cmdList {
			if strings.Index(cmd, currentCmd) == 0 {
				s.content += cmd + "\t"
			}
		}
		s.InputChar('\n')
		s.InputChar('>')
		s.content += currentCmd
		s.currentLine = currentCmd
	}
	s.tableKey = true
}
func (s *Shell) Data() string {
	return ">" + s.content + "_"
}
