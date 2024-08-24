package entities

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Turret_Render [][]rune

var TurretRender = Turret_Render{
	{' ', ' ', '^', ' ', ' '},
	{'<', '[', '0', ']', '>'},
	{' ', ' ', 'V', ' ', ' '}}

type Turret struct {
	*tl.Entity
	Spaceship *Spaceship
	X         int
	Y         int
}

func (turret *Turret) Shoot(direction Direction){
	bullet := Bullet{
		Entity:    tl.NewEntity(turret.X + 2, turret.Y + 1, 1, 1),
		Face:      direction,
		Spaceship: turret.Spaceship,
		Enemy:     true,
	}
	turret.Spaceship.Level.AddEntity(&bullet)
}

func (turret *Turret) Draw(screen *tl.Screen) {
	turret.X, turret.Y = turret.Position()
	SpaceshipX, SpaceshipY := turret.Spaceship.X, turret.Spaceship.Y

	switch {
	case turret.X > SpaceshipX && turret.Y > SpaceshipY:
		turret.Shoot(NORTHWEST)
	case turret.X > SpaceshipX && turret.Y < SpaceshipY:
		turret.Shoot(SOUTHWEST)
	case turret.X > SpaceshipX && turret.Y == SpaceshipY:
		turret.Shoot(WEST)
	case turret.X < SpaceshipX && turret.Y > SpaceshipY:
		turret.Shoot(NORTHEAST)
	case turret.X < SpaceshipX && turret.Y < SpaceshipY:
		turret.Shoot(SOUTHEAST)
	case turret.X < SpaceshipX && turret.Y == SpaceshipY:
		turret.Shoot(EAST)
	case turret.X == SpaceshipX && turret.Y > SpaceshipY:
		turret.Shoot(NORTH)
	case turret.X == SpaceshipX && turret.Y < SpaceshipY:
		turret.Shoot(SOUTH)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			turret.SetCell(i, j, &tl.Cell{Fg: tl.ColorRed, Ch: TurretRender[j][i]})
		}
	}
	turret.Entity.Draw(screen)
}

func SpawnTurret(spaceship *Spaceship) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				spawn_point := rand.Intn(4)
				x, y := 0, 0

				switch spawn_point {
				case 0:
					x = spaceship.X + 20
					y = spaceship.Y + 20
				case 1:
					x = spaceship.X - 20
					y = spaceship.Y - 20
				case 2:
					x = spaceship.X - 20
					y = spaceship.Y + 20
				case 3:
					x = spaceship.X + 20
					y = spaceship.Y - 20
				}

				turret := Turret{
					Entity:    tl.NewEntity(x, y, 5, 3),
					Spaceship: spaceship,
					X:         x,
					Y:         y,
				}
				spaceship.Level.AddEntity(&turret)
			}
		}
	}()
}

func (turret *Turret) Collide(collision tl.Physical) {
	if Bullet, ok := collision.(*Bullet); ok {
		if (!Bullet.Enemy){
			turret.Spaceship.Level.RemoveEntity(Bullet)
			turret.Spaceship.Level.RemoveEntity(turret)
			turret.Spaceship.Score += 5
		}
	}
	if Bigbullet, ok := collision.(*Bigbullet); ok {
		turret.Spaceship.Level.RemoveEntity(Bigbullet)
		turret.Spaceship.Level.RemoveEntity(turret)
		turret.Spaceship.Score += 5
	}
}
