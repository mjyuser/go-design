package prototype

import "testing"

func TestNewMessage(t *testing.T) {
	message := NewMessage("hello", "world", []Tag{"origin"})

	newMessage := message.Clone()

	nn,ok := newMessage.(*Message)
	if !ok {
		t.Errorf("newmessage shoule be message type")
	}

	nn.SetTitle("gogogo")

	if message.Title == nn.Title {
		t.Errorf("tiitle should not be equal")
	}

	t.Log(message)
	t.Log(nn)
}
