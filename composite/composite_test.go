package composite

import (
	"fmt"
	"testing"
)

func Test_Composite(t *testing.T) {
	form := NewForm("product", "add product", "/product/add")
	form.Add(NewInput("name", "Name", "text"))
	form.Add(NewInput("description", "Description", "text"))
	fmt.Println(form.fields)
	picture := NewFieldSet("photo", "product photo")
	picture.Add(NewInput("caption", "Caption", "text"))
	picture.Add(NewInput("image", "Image", "file"))
	form.Add(picture)
	LoadProductData(form)
	t.Log(form.Render())
}

func LoadProductData(element FormElement) {
	data := map[string]interface{}{
		"name": "Apple MacBook",
		"description": "A decent laptop.",
		"photo": map[string]interface{}{
			"caption": "Front photo",
			"image": "photo1.png",
		},
	}

	element.SetData(data)
}
