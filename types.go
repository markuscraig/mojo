package mojo

type Node interface {
	GetChildren() []Node
	Render(*DrawContext) error
}

type Shape interface {
	Node
}

type Stroke struct {
	Width float64
	Color RGBA
}

type Fill struct {
	Color RGBA
}

type XY struct {
	X float64
	Y float64
}

type Rotation struct {
	Times   uint32
	Degrees float64
}
