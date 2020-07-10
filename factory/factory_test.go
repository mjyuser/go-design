package factory

import "testing"

func TestNewFactory(t *testing.T) {
	factory := NewFactory("+")
	plus := factory.Create()
	plus.SetA(3)
	plus.SetB(4)
	if plus.Result() != 7 {
		t.Errorf("plus error, res: %v", plus.Result())
	}

	minusFactory := NewFactory("-")
	minus := minusFactory.Create()
	minus.SetA(5)
	minus.SetB(4)
	if minus.Result() != 1 {
		t.Errorf("minus error, res: %v", minus.Result())
	}

	plusF := PlusOperatorFactory{}
	r1 := compute(plusF, 1, 2)
	if r1 != 3 {
		t.Errorf("plus compute failed, res: %v", r1)
	}
	minusF := MinusOperatorFactory{}
	r2 := compute(minusF, 1, 3)
	if r2 != -2 {
		t.Errorf("minus compute failed, res: %v", r2)
	}
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}


