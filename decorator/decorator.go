package decorator

import "fmt"

type InputFormat interface {
	FormatText(text string) string
}


type TextInput struct {}

func (t *TextInput) FormatText(text string) string {
	fmt.Println("text input format")
	return text
}

type TextFormat struct {
	format InputFormat
}

func NewTextFormat(format InputFormat) InputFormat {
	return &TextFormat{
		format: format,
	}
}

func (t *TextFormat) FormatText(text string) string {
	return t.format.FormatText(text)
}

type PlainTextFilter struct {
	format InputFormat
}

func (p *PlainTextFilter) FormatText(text string) string {
	fmt.Println("plain text filter")
	text = p.format.FormatText(text)
	return text
}


type DangerousHTMLTagFilter struct {
	format InputFormat
}

func (d *DangerousHTMLTagFilter) FormatText(text string) string {
	fmt.Println("dangerous html tag filter")
	text = d.format.FormatText(text)
	return text
}

func WrapDangerousHTMLTagFilter(format InputFormat) InputFormat {
	return &DangerousHTMLTagFilter{format: format}
}

func WrapPlainTextFilter(format InputFormat) InputFormat {
	return &PlainTextFilter{format: format}
}

