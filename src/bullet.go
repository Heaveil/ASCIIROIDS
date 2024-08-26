package src

import tl "github.com/JoelOtter/termloop"

type Bullet struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Face      Direction
	Enemy     bool
}

func NewBullet(x, y int, face Direction, enemy bool, spaceship *Spaceship) *Bullet {
	return &Bullet{
		Entity:    tl.NewEntity(x, y, 1, 1),
		Spaceship: spaceship,
		X:         x,
		Y:         y,
		Face:      face,
		Enemy:     enemy,
	}
}

func (bullet *Bullet) Move() {
	bullet.X, bullet.Y = bullet.Position()
	if vector, ok := directionVectors[bullet.Face]; ok {
		bullet.SetPosition(bullet.X+vector.dx, bullet.Y+vector.dy)
	}
}

func (bullet *Bullet) Draw(screen *tl.Screen) {
	bullet.Move()
	if bullet.Enemy {
		bullet.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: 'x'})
	} else {
		bullet.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: '+'})
	}
	bullet.Entity.Draw(screen)
}
