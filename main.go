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

type Tile struct {
	Position gomesmath.Vector2
	Owner    int
	Health   int
}

var player_tiles []Tile

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

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_SPACE, func(data any) {
		debug_player_tiles()
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
		player_tiles = make([]Tile, 1)
		player_tiles[0] = Tile{
			Position: position,
			Owner:    1,
			Health:   1,
		}

		return true
	}

	message := ""

	for _, tile := range player_tiles {
		if tile.Position == position {
			message = "it's already registered"
			continue
		}

		if math.Abs(float64(tile.Position.X-position.X)) > 1 || math.Abs(float64(tile.Position.Y-position.Y)) > 1 {
			message = "is not adjacent to current collection"
			continue
		} else {
			player_tiles = append(player_tiles, Tile{
				Position: position,
				Owner:    1,
				Health:   1,
			})
			return true
		}
	}

	println("Tile cannot be added to player collection because", message)
	return false
}
