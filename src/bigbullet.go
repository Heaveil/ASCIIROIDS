package src

import tl "github.com/JoelOtter/termloop"

type BigBullet struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
	Face      Direction
}

func NewBigBullet(x, y int, face Direction, spaceship *Spaceship) *BigBullet {
	return &BigBullet{
		Entity:    tl.NewEntity(x, y, 3, 3),
		Spaceship: spaceship,
		X:         x,
		Y:         y,
		Face:      face,
	}
}

func (bullet *BigBullet) Move() {
	bullet.X, bullet.Y = bullet.Position()
	if vector, ok := directionVectors[bullet.Face]; ok {
		bullet.SetPosition(bullet.X+vector.dx, bullet.Y+vector.dy)
	}
}

func (bullet *BigBullet) Draw(screen *tl.Screen) {
	bullet.Move()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bullet.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: BIGBULLET[j][i]})
		}
	}
	bullet.Entity.Draw(screen)
}
