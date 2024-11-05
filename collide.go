package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
  "github.com/quartercastle/vector"
)

type Ball struct {
  Ava *canvas.Circle
  Velo vector.Vector
}

func pos2Vector(ball *Ball) vector.Vector {
  radius := ball.Ava.Size().Width / 2
  pos := ball.Ava.Position()
  return vector.Vector{float64(pos.X + radius), float64(pos.Y + radius)}
}

func detectCollision(ball1, ball2 *Ball) bool {
  v1 := pos2Vector(ball1)
  v2 := pos2Vector(ball2)
  dist := v1.Sub(v2).Magnitude()
  diameter := float64(ball1.Ava.Size().Width)
  return dist <= diameter
}

func fixOverlap(ball1, ball2 *Ball) {
  v1 := pos2Vector(ball1)
  v2 := pos2Vector(ball2)
  diameter := float64(ball1.Ava.Size().Width)
  dist := v1.Sub(v2)
  if dist.Magnitude() < diameter {
    res := dist.Unit().Scale((diameter - dist.Magnitude()))
    cur := ball2.Ava.Position()
    ball2.Ava.Move(fyne.NewPos(cur.X - float32(res[0]), cur.Y - float32(res[1])))
  }
}

func collide(ball1, ball2 *Ball) {
  pos1 := pos2Vector(ball1)
  pos2 := pos2Vector(ball2)
  normal := pos1.Sub(pos2).Unit()
  relVel := ball1.Velo.Sub(ball2.Velo)
  sepVel := relVel.Dot(normal)
  sepVelVec := normal.Scale(-sepVel)
  ball1.Velo = ball1.Velo.Add(sepVelVec)
  ball2.Velo = ball2.Velo.Add(sepVelVec.Scale(-1))
}

func wellBounce(ball *Ball, width, height float32, radius float32) {
  b := pos2Vector(ball)
  var axis vector.Vector
  if b[0] <= float64(radius) || b[0] >= float64(width - radius) {
    axis = vector.Vector{0, 1}
  }
  if b[1] <= float64(radius) || b[1] >= float64(height - radius) {
    axis = vector.Vector{1, 0}
  }
  if axis != nil {
    angle := ball.Velo.Angle(axis)
    ball.Velo = ball.Velo.Rotate(2 * angle)
  }
}
