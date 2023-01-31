package shell

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type Input struct{}

func (s *Input) Start() error {
	if err := keyboard.Open(); err != nil {
		return err
	}
	go s.listening()
	return nil
}

func (s *Input) Stop() {
	keyboard.Close()
}

func (s *Input) listening() {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		}
	}
}
