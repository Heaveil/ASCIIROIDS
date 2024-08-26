package src

import tl "github.com/JoelOtter/termloop"

type Spaceship struct {
	*tl.Entity
	Level   *tl.BaseLevel
	Score   int
	X       int
	Y       int
	Face    Direction
	Powered bool
	Master  bool
	Vim     bool
}

func bulletCoords(big bool, spaceship *Spaceship) (x, y int) {
	if big {
		switch spaceship.Face {
		case NORTH:
			x = spaceship.X + 1
			y = spaceship.Y - 2
		case NORTHEAST:
			x = spaceship.X + 1
			y = spaceship.Y - 2
		case EAST:
			x = spaceship.X + 5
			y = spaceship.Y
		case SOUTHEAST:
			x = spaceship.X + 3
			y = spaceship.Y + 2
		case SOUTH:
			x = spaceship.X + 1
			y = spaceship.Y + 2
		case SOUTHWEST:
			x = spaceship.X + 1
			y = spaceship.Y + 2
		case WEST:
			x = spaceship.X - 3
			y = spaceship.Y
		case NORTHWEST:
			x = spaceship.X + 1
			y = spaceship.Y - 2
		}
	} else {
		switch spaceship.Face {
		case EAST:
			x = spaceship.X + 5
			y = spaceship.Y + 1
		case WEST:
			x = spaceship.X
			y = spaceship.Y + 1
		default:
			x = spaceship.X + 2
			y = spaceship.Y + 1
		}
	}
	return
}

func (spaceship *Spaceship) Shoot() {
	if spaceship.Powered {
		x, y := bulletCoords(true, spaceship)
		bullet := NewBigBullet(x, y, spaceship.Face, spaceship)
		spaceship.Level.AddEntity(bullet)
	} else {
		x, y := bulletCoords(false, spaceship)
		bullet := NewBullet(x, y, spaceship.Face, false, spaceship)
		spaceship.Level.AddEntity(bullet)
	}
}

func (spaceship *Spaceship) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		spaceship.X, spaceship.Y = spaceship.Position()
		if spaceship.Master {
			if spaceship.Vim {
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
			} else {
				switch event.Ch {
				case 'w':
					spaceship.Face = NORTH
					spaceship.SetPosition(spaceship.X, spaceship.Y-1)
				case 'd':
					spaceship.Face = EAST
					spaceship.SetPosition(spaceship.X+1, spaceship.Y)
				case 's':
					spaceship.Face = SOUTH
					spaceship.SetPosition(spaceship.X, spaceship.Y+1)
				case 'a':
					spaceship.Face = WEST
					spaceship.SetPosition(spaceship.X-1, spaceship.Y)
				default:
					if event.Key == tl.KeySpace {
						spaceship.Shoot()
					}
				}
			}
		} else {
			switch event.Key {
			case tl.KeyArrowUp:
				spaceship.Face = NORTH
				spaceship.SetPosition(spaceship.X, spaceship.Y-1)
			case tl.KeyArrowRight:
				spaceship.Face = EAST
				spaceship.SetPosition(spaceship.X+1, spaceship.Y)
			case tl.KeyArrowDown:
				spaceship.Face = SOUTH
				spaceship.SetPosition(spaceship.X, spaceship.Y+1)
			case tl.KeyArrowLeft:
				spaceship.Face = WEST
				spaceship.SetPosition(spaceship.X-1, spaceship.Y)
			case tl.KeyEnter:
				spaceship.Shoot()
			}
		}
	}
}

func (spaceship *Spaceship) Render() {
	allRenders := []Render{SHIP_NORTH, SHIP_NORTHEAST, SHIP_EAST, SHIP_SOUTHEAST, SHIP_SOUTH, SHIP_SOUTHWEST, SHIP_WEST, SHIP_NORTHWEST}
	direction := allRenders[spaceship.Face]
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
