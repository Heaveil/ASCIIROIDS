package main

import (
    tl "github.com/JoelOtter/termloop"
    "asciiroids/entities"
)

func main() {
    game := tl.NewGame()
    game.Screen().SetFps(30)
    level := tl.NewBaseLevel(tl.Cell{})
    spaceship := entities.Spaceship{
        Entity: tl.NewEntity(1, 1, 5, 3),
        Level: level,
        Face: entities.NORTH,
    }
    level.AddEntity(&spaceship)
    game.Screen().SetLevel(level)
    game.Start()
}
