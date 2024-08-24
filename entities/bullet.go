package entities

import tl "github.com/JoelOtter/termloop"

type Bullet struct {
    *tl.Entity
    X int
    Y int
    face Direction
}
