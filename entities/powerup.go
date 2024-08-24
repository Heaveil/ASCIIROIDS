package entities

import (
	tl "github.com/JoelOtter/termloop"
	"time"
)

type Powerup_Render [][]rune

var POWERUP = Powerup_Render{
	{' ', '$', '$'},
	{' ', '$', ' '},
	{'$', '$', ' '},
}

type Powerup struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
}

func NewPowerup(x, y int, spaceship *Spaceship) *Powerup {
	return &Powerup{
		Entity:    tl.NewEntity(x, y, 3, 3),
		Spaceship: spaceship,
		X:         x,
		Y:         y,
	}
}

func SpawnPowerup(spaceship *Spaceship) {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				x := 5
				y := 5

				powerup := NewPowerup(x, y, spaceship)
				spaceship.Level.AddEntity(powerup)
			}
		}
	}()
}

func (powerup *Powerup) Draw(screen *tl.Screen) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			powerup.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: POWERUP[j][i]})
		}
	}
	powerup.Entity.Draw(screen)
}
