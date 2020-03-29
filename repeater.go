package mojo

type CompositeType uint8

const (
	Above CompositeType = iota
	Below
)

type Repeater struct {
	Node
	Copies    uint32
	Offset    float64
	Composite CompositeType
	Children  []Node
}

type RepeatTransform struct {
	AnchorPoint  XY
	Position     XY
	Scale        XY
	Rotation     Rotation
	StartOpacity float64
	EndOpacity   float64
}

func (r Repeater) GetChildren() []Node {
	return r.Children
}

func (r Repeater) Render(dc *DrawContext) error {
	return nil
}
