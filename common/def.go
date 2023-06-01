package common

import (
	"github.com/eiannone/keyboard"
)

type Pos struct {
	X int
	Y int
}
type IContentHooker interface {
	CurrentLineChange()
}
type ITextContent interface {
	SetCurrentLine(content string)
}
type IKeyHooker interface {
	KeyEnter()
	KeyAfterEnter()
	KeyTable()
	KeyBeforeBackspace() bool
}
type IDataSource interface {
	LinesData() []string
	Listener(listener IDataListener)
}
type IDataListener interface {
	DataChanged()
}

type IInputReceiver interface {
	IKeyInputReceiver
	ITextInputReceiver
}
type IKeyInputReceiver interface {
	InputKey(key keyboard.Key)
}
type ITextInputReceiver interface {
	InputChar(char rune)
}

type ICmdReceiver interface {
	Cmd(cmd string)
	CmdList() []string
}
