package builder

import "fmt"

// 创建一组具有相同部件的不同物品, 通过direct的引导过程进行实例化



type CarBuilder interface {
	SetSeats()
	SetEngine()
	SetTripComputer()
	SetGPS()
	GetProduct() Product
}

type Direct interface {
	SetBuilder(builder CarBuilder)
	Construct()
	GetProduct() Product
}

type Product interface {
	GetProduct() Product
}
func (c *Car) GetProduct() Product {
	return c
}

type Car struct {
	Seats string
	Engine string
	TripComputer string
	GPS string
}

func (c *Car) SetSeats() {
	fmt.Println("set car seats")
}
func (c *Car) SetEngine() {
	fmt.Println("set car engine")
}
func (c *Car) SetTripComputer() {
	fmt.Println("set car trip computer")
}
func (c *Car) SetGPS() {
	fmt.Println("set car gps")
}

type CarDirect struct {
	builder CarBuilder
}

func (c *CarDirect) SetBuilder(builder CarBuilder) {
	c.builder = builder
}

func NewDirect(builder CarBuilder) Direct {
	direct := &CarDirect{}
	direct.SetBuilder(builder)

	return direct
}

func (c CarDirect) GetProduct() Product {
	return c.builder.GetProduct()
}

func (c CarDirect) Construct() {
	c.builder.SetSeats()
	c.builder.SetEngine()
	c.builder.SetGPS()
	c.builder.SetTripComputer()
}
