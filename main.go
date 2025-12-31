package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func main() {
	gomesengine.Init("Gaia", 800, 600)

	game()

	gomesengine.Run()
}

func game() {
	settings()
	scene()
}

func settings() {
	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})
}

func scene() {
	var grid = math.Vector2{X: 10, Y: 10}
	var rect = utils.RectSpecs{
		PosX:   0,
		PosY:   0,
		Width:  32,
		Height: 32,
	}

	for x := 0; x < grid.X; x++ {
		for y := 0; y < grid.Y; y++ {
			lifecycle.Register(&lifecycle.GameObject{
				Render: func() {
					r := rect
					r.PosX = x * (rect.Width + 8)
					r.PosY = y * (rect.Height + 8)

					render.DrawRect(r, render.Yellow)
				},
			})
		}
	}
}
