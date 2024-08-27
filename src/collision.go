package src

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"os"
	"time"
)

func (bullet *BigBullet) Collide(collision tl.Physical) {
	if Asteroid, ok := collision.(*Asteroid); ok {
		bullet.Spaceship.Level.RemoveEntity(Asteroid)
		bullet.Spaceship.Level.RemoveEntity(bullet)
		if Asteroid.Big {
			Asteroid.Split(bullet.Spaceship)
			bullet.Spaceship.Score += 3
		} else {
			bullet.Spaceship.Score += 1
		}
	}

	if Bullet, ok := collision.(*Bullet); ok {
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Level.RemoveEntity(Bullet)
	}

	if Turret, ok := collision.(*Turret); ok {
		bullet.Spaceship.Level.AddEntity(NewMissile(Turret.X, Turret.Y, bullet.Spaceship))
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Level.RemoveEntity(Turret)
		bullet.Spaceship.Score += 2
	}

	if Missile, ok := collision.(*Missile); ok {
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Level.RemoveEntity(Missile)
		bullet.Spaceship.Score += 3
	}
}

func (bullet *Bullet) Collide(collision tl.Physical) {
	if Asteroid, ok := collision.(*Asteroid); ok {
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Level.RemoveEntity(Asteroid)
		if Asteroid.Big {
			Asteroid.Split(bullet.Spaceship)
			bullet.Spaceship.Score += 3
		} else {
			bullet.Spaceship.Score += 1
		}
	}

	if Bullet, ok := collision.(*Bullet); ok {
		if Bullet.Enemy != bullet.Enemy {
			bullet.Spaceship.Level.RemoveEntity(bullet)
			bullet.Spaceship.Level.RemoveEntity(Bullet)
		}
	}

	if Turret, ok := collision.(*Turret); ok {
		if !bullet.Enemy {
			bullet.Spaceship.Level.AddEntity(NewMissile(Turret.X, Turret.Y, bullet.Spaceship))
			bullet.Spaceship.Level.RemoveEntity(bullet)
			bullet.Spaceship.Level.RemoveEntity(Turret)
			bullet.Spaceship.Score += 2
		}
	}

	if Missile, ok := collision.(*Missile); ok {
		bullet.Spaceship.Level.RemoveEntity(bullet)
		bullet.Spaceship.Level.RemoveEntity(Missile)
		bullet.Spaceship.Score += 3
	}
}

func (spaceship *Spaceship) Collide(collision tl.Physical) {

	if Powerup, ok := collision.(*Powerup); ok {
		spaceship.Powered = true
		spaceship.Level.RemoveEntity(Powerup)

		// Timer for powerup
		ticker := time.NewTicker(5 * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					spaceship.Powered = false
				}
			}
		}()
	}

	if Bullet, ok := collision.(*Bullet); ok {
		if Bullet.Enemy {
			fmt.Printf("\nSCORE | %d\n", spaceship.Score)
			os.Exit(0)
		}
	}

	if _, ok := collision.(*Asteroid); ok {
		fmt.Printf("\nSCORE | %d\n", spaceship.Score)
		os.Exit(0)
	}

	if _, ok := collision.(*Turret); ok {
		fmt.Printf("\nSCORE | %d\n", spaceship.Score)
		os.Exit(0)
	}

	if _, ok := collision.(*Missile); ok {
		fmt.Printf("\nSCORE | %d\n", spaceship.Score)
		os.Exit(0)
	}

}
