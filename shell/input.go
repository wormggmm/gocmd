package shell

import (
	"github.com/wormggmm/gocmd/common"

	"github.com/eiannone/keyboard"
)

type Input struct {
	receiver common.IInputReceiver
	exitCh   chan interface{}
}

func NewInput() *Input {
	return &Input{
		exitCh: make(chan interface{}),
	}
}
func (s *Input) Start() error {
	if err := keyboard.Open(); err != nil {
		return err
	}
	go s.listening()
	return nil
}

func (s *Input) Stop() {
	close(s.exitCh)
	keyboard.Close()
}

func (s *Input) SetReceiver(receiver common.IInputReceiver) {
	s.receiver = receiver
}
func (s *Input) listening() {
	stop := false
	for !stop {
		select {
		case <-s.exitCh:
			stop = true
		default:
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			// fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
			if s.receiver == nil {
				continue
			}
			if key != 0 {
				s.receiver.InputKey(key)
			} else {
				s.receiver.InputChar(char)
			}
		}
	}
}
