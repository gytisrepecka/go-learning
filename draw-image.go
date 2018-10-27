// Copyright 2018 Gytis Repečka. All rights reserved.
// Use of this source code is governed by a GNU GPL
// license that can be found in the LICENSE file.

package main

import (
  "fmt"
  "os"
  "image"
  "image/color"
  "image/png"
  "image/draw"
)

type Circle struct {
  imageToDraw draw.Image
  x, y, r int
  c color.Color
}

// https://en.wikipedia.org/wiki/Midpoint_circle_algorithm
func (circle *Circle) Draw() {
  x := circle.r - 1
  y := 0
  dx := 1
  dy := 1
  err := dx - (circle.r * 2)

  for x > y {
    circle.imageToDraw.Set(circle.x + x, circle.y + y, circle.c)
    circle.imageToDraw.Set(circle.x + y, circle.y + x, circle.c)
    circle.imageToDraw.Set(circle.x - y, circle.y + x, circle.c)
    circle.imageToDraw.Set(circle.x - x, circle.y + y, circle.c)
    circle.imageToDraw.Set(circle.x - x, circle.y - y, circle.c)
    circle.imageToDraw.Set(circle.x - y, circle.y - x, circle.c)
    circle.imageToDraw.Set(circle.x + y, circle.y - x, circle.c)
    circle.imageToDraw.Set(circle.x + x, circle.y - y, circle.c)

    if err <= 0 {
      y++
      err += dy
      dy += 2
    }
    if err > 0 {
      x--
      dx += 2
      err += dx - (circle.r * 2)
    }

  }
}

func main() {
  outFileName := "draw-image.png"

  // Create 400x300 image
  imageToDraw := image.NewRGBA(image.Rect(0, 0, 400, 300))

  var myCircle Circle
  

  for r := 130; 0 < r; r-- {
    myCircle = Circle{imageToDraw, 200, 150, r, color.RGBA{0, 200, 0, 255}}
    myCircle.Draw()
  }

  f, errf := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE, 0600)
  if errf != nil {
    fmt.Println(errf)
    return
  }
  defer f.Close()
  png.Encode(f, imageToDraw)
  fmt.Printf("Image written: %s\n", outFileName)
}
