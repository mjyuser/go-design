package decorator

import "testing"

func Test_Decorator(t *testing.T) {
	format := NewTextFormat(&TextInput{})
	format = WrapPlainTextFilter(format)
	format = WrapDangerousHTMLTagFilter(format)
	format.FormatText("hello world")
}
