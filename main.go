package main

import (
	"time"

	"github.com/wormggmm/gocmd/common"
	"github.com/wormggmm/gocmd/show"
)

type TestDataSource struct {
	Str string
}

func (s *TestDataSource) Data() string {
	return s.Str
}
func main() {
	// fmt.Print("\r\033[K")
	// c := &common.ScreenController{}
	// c.Reset()
	// c.ClearAll()
	// rand.Seed(time.Now().UnixMilli())
	// for i := 0; i < 50; i++ {
	// 	row := rand.Intn(50)
	// 	col := rand.Intn(300)
	// 	c.SetCursorPos(row, col)
	// 	fmt.Print("A")
	// 	time.Sleep(500 * time.Millisecond)
	// }
	// c.ClearLine()
	// c.FontColor(common.EnumColor.Red)
	// c.FontBackColor(common.EnumColor.Green)
	// c.Blink()
	// fmt.Print("hello")
	// cmd.CMD().Start()
	// time.Sleep(2 * time.Second)
	// cmd.CMD().Stop()
	// c.ClearAll()
	// c.Reset()
	dataSrc := &TestDataSource{}
	b := show.NewBlock(10, 10, 20, 10, dataSrc)
	b.SetFrame('*', true, common.EnumColor.Red)
	mgr := show.NewManager()
	mgr.AddBlock(b)
	mgr.Start()

	for i := 0; i < 20; i++ {
		dataSrc.Str += "A"
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
	c := &common.ScreenController{}
	c.Reset()
	c.ClearAll()
	c.SetCursorPos(0, 0)
}
