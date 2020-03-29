package mojo

type RGBA struct {
	R float64
	G float64
	B float64
	A float64
}

func NewRGBA(r, g, b, a float64) RGBA {
	return RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
