package composite

import (
	"fmt"
	"strings"
)

// 组合模式

type FormElement interface {
	GetName() string
	SetData(data interface{})
	GetData() interface{}

	AbstractFormElement
}

type AbstractFormElement interface {
	Render() string
}

type Element struct {
	Name string
	Title string
	Data interface{}
}

func (e *Element) GetName() string {
	return e.Name
}

func (e *Element) SetData(data interface{}) {
	e.Data = data
}

func (e *Element) GetData() interface{} {
	return e.Data
}

func (e *Element) Render() string {
	fmt.Println("exec element render")
	return ""
}


type Composite interface {
	FormElement
	Add(field FormElement) FormElement
	Remove(field FormElement) FormElement
}

type FieldComposite struct {
	fields map[string]FormElement
	FormElement
}

func (f *FieldComposite) Add(element FormElement) FormElement {
	name := element.GetName()
	f.fields[name] = element
	if name == "picture" {
		fmt.Println(f.Render())
	}
	return f
}

func (f *FieldComposite) Remove(element FormElement) FormElement {
	if _, ok := f.fields[element.GetName()]; ok {
		delete(f.fields, element.GetName())
	}

	return f
}

func (f *FieldComposite) SetData(data interface{}) {
	d := data.(map[string]interface{})
	for key, val := range d {
		if element, ok := f.fields[key]; ok {
			element.SetData(val)
		}
	}
}
func (f *FieldComposite) GetData() interface{} {
	var data []interface{}
	for _, v := range f.fields {
		data = append(data, v.GetData())
	}

	return data
}

func (f *FieldComposite) Render() string {
	var output strings.Builder
	for _, v := range f.fields {
		output.WriteString(v.Render())
	}

	return output.String()
}

type FieldSet struct {
	*FieldComposite
}

func (f *FieldSet) Render() string {
	element := f.FormElement.(*Element)
	output := f.FieldComposite.Render()
	return fmt.Sprintf("<fieldset><legend>%s</legend>\n%s</fieldset>\n", element.Title, output)
}

func NewFieldSet(name, title string) *FieldSet {
	return &FieldSet{
		&FieldComposite{
			fields:      make(map[string]FormElement),
			FormElement: &Element{
				Name:  name,
				Title: title,
			},
		},
	}
}


type Input struct {
	FormElement
	InputType string
}

func (i *Input) Render() string {
	element := i.FormElement.(*Element)
	return fmt.Sprintf("<label for=\"%s\">%s</label>\n" +
		"<input name=\"%s\" type=\"%s\" value=\"%v\">\n", element.Name, element.Title, element.Name, i.InputType, element.Data)
}

func NewInput(name, title, inputType string) *Input {
	return &Input{
		FormElement: &Element{
			Name:        name,
			Title:       title,
		},
		InputType:   inputType,
	}
}


type Form struct {
	Url string
	*FieldSet
}

func (f *Form) Render() string {
	output := f.FieldComposite.Render()
	element := f.FormElement.(*Element)
	return fmt.Sprintf("<form action=\"%s\">\n<h3>%s</h3>\n%s</form>\n", f.Url, element.Title, output)
}

func NewForm(name, title, url string) *Form {
	return &Form{
		Url:            url,
		FieldSet: NewFieldSet(name, title),
	}
}