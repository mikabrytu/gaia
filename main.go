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
	grid := math.Vector2{X: 10, Y: 10}

	for x := 0; x < grid.X; x++ {
		for y := 0; y < grid.Y; y++ {
			tile(x, y)
		}
	}
}

func tile(x int, y int) {
	size := 32
	offset := 8

	rect := utils.RectSpecs{
		PosX:   x * (size + offset),
		PosY:   y * (size + offset),
		Width:  size,
		Height: size,
	}

	color := render.Yellow

	lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			events.Subscribe(events.Input, events.INPUT_MOUSE_CLICK_DOWN, func(data any) {
				click := data.(events.InputMouseClickDownEvent)

				if (click.Position.X > rect.PosX && click.Position.X < (rect.PosX+rect.Width)) && (click.Position.Y > rect.PosY && click.Position.Y < (rect.PosY+rect.Height)) {
					color = render.Green
				}
			})
		},
		Render: func() {
			render.DrawRect(rect, color)
		},
	})
}
