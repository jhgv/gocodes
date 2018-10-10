package objects

import (
	"bytes"
	"strings"
)

type TextHelper struct {
}

func (th *TextHelper) UpperText(text string) string {
	return strings.ToUpper(text)
}

func (th *TextHelper) LowerText(text string) string {
	return strings.ToLower(text)
}

func (th *TextHelper) InvertText(text string) string {
	var buf bytes.Buffer
	for pos := range text {
		buf.WriteString(string(text[len(text)-pos-1]))
	}
	return buf.String()
}
