package src

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Asteroid struct {
	*tl.Entity
	X    int
	Y    int
	Face Direction
	Big  bool
}

func NewAsteroid(x, y int, face Direction, big bool) *Asteroid {
	width, height := 5, 3
	if big {
		width, height = 9, 4
	}
	return &Asteroid{
		Entity: tl.NewEntity(x, y, width, height),
		X:      x,
		Y:      y,
		Face:   face,
		Big:    big,
	}
}

func generateAsteroid(spaceship *Spaceship) (x, y int, direction Direction, size bool) {

	whichSize := rand.Intn(2)
	switch whichSize {
	case 0:
		size = true
	case 1:
		size = false
	}

	whichQuardrant := rand.Intn(4)
	distance := 30
	whichPoint := rand.Intn(60) - 30
	whichDirection := rand.Intn(3)

	switch whichQuardrant {
	case 0:
		x = spaceship.X - distance
		y = spaceship.Y + whichPoint
		switch whichDirection {
		case 0:
			direction = NORTHEAST
		case 1:
			direction = EAST
		case 2:
			direction = SOUTHEAST
		}
	case 1:
		x = spaceship.X + distance
		y = spaceship.Y + whichPoint
		switch whichDirection {
		case 0:
			direction = NORTHWEST
		case 1:
			direction = WEST
		case 2:
			direction = SOUTHWEST
		}
	case 2:
		x = spaceship.X + whichPoint
		y = spaceship.Y + distance
		switch whichDirection {
		case 0:
			direction = SOUTHWEST
		case 1:
			direction = SOUTH
		case 2:
			direction = SOUTHEAST
		}
	case 3:
		x = spaceship.X + whichPoint
		y = spaceship.Y - distance
		switch whichDirection {
		case 0:
			direction = NORTHWEST
		case 1:
			direction = NORTH
		case 2:
			direction = NORTHEAST
		}
	}
	return
}

func SpawnAsteroid(spaceship *Spaceship) {
	ticker := time.NewTicker(400 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				asteroid := NewAsteroid(generateAsteroid(spaceship))
				spaceship.Level.AddEntity(asteroid)
			}
		}
	}()
}

func (asteroid *Asteroid) Split(spaceship *Spaceship) {
	offsetX := 3
	offsetY := 3
	asteroid1 := NewAsteroid(asteroid.X+offsetX, asteroid.Y+offsetY, asteroid.Face, false)
	asteroid2 := NewAsteroid(asteroid.X-offsetX, asteroid.Y-offsetY, asteroid.Face, false)
	spaceship.Level.AddEntity(asteroid1)
	spaceship.Level.AddEntity(asteroid2)
}

func (asteroid *Asteroid) Render() {
	width, height, render := 5, 3, ASTEROID_SMALL
	if asteroid.Big {
		width, height, render = 9, 4, ASTEROID_BIG
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			asteroid.SetCell(i, j, &tl.Cell{Fg: tl.ColorWhite, Ch: render[j][i]})
		}
	}
}

func (asteroid *Asteroid) Move() {
	asteroid.X, asteroid.Y = asteroid.Position()
	if vector, ok := directionVectors[asteroid.Face]; ok {
		asteroid.SetPosition(asteroid.X+vector.dx, asteroid.Y+vector.dy)
	}
}

func (asteroid *Asteroid) Draw(screen *tl.Screen) {
	asteroid.Move()
	asteroid.Render()
	asteroid.Entity.Draw(screen)
}
