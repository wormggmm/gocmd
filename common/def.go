package common

type Pos struct {
	X int
	Y int
}

type IDataSource interface {
	Data() string
}
