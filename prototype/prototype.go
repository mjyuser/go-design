package prototype

import "time"

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}




// 消息模版
type Message struct {
	Title string
	Content string
	CreateAt time.Time

	Tags []Tag
}

type Tag string


func (m *Message) Clone() Cloneable {
	mm := *m
	mm.CreateAt = time.Now().Add(time.Minute)

	return &mm
}

func NewMessage(title string, content string, tags []Tag) *Message {
	return &Message{
			Title: title,
			Content: content,
			CreateAt: time.Now(),
			Tags: tags,
	}
}

func (m *Message) SetTitle(title string) {
	m.Title = title
}

func (m *Message) SetContent(content string) {
	m.Content = content
}

func (m *Message) SetTags(tags []Tag) {
	m.Tags = tags
}

