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
