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

type Transform struct {
	AnchorPoint  XY
	Position     XY
	Scale        XY
	Rotation     Rotation
	StartOpacity float64
	EndOpacity   float64
}
