package common

import (
	"github.com/eiannone/keyboard"
)

type Pos struct {
	X int
	Y int
}

type IDataSource interface {
	Data() string
	Listener(listener IDataListener)
}
type IDataListener interface {
	DataChanged()
}

type IInputReceiver interface {
	InputChar(char rune)
	InputKey(key keyboard.Key)
}

type ICmdReceiver interface {
	Cmd(cmd string)
}
