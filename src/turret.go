package src

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Turret struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Cooldown  bool
}

func NewTurret(x, y int, spaceship *Spaceship) *Turret {
	return &Turret{
		Entity:    tl.NewEntity(x, y, 5, 3),
		Spaceship: spaceship,
		X:         x,
		Y:         y,
		Cooldown:  false,
	}
}

func SpawnTurret(spaceship *Spaceship) {
	ticker := time.NewTicker(7 * time.Second)
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
				turret := NewTurret(x, y, spaceship)
				spaceship.Level.AddEntity(turret)
			}
		}
	}()
}

func (turret *Turret) Shoot(direction Direction) {
	ticker := time.NewTicker(1 * time.Second)
	if !turret.Cooldown {
		bullet := NewBullet(turret.X+2, turret.Y+1, direction, true, turret.Spaceship)
		turret.Spaceship.Level.AddEntity(bullet)
		turret.Cooldown = true
		go func() {
			for {
				select {
				case <-ticker.C:
					turret.Cooldown = false
				}
			}
		}()
	}
}

func (turret *Turret) Move() {
	dx, dy := 0, 0
	turret.X, turret.Y = turret.Position()
	spaceshipX, spaceshipY := turret.Spaceship.Position()

	if turret.X > spaceshipX {
		dx = -1
	} else if turret.X < spaceshipX {
		dx = 1
	}

	if turret.Y > spaceshipY {
		dy = -1
	} else if turret.Y < spaceshipY {
		dy = 1
	}

	if direction, ok := determineDirection(dx, dy); ok {
		turret.Shoot(direction)
	}
}

func (turret *Turret) Draw(screen *tl.Screen) {
	turret.Move()
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			turret.SetCell(i, j, &tl.Cell{Fg: tl.ColorMagenta, Ch: TURRET[j][i]})
		}
	}
	turret.Entity.Draw(screen)
}
