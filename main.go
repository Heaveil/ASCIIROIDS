package main

import (
	"asciiroids/src"
	"flag"
	tl "github.com/JoelOtter/termloop"
)

func main() {

	// Get which flags are called
	multi := flag.Bool("multi", false, "Enable multipayer")
	vim := flag.Bool("vim", false, "Enable vim keys")
	flag.Parse()

	// Create the game
	level := tl.NewBaseLevel(tl.Cell{})
	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Screen().SetFps(15)

	// Create Spaceship
	spaceship := &src.Spaceship{
		Entity: tl.NewEntity(0, 0, 5, 3),
		Level:  level,
		Face:   src.NORTH,
		Master: true,
		Vim:    *vim,
	}
	level.AddEntity(spaceship)

	// Create second spaceship
	if *multi {
		spaceship2 := &src.Spaceship{
			Entity: tl.NewEntity(7, 0, 5, 3),
			Level:  level,
			Face:   src.NORTH,
			Master: false,
		}
		level.AddEntity(spaceship2)
	}

	// Spawn entities
	src.SpawnAsteroid(spaceship)
	src.SpawnTurret(spaceship)
	src.SpawnPowerup(spaceship)

	// Start game
	game.Start()
}
