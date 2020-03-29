package mojo

const (
	defaultRepeaterCopies = uint32(3)
)

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
	Transform *Transform
	Children  []Node
}

func (r Repeater) GetChildren() []Node {
	return r.Children
}

func (r Repeater) Render(dc *DrawContext) error {
	// use default properties if needed
	if r.Copies == 0 {
		r.Copies = defaultRepeaterCopies
	}

	return nil
}
