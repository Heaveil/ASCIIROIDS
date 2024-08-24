package entities

import tl "github.com/JoelOtter/termloop"

type Bullet struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Face      Direction
}

func (bullet *Bullet) Draw(screen *tl.Screen) {
	bullet.X, bullet.Y = bullet.Position()
	switch bullet.Face {
	case NORTH:
		bullet.SetPosition(bullet.X, bullet.Y-1)
	case NORTHEAST:
		bullet.SetPosition(bullet.X+1, bullet.Y-1)
	case EAST:
		bullet.SetPosition(bullet.X+1, bullet.Y)
	case SOUTHEAST:
		bullet.SetPosition(bullet.X+1, bullet.Y+1)
	case SOUTH:
		bullet.SetPosition(bullet.X, bullet.Y+1)
	case SOUTHWEST:
		bullet.SetPosition(bullet.X-1, bullet.Y+1)
	case WEST:
		bullet.SetPosition(bullet.X-1, bullet.Y)
	case NORTHWEST:
		bullet.SetPosition(bullet.X-1, bullet.Y-1)
	}
	bullet.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: '+'})
	bullet.Entity.Draw(screen)
}

func (bullet *Bullet) Collide(collision tl.Physical) {
	if Asteroids, ok := collision.(*Asteroids); ok {
		bullet.Spaceship.Level.RemoveEntity(Asteroids)
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Score += 5

	}
}
