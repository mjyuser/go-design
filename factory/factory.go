package factory

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(num int) {
	o.a = num
}
func (o *OperatorBase) SetB(num int) {
	o.b = num
}


type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

type PlusOperatorFactory struct {}
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

type MinusOperatorFactory struct {}
func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}

func NewFactory(typ string) OperatorFactory{
	switch typ {
	case "+":
		return &PlusOperatorFactory{}
	case "-":
		return &MinusOperatorFactory{}
	}

	return nil
}

func AbstractFactory() OperatorFactory {
	var factory OperatorFactory
	factory = &PlusOperatorFactory{}
	compute := func(factory OperatorFactory, a, b int) int {
		op := factory.Create()
		op.SetA(a)
		op.SetB(b)
		return op.Result()
	}

	_ = compute(factory, 1, 2)

	return factory
}