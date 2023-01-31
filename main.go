package main

import (
	"time"

	"github.com/wormggmm/gocmd/common"
	"github.com/wormggmm/gocmd/show"
)

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

	b := show.NewBlock(10, 10, 50, 1)
	for i := 0; i < 2; i++ {
		b.Printf("01234567890123456789\n")
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
	c := &common.ScreenController{}
	c.Reset()
	c.ClearAll()
	c.SetCursorPos(0, 0)
}
