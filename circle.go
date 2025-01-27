package mojo

type Circle struct {
	Shape
	X float64
	Y float64
	R float64
	*Stroke
	*Fill
}

func (c Circle) GetChildren() []Node {
	return []Node{}
}

func (c Circle) Render(dc *DrawContext) error {
	// draw circle and fill
	ggdc := dc.ggCtx
	ggdc.DrawCircle(c.X, c.Y, c.R)
	color := c.Fill.Color
	ggdc.SetRGBA(color.R, color.G, color.B, color.A)
	ggdc.FillPreserve()

	// draw stroke
	color = c.Stroke.Color
	ggdc.SetLineWidth(c.Stroke.Width)
	ggdc.SetRGBA(color.R, color.G, color.B, color.A)
	ggdc.Stroke()
	return nil
}
