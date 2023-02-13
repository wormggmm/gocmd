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
	sh.BaseShell.InputChar('>')
	return sh
}
func (s *Shell) SetReceiver(receiver common.ICmdReceiver) {
	s.receiver = receiver
}
func (s *Shell) KeyBeforeBackspace() (isBreak bool) {
	if len(s.currentLine) == 1 && s.currentLine[0] == '>' {
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
func (s *Shell) setCurrentLine(content string) {
	s.currentLine = ">" + content
}
func (s *Shell) getCurrentLine() string {
	cmd := s.currentLine
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
				s.currentLine += cmd + "\t"
			}
		}
		// s.InputChar('\n')
		// s.content += currentCmd
		s.enterLine()
		s.setCurrentLine(currentCmd)
	}
	s.tableKey = true
}
func (s *Shell) Data() string {
	content := ""
	for _, line := range s.lines {
		content += (line + "\n")
	}
	content += (s.currentLine)
	return content + "_"
}
func (s *Shell) LinesData() []string {
	lines := append(s.lines, s.currentLine)
	return lines
}
