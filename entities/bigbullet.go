package entities

import tl "github.com/JoelOtter/termloop"

type BigBullet [][]rune

var BIGBULLET = BigBullet{
	{'+', '+', '+'},
	{'+', '+', '+'},
	{'+', '+', '+'},
}

type Bigbullet struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Face      Direction
}

func (bullet *Bigbullet) Draw(screen *tl.Screen) {
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

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bullet.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: BIGBULLET[j][i]})
		}
	}

	bullet.Entity.Draw(screen)

}

func (bullet *Bigbullet) Collide(collision tl.Physical) {
	if Asteroids, ok := collision.(*Asteroids); ok {
		bullet.Spaceship.Level.RemoveEntity(Asteroids)
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Score += 5

		if Asteroids.Big {
			Asteroids.Split(bullet.Spaceship)
		}
	}
}
