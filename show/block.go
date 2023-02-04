package show

import (
	"fmt"

	"github.com/wormggmm/gocmd/utils"

	"github.com/wormggmm/gocmd/common"
)

// 封装print函数，对换行的支持，以及超过width后的折行
type Block struct {
	controller *common.ScreenController
	manager    *Manager
	id         int64
	row        int
	col        int
	height     int
	width      int
	cursorPos  *common.Pos
	dataSrc    common.IDataSource
	frameChar  rune
	frameBlink bool
	frameColor common.TColorEnum
}

func NewBlock(row, col, height, width int, dataSrc common.IDataSource) *Block {
	b := &Block{
		id:         utils.GenUUID().Int64(),
		row:        row,
		col:        col,
		height:     height,
		width:      width,
		controller: &common.ScreenController{},
		cursorPos:  &common.Pos{},
		dataSrc:    dataSrc,
		frameColor: common.EnumColor.Idle,
	}
	dataSrc.Listener(b)
	return b
}
func (s *Block) SetManager(manager *Manager) {
	s.manager = manager
}
func (s *Block) DataChanged() {
	s.manager.FlushDraw(s)
}
func (s *Block) ID() int64 {
	return s.id
}
func (s *Block) setDataSource(dataSrc common.IDataSource) {
	s.dataSrc = dataSrc
}
func (s *Block) setCursorPosY(y int) {
	s.cursorPos.Y = y
	s.applyCursorPos()
}
func (s *Block) setCursorPosX(x int) {
	s.cursorPos.X = x
	s.applyCursorPos()
}
func (s *Block) setCursorPos(x, y int) {
	s.cursorPos.X = x
	s.cursorPos.Y = y
	s.applyCursorPos()
}
func (s *Block) globalCursorPosY() int {
	return s.row + s.cursorPos.Y
}
func (s *Block) globalCursorPosX() int {
	return s.col + s.cursorPos.X
}

func (s *Block) ApplyCursorPos() {
	s.applyCursorPos()
}
func (s *Block) applyCursorPos() {
	s.controller.SetCursorPos(s.globalCursorPosY(), s.globalCursorPosX())
}
func (s *Block) SetFrame(char rune, blink bool, color common.TColorEnum) {
	s.frameChar = char
	s.frameBlink = blink
	s.frameColor = color
}
func (s *Block) Clear() {
	s.setCursorPos(0, 0)
	bankStr := ""
	for i := 0; i < s.height*s.width; i++ {
		bankStr += " "
	}
	s.Printf(bankStr)
	s.setCursorPos(0, 0)
}
func (s *Block) drawFrame() {
	if s.frameChar == 0 {
		return
	}
	s.setCursorPos(-1, -1)
	frameWidthStr := ""
	frameWidthStr2 := ""
	for i := 0; i < s.width+2; i++ {
		frameWidthStr += string(s.frameChar)
		if len(frameWidthStr2) == 0 || len(frameWidthStr2) == s.width+1 {
			frameWidthStr2 += string(s.frameChar)
		} else {
			frameWidthStr2 += " "
		}
	}
	if s.frameBlink {
		s.controller.Blink()
	}
	if s.frameColor != common.EnumColor.Idle {
		s.controller.FontColor(s.frameColor)
	}
	fmt.Print(frameWidthStr)
	for i := 0; i < s.height; i++ {
		s.setCursorPosY(s.cursorPos.Y + 1)
		fmt.Print(frameWidthStr2)
	}
	s.setCursorPosY(s.cursorPos.Y + 1)
	fmt.Println(frameWidthStr)
	s.controller.Reset()
}
func (s *Block) Draw() {
	if s.dataSrc == nil {
		return
	}
	s.Clear()
	s.drawFrame()
	s.setCursorPos(0, 0)
	s.Printf(s.dataSrc.Data())
}
func (s *Block) Printf(format string, argv ...interface{}) {
	s.applyCursorPos()
	str := fmt.Sprintf(format, argv...)
	subStr := ""
	for _, c := range str {
		if c == '\n' || s.cursorPos.X+len(subStr)+1 > s.width {
			fmt.Print(subStr)
			subStr = ""
			s.setCursorPos(0, s.cursorPos.Y+1)
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
