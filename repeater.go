package mojo

type Repeater struct {
	Node
	Children []Node
}

func (r Repeater) GetChildren() []Node {
	return r.Children
}

func (r Repeater) Render(dc *DrawContext) error {
	return nil
}
