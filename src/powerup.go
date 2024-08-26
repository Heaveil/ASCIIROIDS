package src

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

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
				distance := 20
				spawnPoint := rand.Intn(8)
				x, y := spaceship.X, spaceship.Y
				if vector, ok := directionVectors[Direction(spawnPoint)]; ok {
					x += vector.dx * distance
					y += vector.dy * distance
				}
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
