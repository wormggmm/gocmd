package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/wormggmm/gocmd/shell"

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
	linesStr := "40"
	columnsStr := "200"
	if len(os.Args) >= 3 {
		linesStr = os.Args[1]
		columnsStr = os.Args[2]
	}
	fmt.Printf("lines:%s columns:%s\n", linesStr, columnsStr)
	lines, err := strconv.Atoi(linesStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// columns, err := strconv.Atoi(columnsStr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	time.Sleep(1 * time.Second)
	c := &common.ScreenController{}
	c.Reset()
	c.ClearAll()
	dataSrc := &TestDataSource{}
	b1 := show.NewBlock(2, 2, 5, 10, dataSrc)
	b1.SetFrame('*', true, common.EnumColor.Red)
	mgr := show.NewManager()
	mgr.AddBlock(b1)
	mgr.Start()

	sh := &shell.Shell{}
	b2 := show.NewBlock(lines-6, 2, 5, 10, sh)
	b2.SetFrame('O', true, common.EnumColor.Blue)
	mgr.AddBlock(b2)

	for i := 0; i < 20; i++ {
		dataSrc.Str += "A"
		time.Sleep(300 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
	c.Reset()
	c.ClearAll()
	c.SetCursorPos(0, 0)
}
