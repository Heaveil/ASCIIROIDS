package entities

import tl "github.com/JoelOtter/termloop"

type Bullet struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Face      Direction
	Enemy	  bool
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

	if (bullet.Enemy){
		bullet.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: 'x'})
	} else {
		bullet.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: '+'})
	}

	bullet.Entity.Draw(screen)

}

func (asteroid *Asteroids) Split(Spaceship *Spaceship) {
	offsetX := 3
	offsetY := 3

	asteroid1 := NewSmallAsteroid(asteroid.X+offsetX, asteroid.Y+offsetY, asteroid.Face)
	asteroid2 := NewSmallAsteroid(asteroid.X-offsetX, asteroid.Y-offsetY, asteroid.Face)

	Spaceship.Level.AddEntity(&asteroid1)
	Spaceship.Level.AddEntity(&asteroid2)
}

func (bullet *Bullet) Collide(collision tl.Physical) {
	if Asteroids, ok := collision.(*Asteroids); ok {
		bullet.Spaceship.Level.RemoveEntity(Asteroids)
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Score += 5

		if Asteroids.Big {
			Asteroids.Split(bullet.Spaceship)
		}
	}
}
