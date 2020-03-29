package main

import (
	"github.com/markuscraig/mojo"
)

func main() {
	// build the scene
	scene := &mojo.Scene{
		Width:  800,
		Height: 600,
		Children: []mojo.Node{
			mojo.Repeater{
				Children: []mojo.Node{
					mojo.Circle{
						X: 100.0,
						Y: 200.0,
						R: 50.0,
						Stroke: mojo.Stroke{
							Width: 10.0,
							Color: mojo.NewRGBA(1, 0, 0, 1),
						},
						Fill: mojo.Fill{
							Color: mojo.NewRGBA(0, 0, 1, 1),
						},
					},
				},
			},
		},
	}

	// render to a png image
	scene.RenderPNG("out.png")
}
