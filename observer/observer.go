package observer

import (
	"fmt"
	"time"
)

type Observer interface {
	Notify(data interface{})
}

type EventSubject struct {
	listener map[EventType][]Observer
}

type EventType string

const (
	ReportNews EventType = "report_news"
	ReportSportsNews EventType = "report_sports_news"
	AllNews = "*"
)

func (e *EventSubject) NotifyObserver(eventType EventType, data interface{}) {
	if e.listener == nil {
		e.listener = make(map[EventType][]Observer)
	}
	listeners, ok := e.listener[eventType]
	if !ok {
		return
	}
	for _, listener := range listeners {
		listener.Notify(data)
	}

}

func (e *EventSubject) Subscribe(eventType EventType, observer Observer) {
	if e.listener == nil {
		e.listener = make(map[EventType][]Observer)
	}

	if e.listener[AllNews] == nil {
		e.listener[AllNews] = make([]Observer, 0)
	}

	if e.listener[eventType] == nil {
		e.listener[eventType] = make([]Observer, 0)
	}

	e.listener[AllNews] = append(e.listener[AllNews], observer)
	e.listener[eventType] = append(e.listener[eventType], observer)
}

func (e *EventSubject) DeSubscribe(eventType EventType, observer Observer) {
	if e.listener == nil {
		return
	}

	if e.listener[eventType] == nil {
		return
	}

	for i, o := range e.listener[eventType] {
		if o == observer {
			e.listener[eventType] = append(e.listener[eventType][0:i-1], e.listener[eventType][i:]...)
			if len(e.listener[eventType]) == 0 {
				delete(e.listener, eventType)
			}
		}
	}
}

type Subject interface {
	Subscribe(eventType EventType, observer Observer)
	DeSubscribe(eventType EventType, observer Observer)
	NotifyObserver(eventType EventType, data interface{})
}

type PeoplesDaily struct {
	Title string
	EventSubject
}

func (p *PeoplesDaily) Report() {
	content := fmt.Sprintf("people's daily report message. [%s] \n", time.Now().Format("2006-01-02 15:04:05"))
	p.NotifyObserver(ReportNews, content)
}

type People struct {
	Name string
}

func NewPeople(name string) *People {
	return &People{
		Name: name,
	}
}

func (p *People) Notify(data interface{}) {
	fmt.Printf("user: [%s] receive subject data: %v\n", p.Name, data)
}

