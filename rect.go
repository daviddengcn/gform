package gform

import (
    "github.com/AllenDang/w32"
)

type Rect w32.RECT

func NewEmptyRect() *Rect {
    var newRect Rect
    w32.SetRectEmpty((*w32.RECT)(&newRect))

    return &newRect
}

func NewRect(left, top, right, bottom int) *Rect {
    return &Rect{Left: int32(left), Top: int32(top), Right: int32(right), Bottom: int32(bottom)}
}

func (this *Rect) Data() (left, top, right, bottom int) {
    left = int(this.Left)
    top = int(this.Top)
    right = int(this.Right)
    bottom = int(this.Bottom)
    return
}

func (r *Rect) GetW32Rect() *w32.RECT {
    return (*w32.RECT)(r)
}

func (r *Rect) Set(left, top, right, bottom int) {
    w32.SetRect((*w32.RECT)(r), left, top, right, bottom)
}

func (r *Rect) IsEqual(rect *Rect) bool {
    return w32.EqualRect((*w32.RECT)(r), (*w32.RECT)(rect))
}

func (r *Rect) Inflate(x, y int) {
    w32.InflateRect((*w32.RECT)(r), x, y)
}

func (r *Rect) Intersect(src *Rect) {
    w32.IntersectRect((*w32.RECT)(r), (*w32.RECT)(r), (*w32.RECT)(src))
}

func (r *Rect) IsEmpty() bool {
    return w32.IsRectEmpty((*w32.RECT)(r))
}

func (r *Rect) Offset(x, y int) {
    w32.OffsetRect((*w32.RECT)(r), x, y)
}

func (r *Rect) IsPointIn(x, y int) bool {
    return w32.PtInRect((*w32.RECT)(r), x, y)
}

func (r *Rect) Substract(src *Rect) {
    w32.SubtractRect((*w32.RECT)(r), (*w32.RECT)(r), (*w32.RECT)(src))
}

func (r *Rect) Union(src *Rect) {
    w32.UnionRect((*w32.RECT)(r), (*w32.RECT)(r), (*w32.RECT)(src))
}
