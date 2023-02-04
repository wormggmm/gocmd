package comp

import (
	"fmt"
	"strings"

	"github.com/wormggmm/gocmd/common"
)

type ProgressBar struct {
	tag       string
	char      rune
	charEmpty rune
	length    int
	progress  float64
}

func NewProgressBar(tag string, char rune, length int, charEmpty ...rune) *ProgressBar {
	emptyChar := ' '
	if len(charEmpty) >= 1 {
		emptyChar = charEmpty[0]
	}
	pb := &ProgressBar{
		tag,
		char,
		emptyChar,
		length,
		0,
	}
	return pb
}
func (s *ProgressBar) Listener(listener common.IDataListener) {}
func (s *ProgressBar) Set(progress float64) {
	s.progress = progress
}
func (s *ProgressBar) Data() string {
	c := int(float64(s.length) * s.progress)
	str := fmt.Sprintf("%s%s (%f)",
		s.tag, strings.Repeat(string(s.char), c)+strings.Repeat(string(s.charEmpty), s.length-c), s.progress)
	return str
}
