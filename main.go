package main

import (
	"asciiroids/entities"
	"flag"
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
		Master: true,
	}
	level.AddEntity(&spaceship)

	multi := flag.Bool("multi", false, "Enable multipayer")
	bigA := flag.Bool("bigA", false, "Enable big asteroid")
	smallA := flag.Bool("smallA", false, "Enable small asteroid")
	missile := flag.Bool("missile", false, "Enable missle")
	turret := flag.Bool("turret", false, "Enable turret")
	powerup := flag.Bool("powerup", false, "Enable powerup")

	flag.Parse()

	if *multi {
		spaceship2 := entities.Spaceship{
			Entity: tl.NewEntity(7, 0, 5, 3),
			Level:  level,
			Face:   entities.NORTH,
			Master: false,
		}
		level.AddEntity(&spaceship2)

		entities.SpawnAsteroids(&spaceship)
		entities.SpawnMissile(&spaceship)
		entities.SpawnPowerup(&spaceship)
		entities.SpawnTurret(&spaceship)
	} else if *missile {
		missile := entities.Missile{
			Entity:    tl.NewEntity(0, -20, 3, 3),
			Spaceship: &spaceship,
			X:         10,
			Y:         10,
		}
		spaceship.Level.AddEntity(&missile)
	} else if *turret {
		turret := entities.Turret{
			Entity:    tl.NewEntity(30, 0, 5, 3),
			Spaceship: &spaceship,
			X:         30,
			Y:         0,
			Cooldown:  false,
		}
		spaceship.Level.AddEntity(&turret)
	} else if *bigA {
		aster := entities.Asteroids{
			Entity: tl.NewEntity(0, -20, 9, 4),
			X:      0,
			Y:      -20,
			Big:    true,
			Face:   entities.SOUTH,
		}
		spaceship.Level.AddEntity(&aster)
	} else if *smallA {
		aster := entities.Asteroids{
			Entity: tl.NewEntity(0, -20, 5, 3),
			X:      0,
			Y:      -20,
			Big:    false,
			Face:   entities.SOUTH,
		}
		spaceship.Level.AddEntity(&aster)
	} else if *powerup {
		powerup := entities.NewPowerup(10, 10, &spaceship)
		spaceship.Level.AddEntity(powerup)
	} else {
		entities.SpawnAsteroids(&spaceship)
		entities.SpawnMissile(&spaceship)
		entities.SpawnPowerup(&spaceship)
		entities.SpawnTurret(&spaceship)

	}

	game.Screen().SetLevel(level)
	game.Start()
}
