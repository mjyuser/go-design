package bridge

import "testing"

func Test_Bridge(t *testing.T) {
	htmlRenderer := NewHTMLRenderer()
	jsonRenderer := NewJsonRenderer()

	page := NewSimplePage(htmlRenderer, "Home", "Welcome to our website")
	t.Log("HTML view of a simple content page:\n")
	t.Log(page.View())

	product := NewProduct("123", "Star Wars, episode1", "A long time ago in a galaxy far, far away...", "/images/star-wars.jpeg", 128.1)
	page = NewProductPage(product, htmlRenderer)

	t.Log("HTML view of a product page, same client code:\n")
	t.Log(page.View())

	page.ChangeRenderer(jsonRenderer)
	t.Log("JSON view of a simple content page, with the same client code:\n")
	t.Log(page.View())
}
