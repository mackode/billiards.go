package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  col "golang.org/x/image/colornames"
  "github.com/quartercastle/vector"
  "os"
  "time"
)

func main() {
  a := app.New()
  w := a.NewWindow("Pool Billiard")

  width := float32(650)
  height := float32(700)
  w.Resize(fyne.NewSize(width, height))
  w.SetFixedSize(true)

  radius := float32(30)
  cue := drawCircle(col.White, 200, 200, radius)
  obj := drawCircle(col.Red, 400, 200, radius)
  cueBall := Ball{Ava: cue, Velo: vector.Vector{0, 0}}
  objBall := Ball{Ava: cue, Velo: vector.Vector{0, 0}}

  play := NewTapRect(width, height, func(pos fyne.Position) {
    shootBall(&cueBall, pos)
  })
  play.Resize(fyne.NewSize(width, height))

  objs := []fyne.CanvasObject{play, cue, obj}
  con := container.NewWithoutLayout(objs...)
  w.SetContent(con)

  w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
    key := string(ev.Name)
    switch key {
      case "Q":
        os.Exit(0)
    }
  })

  go func() {
    for {
      select {
        case <- time.After(time.Duration(5) * time.Millisecond):
          driveBall(&cueBall)
          driveBall(&objBall)
          slowBall(&cueBall)
          slowBall(&objBall)

          if detectCollision(&cueBall, &objBall) {
            collide(&cueBall, &objBall)
            fixOverlap(&cueBall, &objBall)
          }
          wallBounce(&cueBall, width, height, radius)
          wallBounce(&objBall, width, height, radius)
      }
    }
  }()

  w.ShowAndRun()
}
