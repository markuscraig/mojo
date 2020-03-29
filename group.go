package mojo

import (
	"errors"
)

type Group struct {
	Node
	Children  []Node
	Transform *Transform
}

func NewGroup() Group {
	return Group{
		Children:  []Node{},
		Transform: defaultTransform,
	}
}

func (g Group) GetChildren() []Node {
	return g.Children
}

func (g Group) Render(dc *DrawContext) error {
	// recursively render all nodes
	g.renderNodes(dc, g.Children)
	return nil
}

func (g Group) renderNodes(dc *DrawContext, nodes []Node) error {
	// validate params
	if dc == nil {
		return errors.New("Invalid nil DrawContext given; cannot render group")
	}
	if nodes == nil {
		return errors.New("Invalid nil nodes given; cannot render group")
	}

	// recursively render nodes (breadth-first)
	for _, n := range nodes {
		n.Render(dc)
		g.renderNodes(dc, n.GetChildren())
	}
	return nil
}
