package common

import (
	"fmt"
)

type TColorEnum int
type colorEnum struct {
	Idle   TColorEnum
	Black  TColorEnum
	Red    TColorEnum
	Green  TColorEnum
	Yellow TColorEnum
	Blue   TColorEnum
	Purple TColorEnum
	Ching  TColorEnum
	White  TColorEnum
}

var EnumColor = colorEnum{
	Idle:   -1,
	Black:  0,
	Red:    1,
	Green:  2,
	Yellow: 3,
	Blue:   4,
	Purple: 5,
	Ching:  6,
	White:  7,
}

type ScreenController struct {
	BaseController
}

func (s *ScreenController) ClearLine() {
	s.MoveCursorToLineBegin()
	s.ClearCursorToEnd()
}
func (s *ScreenController) MoveCursor(row, col int) {
	if row < 0 {
		row = -row
		s.MoveCursorLeft(row)
	} else {
		s.MoveCursorRight(row)
	}
	if col < 0 {
		col = -col
		s.MoveCursorUp(col)
	} else {
		s.MoveCursorDown(col)
	}
}

type BaseController struct {
}

func (s *BaseController) MoveCursorToLineBegin() {
	fmt.Print("\r")
}
func (s *BaseController) ClearCursorToEnd() {
	fmt.Print("\033[K")
}
func (s *BaseController) ClearAll() {
	fmt.Print("\033[2J")
}
func (s *BaseController) SaveCursorPos() {
	fmt.Print("\033[s")
}
func (s *BaseController) LoadCursorPos() {
	fmt.Print("\033[u")
}
func (s *BaseController) SetCursorPos(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}
func (s *BaseController) MoveCursorLeft(num int) {
	fmt.Printf("\033[%dD", num)
}
func (s *BaseController) MoveCursorRight(num int) {
	fmt.Printf("\033[%dC", num)
}
func (s *BaseController) MoveCursorUp(num int) {
	fmt.Printf("\033[%dA", num)
}
func (s *BaseController) MoveCursorDown(num int) {
	fmt.Printf("\033[%dB", num)
}
func (s *BaseController) FontColor(color TColorEnum) {
	fmt.Printf("\033[3%dm", color)
}
func (s *BaseController) FontBackColor(color TColorEnum) {
	fmt.Printf("\033[4%dm", color)
}
func (s *BaseController) Blink() {
	fmt.Print("\033[32;5m")
}
func (s *BaseController) Reset() {
	fmt.Print("\033[0m")
}
