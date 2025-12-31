package main

import (
	"math"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	gomesmath "github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var player_tiles []gomesmath.Vector2

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
	grid := gomesmath.Vector2{X: 10, Y: 10}

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
					if try_add_tile(gomesmath.Vector2{X: x, Y: y}) {
						color = render.Green
					}
				}
			})
		},
		Render: func() {
			render.DrawRect(rect, color)
		},
	})
}

func try_add_tile(position gomesmath.Vector2) bool {
	if player_tiles == nil {
		player_tiles = make([]gomesmath.Vector2, 1)
		player_tiles[0] = position
		return true
	}

	for _, tile := range player_tiles {
		if math.Abs(float64(tile.X-position.X)) > 1 || math.Abs(float64(tile.Y-position.Y)) > 1 {
			continue
		} else {
			player_tiles = append(player_tiles, position)
			return true
		}
	}

	println("Tile is not adjacent to player collection")
	return false
}
