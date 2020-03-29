package mojo

import (
	"errors"
	"image"

	"github.com/fogleman/gg"
)

type Scene struct {
	Width    uint32
	Height   uint32
	Children []Node
}

func NewScene(w uint32, h uint32) *Scene {
	return &Scene{
		Width:    w,
		Height:   h,
		Children: []Node{},
	}
}

func (s *Scene) Render() (image.Image, error) {
	// create the drawing context
	dc := NewDrawContext(GGBackend, s.Width, s.Height)

	// recursively render all nodes
	s.renderNodes(dc, s.Children)

	// render all children in-order
	for _, c := range s.Children {
		c.Render(dc)
	}

	// return the composite image
	return dc.Image()
}

func (s *Scene) renderNodes(dc *DrawContext, nodes []Node) error {
	// validate params
	if dc == nil {
		return errors.New("Invalid nil DrawContext given; cannot render nodes")
	}
	if nodes == nil {
		return errors.New("Invalid nil nodes given; cannot render nodes")
	}

	// recursively render nodes (breadth-first)
	for _, n := range nodes {
		n.Render(dc)
		s.renderNodes(dc, n.GetChildren())
	}
	return nil
}

func (s *Scene) RenderPNG(path string) error {
	// render the scene to an image
	im, err := s.Render()
	if err != nil {
		return err
	}

	// save the rendered image to a png file and return any errors
	return gg.SavePNG(path, im)
}
