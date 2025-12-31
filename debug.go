package main

import "fmt"

func debug_player_tiles() {
	if len(player_tiles) == 0 {
		println("Player doesn't have any tile")
		return
	}

	for i := 0; i < len(player_tiles); i++ {
		print(fmt.Sprintf("Tile %v is located at {X: %v, Y:%v}, Owned by %v and has %v health\n", i, player_tiles[i].Position.X, player_tiles[i].Position.Y, player_tiles[i].Owner, player_tiles[i].Health))
	}
}
