package src

import tl "github.com/JoelOtter/termloop"

type Missile struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
}

func NewMissile(x, y int, spaceship *Spaceship) *Missile {
	return &Missile{
		Entity:    tl.NewEntity(x, y, 3, 3),
		Spaceship: spaceship,
		X:         x,
		Y:         y,
	}
}

func (missile *Missile) Move() {
	dx, dy := 0, 0
	missile.X, missile.Y = missile.Position()
	spaceshipX, spaceshipY := missile.Spaceship.Position()

	if missile.X > missile.Y {
		dx = -1
	} else if missile.X < spaceshipX {
		dx = 1
	}

	if missile.Y > spaceshipY {
		dy = -1
	} else if missile.Y < spaceshipY {
		dy = 1
	}

	missile.SetPosition(missile.X+dx, missile.Y+dy)
}

func (missile *Missile) Draw(screen *tl.Screen) {
	missile.Move()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			missile.SetCell(i, j, &tl.Cell{Fg: tl.ColorRed, Ch: MISSILE[j][i]})
		}
	}
	missile.Entity.Draw(screen)
}
