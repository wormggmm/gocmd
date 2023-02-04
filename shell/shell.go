package shell

import "github.com/wormggmm/gocmd/common"

type Shell struct {
	*BaseShell
	receiver common.ICmdReceiver
}

func NewShell(receiver common.ICmdReceiver) *Shell {
	sh := &Shell{
		receiver: receiver,
	}
	sh.BaseShell = NewBaseShell(sh)
	return sh
}
func (s *Shell) KeyEnter() {
	cmd := s.currentLine
	if s.receiver != nil {
		s.receiver.Cmd(cmd)
	}
}
func (s *Shell) Data() string {
	return ">" + s.content + "_"
}
