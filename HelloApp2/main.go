package main

import (
    "fmt"
    "geometry"
    "github.com/hackebrot/turtle"
)

func main() {
    emoji, ok := turtle.Emojis["smiley"]
    if !ok {
        fmt.Println("No emoji found.")
    } else {
        fmt.Println(emoji.Char)
    }

    pt1 := geometry.Point{X: 2, Y: 3}
    fmt.Println(pt1)
    fmt.Println(pt1.Length())
}
