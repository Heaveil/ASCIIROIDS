package main

import (
	"asciiroids/entities"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(15)
	level := tl.NewBaseLevel(tl.Cell{})
	spaceship := entities.Spaceship{
		Entity: tl.NewEntity(0, 0, 5, 3),
		Level:  level,
		Face:   entities.NORTH,
	}

	level.AddEntity(&spaceship)
	entities.SpawnAsteroids(&spaceship)
	entities.SpawnMissile(&spaceship)
	entities.SpawnPowerup(&spaceship)
	entities.SpawnTurret(&spaceship)
	game.Screen().SetLevel(level)
	game.Start()
}
