package entities

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Asteroid_Render [][]rune

var SMALL_ASTEROID = Asteroid_Render{
	{' ', '_', '_', '_', ' '},
	{'/', ' ', ' ', ' ', '\\'},
	{'\\', '_', '_', '_', '/'},
}

var BIG_ASTEROID = Asteroid_Render{
	{' ', ' ', '_', '_', '_', '_', '_', ' ', ' '},
	{' ', '/', ' ', ' ', ' ', ' ', ' ', '\\', ' '},
	{'[', ' ', ' ', ' ', 'x', ' ', ' ', ' ', ']'},
	{' ', '\\', '_', '_', '_', '_', '_', '/', ' '},
}

type Asteroids struct {
	*tl.Entity
	X    int
	Y    int
	Big  bool
	Face Direction
}

func NewSmallAsteroid(x, y int, face Direction) (asteroid Asteroids) {
	asteroid = Asteroids{
		Entity: tl.NewEntity(x, y, 5, 3),
		X:      x,
		Y:      y,
		Big:    false,
		Face:   face,
	}
	return
}

func NewBigAsteroid(x, y int, face Direction) (asteroid Asteroids) {
	asteroid = Asteroids{
		Entity: tl.NewEntity(x, y, 9, 4),
		X:      x,
		Y:      y,
		Big:    true,
		Face:   face,
	}
	return
}

func GetSide(spaceship *Spaceship) (x, y int, dir Direction) {
	randomInt := rand.Intn(4)
	randomIntDir := rand.Intn(3)
	randomRange := rand.Intn(60) - 30
	switch randomInt {
	case 0:
		x = spaceship.X - 30
		y = spaceship.Y + randomRange
		switch randomIntDir {
		case 0:
			dir = NORTHEAST
		case 1:
			dir = EAST
		case 2:
			dir = SOUTHEAST
		}
	case 1:
		x = spaceship.X + 30
		y = spaceship.Y + randomRange
		switch randomIntDir {
		case 0:
			dir = NORTHWEST
		case 1:
			dir = WEST
		case 2:
			dir = SOUTHWEST
		}
	case 2:
		x = spaceship.X + randomRange
		y = spaceship.Y + 30
		switch randomIntDir {
		case 0:
			dir = SOUTHWEST
		case 1:
			dir = SOUTH
		case 2:
			dir = SOUTHEAST
		}
	case 3:
		x = spaceship.X + randomRange
		y = spaceship.Y - 30
		switch randomIntDir {
		case 0:
			dir = NORTHWEST
		case 1:
			dir = NORTH
		case 2:
			dir = NORTHEAST
		}
	}
	return
}

func SpawnAsteroids(spaceship *Spaceship) {

	ticker := time.NewTicker(275 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				x, y, dir := GetSide(spaceship)

				randomInt := rand.Intn(2)
				var asteroid Asteroids

				switch randomInt {

				case 0:
					asteroid = NewSmallAsteroid(x, y, dir)

				case 1:
					asteroid = NewBigAsteroid(x, y, dir)
				}

				spaceship.Level.AddEntity(&asteroid)
			}
		}
	}()
}

func (asteroid *Asteroids) Render() {
	if asteroid.Big {
		for i := 0; i < 9; i++ {
			for j := 0; j < 4; j++ {
				asteroid.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: BIG_ASTEROID[j][i]})
			}
		}
	} else {
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				asteroid.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: SMALL_ASTEROID[j][i]})
			}
		}
	}
}

func (asteroid *Asteroids) Draw(screen *tl.Screen) {
	asteroid.X, asteroid.Y = asteroid.Position()
	switch asteroid.Face {
	case NORTH:
		asteroid.SetPosition(asteroid.X, asteroid.Y-1)
	case NORTHEAST:
		asteroid.SetPosition(asteroid.X+1, asteroid.Y-1)
	case EAST:
		asteroid.SetPosition(asteroid.X+1, asteroid.Y)
	case SOUTHEAST:
		asteroid.SetPosition(asteroid.X+1, asteroid.Y+1)
	case SOUTH:
		asteroid.SetPosition(asteroid.X, asteroid.Y+1)
	case SOUTHWEST:
		asteroid.SetPosition(asteroid.X-1, asteroid.Y+1)
	case WEST:
		asteroid.SetPosition(asteroid.X-1, asteroid.Y)
	case NORTHWEST:
		asteroid.SetPosition(asteroid.X-1, asteroid.Y-1)
	}
	asteroid.Render()
	asteroid.Entity.Draw(screen)
}
