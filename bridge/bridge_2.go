package bridge

import (
	"fmt"
	"strings"
)

type Renderer interface {
	RenderTitle(title string) string
	RenderTextBlock(text string) string
	RenderImage(url string) string
	RenderLink(url, title string) string
	RenderHeader() string
	RenderFooter() string
	RenderParts(parts []string) string
}

type Page interface {
	View() string
	ChangeRenderer(renderer Renderer)
}

type SimplePage struct {
	Title string
	Content string
	renderer Renderer
}

func (s *SimplePage) ChangeRenderer(renderer Renderer) {
	s.renderer = renderer
}

func NewSimplePage(renderer Renderer, title, content string) Page {
	return &SimplePage{
		Title:    title,
		Content:  content,
		renderer: renderer,
	}
}

func (s *SimplePage) View() string {
	return s.renderer.RenderParts([]string{
		s.renderer.RenderHeader(),
		s.renderer.RenderTitle(s.Title),
		s.renderer.RenderTextBlock(s.Content),
		s.renderer.RenderFooter(),
	})
}

type Product struct {
	Id string
	Title string
	Description string
	Image string
	Price float64
}

func NewProduct(id, title, description, image string, price float64) *Product {
	return &Product{
		Id: id,
		Title: title,
		Description: description,
		Image: image,
		Price: price,
	}
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetTitle() string {
	return p.Title
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetImage() string {
	return p.Image
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

type ProductInterface interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetImage() string
	GetPrice() float64
}

type ProductPage struct {
	product *Product
	renderer Renderer
}

func (p *ProductPage) ChangeRenderer(renderer Renderer) {
	p.renderer = renderer
}

func NewProductPage(product *Product, renderer Renderer) Page {
	return &ProductPage{
		product: product,
		renderer: renderer,
	}
}

func (p *ProductPage) View() string {
	return p.renderer.RenderParts([]string{
		p.renderer.RenderHeader(),
		p.renderer.RenderTitle(p.product.GetTitle()),
		p.renderer.RenderTextBlock(p.product.GetDescription()),
		p.renderer.RenderImage(p.product.GetImage()),
		p.renderer.RenderLink(fmt.Sprintf("%s%s", "/cart/add", p.product.GetId()), "Add to cart"),
		p.renderer.RenderFooter(),
	})
}

type HTMLRenderer struct {}

func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{}
}

func (H *HTMLRenderer) RenderTitle(title string) string {
	return fmt.Sprintf("<h1>%s</h1>", title)
}


func (H *HTMLRenderer) RenderTextBlock(text string) string {
	return fmt.Sprintf("<div class='text'>%s</div>", text)
}

func (H *HTMLRenderer) RenderImage(url string) string {
	return fmt.Sprintf("<img src='%s'>", url)
}

func (H *HTMLRenderer) RenderLink(url, title string) string {
	return fmt.Sprintf("<a href='%s'>%s</a>", url, title)
}

func (H *HTMLRenderer) RenderHeader() string {
	return "<html><body>"
}

func (H *HTMLRenderer) RenderFooter() string {
	return "</body></html>"
}

func (H *HTMLRenderer) RenderParts(parts []string) string {
	return strings.Join(parts, "\n")
}



type JsonRenderer struct {}

func NewJsonRenderer() *JsonRenderer {
	return &JsonRenderer{}
}

func (j *JsonRenderer) RenderTitle(title string) string {
	return fmt.Sprintf("'\"title\": \"%s\"'", title)
}

func (j *JsonRenderer) RenderTextBlock(text string) string {
	return fmt.Sprintf("'\"text\": \"%s\"'", text)
}

func (j *JsonRenderer) RenderImage(url string) string {
	return fmt.Sprintf("'\"img\": \"%s\"'", url)
}

func (j *JsonRenderer) RenderLink(url, title string) string {
	return fmt.Sprintf("'\"link\": {\"href\": \"%s\", \"title\": \"%s\"\"}'", url, title)
}

func (j *JsonRenderer) RenderHeader() string {
	return ""
}

func (j *JsonRenderer) RenderFooter() string {
	return ""
}

func (j *JsonRenderer) RenderParts(parts []string) string {
	return fmt.Sprintf("{\n\"%s\"\n}", strings.Join(parts, ",\n"))
}



