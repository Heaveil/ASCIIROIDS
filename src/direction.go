package src

type Direction int

const (
	NORTH Direction = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

var directionVectors = map[Direction]struct{ dx, dy int }{
	NORTH:     {0, -1},
	NORTHEAST: {1, -1},
	EAST:      {1, 0},
	SOUTHEAST: {1, 1},
	SOUTH:     {0, 1},
	SOUTHWEST: {-1, 1},
	WEST:      {-1, 0},
	NORTHWEST: {-1, -1},
}

func determineDirection(dx, dy int) (Direction, bool) {
	for direction, vector := range directionVectors {
		if vector.dx == dx && vector.dy == dy {
			return direction, true
		}
	}
	return 0, false
}
