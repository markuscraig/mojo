package mojo

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
)

type DrawBackend int

const (
	GGBackend DrawBackend = iota
)

type DrawContext struct {
	backend DrawBackend
	ggCtx   *gg.Context
}

func NewDrawContext(b DrawBackend, width uint32, height uint32) *DrawContext {
	// create and init the gg drawing context
	dc := gg.NewContext(int(width), int(height))

	// create and return the mojo drawing context
	return &DrawContext{
		backend: b,
		ggCtx:   dc,
	}
}

func (dc *DrawContext) Image() (image.Image, error) {
	switch dc.backend {
	case GGBackend:
		return dc.ggCtx.Image(), nil
	}
	return nil, fmt.Errorf("Unsupported drawing backend '%v'; cannot render image", dc.backend)
}
