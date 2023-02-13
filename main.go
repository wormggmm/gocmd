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
	linesStr := "50"
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
	// dataSrc := &TestDataSource{}
	// b1 := show.NewBlock(2, 2, 5, 10, dataSrc)
	// b1.SetFrame('*', true, common.EnumColor.Red)
	mgr := show.NewManager()
	// mgr.AddBlock(b1)
	mgr.Start()
	mgr.SetCursorPos(0, 0)
	input := shell.NewInput()
	sh := shell.NewShell()
	testShellProcessor := &TestShellProcessor{cmd: func(cmd string) {
		if cmd == "exit" {
			input.Stop()
			os.Exit(0)
		}
	},
		cmdList: func() []string {
			return []string{"exit"}
		},
	}
	sh.SetReceiver(testShellProcessor)
	input.SetReceiver(sh)
	b2 := show.NewBlock(lines-6, 2, 5, 10, sh)
	b2.SetFrame('O', false, common.EnumColor.Blue)
	mgr.AddBlock(b2)
	// pb := comp.NewProgressBar("test:", '#', 30, '-')
	// b3 := show.NewBlock(2+5+2, 2, 1, 50, pb)
	// mgr.AddBlock(b3)
	input.Start()
	defer input.Stop()
	for i := 0; i < 20; i++ {
		// pb.Set(float64((1+i)*5) / 100)
		// dataSrc.Str += "A"
		time.Sleep(300 * time.Millisecond)
	}
	time.Sleep(50 * time.Second)
	c.Reset()
	c.ClearAll()
	c.SetCursorPos(0, 0)
}

type TestShellProcessor struct {
	cmd      func(cmd string)
	cmdTable func(cmd string)
	cmdList  func() []string
}

func (s *TestShellProcessor) Cmd(cmd string) {
	s.cmd(cmd)
}

func (s *TestShellProcessor) CmdTable(cmd string) {
	s.cmdTable(cmd)
}

func (s *TestShellProcessor) CmdList() []string {
	return s.cmdList()
}
