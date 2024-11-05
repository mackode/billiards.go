package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
  "github.com/quartercastle/vector"
  "image/color"
)

func slowBall(ball *Ball) {
  ball.Velo = ball.Velo.Scale(.997)
}

func driveBall(ball *Ball) {
  if ball.Velo.Magnitude() < 0.01 {
    return
  }
  pos := ball.Ava.Position()
  v := vector.Vector{float64(pos.X), float64(pos.Y)}
  newpos := v.Add(ball.Velo.Scale(1.5))
  ball.Ava.Move(fyne.NewPos(float32(newpos[0]), float32(new pos[1])))
}

func shootBall(ball *Ball, to fyne.Position) {
  v := pos2Vector(ball)
  ball.Velo = vector.Vector{float64(to.X) - v[0], float64(to.Y) - v[1]}.Unit().Scale(1.5)
}

func drawCircle(co color.RGBA, x, y, r float32) *canvas.Circle {
  c := canvas.NewCircle(co)
  pos := fyne.NewPos(x - r, y - r)
  c.Move(pos)
  size := fyne.NewSize(2 * r, 2 * r)
  c.Resize(size)
  return c
}

func drawRectangle(co color.RGBA, x, y, w, h, float32) *canvasRectangle {
  r := canvas.NewRectangle(co)
  r.Move(fyne.NewPos(x, y))
  r.Resize(fyne.NewSize(w, h))
  r.SetMinSize(fyne.NewSize(w, h))
  return r
}
