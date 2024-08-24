package entities

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"os"
	"time"
)

type Ship_Render [][]rune

var SHIP_NORTH = Ship_Render{
	{' ', '|', ' ', '|', ' '},
	{'|', '-', '-', '-', '|'},
	{' ', '~', '~', '~', ' '}}

var SHIP_NORTHEAST = Ship_Render{
	{' ', ' ', '/', ' ', '/'},
	{'/', '-', '-', '-', '/'},
	{'~', '~', '~', '~', ' '}}

var SHIP_EAST = Ship_Render{
	{'-', '-', '-', ' ', ' '},
	{' ', '|', ' ', '=', '='},
	{'-', '-', '-', ' ', ' '}}

var SHIP_SOUTHEAST = Ship_Render{
	{'~', '~', '~', '~', ' '},
	{'\\', '-', '-', '-', '\\'},
	{' ', ' ', '\\', ' ', '\\'}}

var SHIP_SOUTH = Ship_Render{
	{' ', '~', '~', '~', ' '},
	{'|', '-', '-', '-', '|'},
	{' ', '|', ' ', '|', ' '}}

var SHIP_SOUTHWEST = Ship_Render{
	{' ', '~', '~', '~', '~'},
	{'/', '-', '-', '-', '/'},
	{'/', ' ', '/', ' ', ' '}}

var SHIP_WEST = Ship_Render{
	{' ', ' ', '-', '-', '-'},
	{'=', '=', ' ', '|', ' '},
	{' ', ' ', '-', '-', '-'}}

var SHIP_NORTHWEST = Ship_Render{
	{'\\', ' ', '\\', ' ', ' '},
	{'\\', '-', '-', '-', '\\'},
	{' ', '~', '~', '~', '~'}}

type Spaceship struct {
	*tl.Entity
	Level   *tl.BaseLevel
	X       int
	Y       int
	Face    Direction
	Score   int
	Powered bool
}

func (spaceship *Spaceship) Render() {
	direction := Ship_Render{}
	switch spaceship.Face {
	case NORTH:
		direction = SHIP_NORTH
	case NORTHEAST:
		direction = SHIP_NORTHEAST
	case EAST:
		direction = SHIP_EAST
	case SOUTHEAST:
		direction = SHIP_SOUTHEAST
	case SOUTH:
		direction = SHIP_SOUTH
	case SOUTHWEST:
		direction = SHIP_SOUTHWEST
	case WEST:
		direction = SHIP_WEST
	case NORTHWEST:
		direction = SHIP_NORTHWEST
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			spaceship.SetCell(i, j, &tl.Cell{Fg: tl.ColorBlue, Ch: direction[j][i]})
		}
	}
}

func (spaceship *Spaceship) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := spaceship.Position()
	spaceship.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	spaceship.Render()
	spaceship.Entity.Draw(screen)
}

func get_coords(Spaceship Spaceship) (x, y int) {
	switch Spaceship.Face {
	case EAST:
		x = Spaceship.X + 5
		y = Spaceship.Y + 1
	case WEST:
		x = Spaceship.X
		y = Spaceship.Y + 1
	default:
		x = Spaceship.X + 2
		y = Spaceship.Y + 1
	}
	return
}

func get_coords_big(Spaceship Spaceship) (x, y int) {
	switch Spaceship.Face {
	case NORTH:
		x = Spaceship.X + 1
		y = Spaceship.Y - 2
	case NORTHEAST:
		x = Spaceship.X + 1
		y = Spaceship.Y - 2
	case EAST:
		x = Spaceship.X + 5
		y = Spaceship.Y
	case SOUTHEAST:
		x = Spaceship.X + 3
		y = Spaceship.Y + 2
	case SOUTH:
		x = Spaceship.X + 1
		y = Spaceship.Y + 2
	case SOUTHWEST:
		x = Spaceship.X + 1
		y = Spaceship.Y + 2
	case WEST:
		x = Spaceship.X - 3
		y = Spaceship.Y
	case NORTHWEST:
		x = Spaceship.X + 1
		y = Spaceship.Y - 2
	}
	return
}

func (Spaceship *Spaceship) Shoot() {
	if Spaceship.Powered {
		x, y := get_coords_big(*Spaceship)
		bigbullet := Bigbullet{
			Entity:    tl.NewEntity(x, y, 3, 3),
			Face:      Spaceship.Face,
			Spaceship: Spaceship,
		}
		Spaceship.Level.AddEntity(&bigbullet)
	} else {
		x, y := get_coords(*Spaceship)
		bullet := Bullet{
			Entity:    tl.NewEntity(x, y, 1, 1),
			Face:      Spaceship.Face,
			Spaceship: Spaceship,
			Enemy:     false,
		}
		Spaceship.Level.AddEntity(&bullet)
	}
}

func (spaceship *Spaceship) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		spaceship.X, spaceship.Y = spaceship.Position()
		switch event.Ch {
		case 'k':
			spaceship.Face = NORTH
			spaceship.SetPosition(spaceship.X, spaceship.Y-1)
		case 'g':
			spaceship.Face = NORTHEAST
			spaceship.SetPosition(spaceship.X+1, spaceship.Y-1)
		case 'l':
			spaceship.Face = EAST
			spaceship.SetPosition(spaceship.X+1, spaceship.Y)
		case 'f':
			spaceship.Face = SOUTHEAST
			spaceship.SetPosition(spaceship.X+1, spaceship.Y+1)
		case 'j':
			spaceship.Face = SOUTH
			spaceship.SetPosition(spaceship.X, spaceship.Y+1)
		case 'd':
			spaceship.Face = SOUTHWEST
			spaceship.SetPosition(spaceship.X-1, spaceship.Y+1)
		case 'h':
			spaceship.Face = WEST
			spaceship.SetPosition(spaceship.X-1, spaceship.Y)
		case 's':
			spaceship.Face = NORTHWEST
			spaceship.SetPosition(spaceship.X-1, spaceship.Y-1)
		default:
			if event.Key == tl.KeySpace {
				spaceship.Shoot()
			}
		}
	}
}

func (spaceship *Spaceship) Collide(collision tl.Physical) {
	if _, ok := collision.(*Asteroids); ok {
		fmt.Println("\n\n\n\n")
		fmt.Println(spaceship.Score)
		os.Exit(0)
	}

	if _, ok := collision.(*Missile); ok {
		fmt.Println("\n\n\n\n")
		fmt.Println(spaceship.Score)
		os.Exit(0)
	}

	if _, ok := collision.(*Turret); ok {
		fmt.Println("\n\n\n\n")
		fmt.Println(spaceship.Score)
		os.Exit(0)
	}

	if powerup, ok := collision.(*Powerup); ok {
		spaceship.Powered = true
		spaceship.Level.RemoveEntity(powerup)

		// sets 5 second timer for powerup
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
}
