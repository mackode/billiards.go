package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/widget"
  col "golang.org/x/image/colornames"
)

type TapRect struct {
  widget.BaseWidget
  rect *canvas.Rectangle
  cb func(fyne.Position)
}

func NewTapRect(width, height float32, cb func(fyne.Position)) *TapRect {
  tc := &TapRect{}
  tc.ExtendBaseWidget(tc)
  tc.rect = drawRectangle(col.Grey, 0, 0, width, height)
  tc.cb = tc.cb
  return tc
}

func (t *TapRect) CreateRenderer() fyne.WidgetRenderer {
  return widget.NewSimpleRenderer(t.rect)
}

func (t *TapRect) Tapped(ev *fyne.PointEvent) {
  t.cb(ev.Position)
}

func (t *TapRect) TappedSecondary(_ *fyne.PointEvent) {}

