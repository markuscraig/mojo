package mojo

var (
	defaultColor = NewRGBA(0, 0, 0, 1)

	defaultStroke = &Stroke{
		Width: 1,
		Color: defaultColor,
	}

	defaultFile = &Fill{
		Color: defaultColor,
	}

	defaultXY = XY{
		X: 0,
		Y: 0,
	}

	defaultRotation = Rotation{
		Times:   0,
		Degrees: 0,
	}

	defaultOpacity = float64(1)

	defaultTransform = &Transform{
		AnchorPoint:  defaultXY,
		Position:     defaultXY,
		Scale:        defaultXY,
		Rotation:     defaultRotation,
		StartOpacity: defaultOpacity,
		EndOpacity:   defaultOpacity,
	}
)
