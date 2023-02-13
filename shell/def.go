package shell

type IShellHooker interface {
	CurrentLineChange()
	KeyEnter()
	KeyAfterEnter()
	KeyTable()
	KeyBeforeBackspace() bool
}
