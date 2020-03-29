package mojo

var (
	defaultColor  = NewRGBA(0, 0, 0, 1)
	defaultStroke = &Stroke{
		Width: 1,
		Color: defaultColor,
	}
	defaultFile = &Fill{
		Color: defaultColor,
	}
)
