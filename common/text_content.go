package common

type TextContent struct {
	lines       []string
	currentLine string
}

func (s *TextContent) InputChar(char rune) {
	if char == '\n' {
		s.EnterLine()
	} else {
		s.currentLine += string(char)
	}
}
func (s *TextContent) EnterLine() {
	s.lines = append(s.lines, s.currentLine)
	s.currentLine = ""
}

func (s *TextContent) BackSpace() {
	if len(s.currentLine) > 0 {
		s.currentLine = s.currentLine[0 : len(s.currentLine)-1]
	}
}

func (s *TextContent) CurrentLineLen() int {
	return len(s.currentLine)
}

func (s *TextContent) CurrentLine() string {
	return s.currentLine
}

func (s *TextContent) SetCurrentLine(content string) {
	s.currentLine = content
}
func (s *TextContent) LinesData() []string {
	lines := append(s.lines, s.currentLine)
	return lines
}
