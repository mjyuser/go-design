package observer

import "testing"

func Test_Observer(t *testing.T) {
	zhang := NewPeople("zhang")
	wang := NewPeople("wang")

	daily := &PeoplesDaily{}
	daily.Subscribe(ReportNews, zhang)
	daily.Subscribe(ReportSportsNews, wang)

	daily.NotifyObserver(ReportNews, "hello world")
	daily.NotifyObserver(ReportSportsNews, "report sports news")
	daily.NotifyObserver(AllNews,"report all news")
}
