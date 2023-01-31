package show

import (
	"fmt"

	"github.com/wormggmm/gocmd/common"
)

// TODO:
// 1. 封装print函数，对换行的支持，以及超过width后的折行
type Block struct {
	controller *common.ScreenController
	row        int
	col        int
	height     int
	width      int
	cursorPos  *common.Pos
}

func NewBlock(row, col, height, width int) *Block {
	return &Block{
		row:        row,
		col:        col,
		height:     height,
		width:      width,
		controller: &common.ScreenController{},
		cursorPos:  &common.Pos{},
	}
}
func (s *Block) globalCursorPosY() int {
	return s.row + s.cursorPos.Y
}
func (s *Block) globalCursorPosX() int {
	return s.col + s.cursorPos.X
}
func (s *Block) applyCursorPos() {
	s.controller.SetCursorPos(s.globalCursorPosY(), s.globalCursorPosX())
}
func (s *Block) Printf(format string, argv ...interface{}) {
	s.applyCursorPos()
	str := fmt.Sprintf(format, argv...)
	subStr := ""
	for _, c := range str {
		if c == '\n' || s.cursorPos.X+len(subStr)+1 > s.width {
			fmt.Print(subStr)
			subStr = ""
			s.cursorPos.Y++
			s.cursorPos.X = 0
			s.applyCursorPos()
			if c != '\n' {
				subStr = string(c)
			}
		} else {
			subStr += string(c)
		}
	}
	if len(subStr) > 0 {
		fmt.Print(subStr)
		s.cursorPos.X += len(subStr)
	}
}
