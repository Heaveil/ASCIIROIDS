package entities

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Missile_Render [][]rune

var MissileRender = Missile_Render {
	{' ', '^', ' '},
	{'<', '+', '>'},
	{' ', 'V', ' '}}

type Missile struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
}

func (missile *Missile) Draw(screen *tl.Screen) {
	missile.X, missile.Y = missile.Position()
	SpaceshipX, SpaceshipY := missile.Spaceship.X, missile.Spaceship.Y

	switch {
	case missile.X > SpaceshipX && missile.Y > SpaceshipY:
		missile.SetPosition(missile.X-1, missile.Y-1)
	case missile.X > SpaceshipX && missile.Y < SpaceshipY:
		missile.SetPosition(missile.X-1, missile.Y+1)
	case missile.X > SpaceshipX && missile.Y == SpaceshipY:
		missile.SetPosition(missile.X-1, missile.Y)
	case missile.X < SpaceshipX && missile.Y > SpaceshipY:
		missile.SetPosition(missile.X+1, missile.Y-1)
	case missile.X < SpaceshipX && missile.Y < SpaceshipY:
		missile.SetPosition(missile.X+1, missile.Y+1)
	case missile.X < SpaceshipX && missile.Y == SpaceshipY:
		missile.SetPosition(missile.X+1, missile.Y)
	case missile.X == SpaceshipX && missile.Y > SpaceshipY:
		missile.SetPosition(missile.X, missile.Y-1)
	case missile.X == SpaceshipX && missile.Y < SpaceshipY:
		missile.SetPosition(missile.X, missile.Y+1)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			missile.SetCell(i, j, &tl.Cell{Fg: tl.ColorRed, Ch: MissileRender[j][i]})
		}
	}
	missile.Entity.Draw(screen)
}

func SpawnMissile(spaceship *Spaceship) {
	ticker := time.NewTicker(275 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				spawn_point := rand.Intn(4)
				x,y := 0,0

				switch spawn_point {
				case 0:
					x = spaceship.X + 30
					y = spaceship.Y + 30
				case 1:
					x = spaceship.X - 30
					y = spaceship.Y - 30
				case 2:
					x = spaceship.X - 30
					y = spaceship.Y + 30
				case 3:
					x = spaceship.X + 30
					y = spaceship.Y - 30
				}

				missile := Missile{
					Entity: tl.NewEntity(x, y, 3, 3),
					Spaceship: spaceship,
					X : x,
					Y : y,
				}
				spaceship.Level.AddEntity(&missile)
			}
		}
	}()
}

func (missile *Missile) Collide(collision tl.Physical) {
	if Bullet, ok := collision.(*Bullet); ok {
		missile.Spaceship.Level.RemoveEntity(Bullet)
		missile.Spaceship.Level.RemoveEntity(missile)
		missile.Spaceship.Score += 5
	}
}
