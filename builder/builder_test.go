package builder

import "testing"

func TestNewDirect(t *testing.T) {
	car := &Car{}
	direct := NewDirect(car)
	direct.Construct()
	product := direct

	t.Log(product)
}
