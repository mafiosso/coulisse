package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Layer struct {
	Rect    Rectangle
	Sprites []Spriter
	Effects []any
}

func (l *Layer) AddSprite(s Spriter) {
	l.Sprites = append(l.Sprites, s)
}

func (l *Layer) RemoveSprite(n int) error {
	if n < 0 || n >= len(l.Sprites) {
		return fmt.Errorf("index out of bounds")
	}
	l.Sprites = append(l.Sprites[:n], l.Sprites[n+1:]...)
	return nil
}

func (l *Layer) MoveSprite(n, newIndex int) error {
	if n < 0 || n >= len(l.Sprites) || newIndex < 0 || newIndex >= len(l.Sprites) {
		return fmt.Errorf("index out of bounds")
	}
	sprite := l.Sprites[n]
	l.Sprites = append(l.Sprites[:n], l.Sprites[n+1:]...)
	l.Sprites = append(l.Sprites[:newIndex], append([]Spriter{sprite}, l.Sprites[newIndex:]...)...)
	return nil
}

func (l *Layer) AddEffect(e any) {
	l.Effects = append(l.Effects, e)
}

func (l *Layer) RemoveEffect(n int) error {
	if n < 0 || n >= len(l.Effects) {
		return fmt.Errorf("index out of bounds")
	}
	l.Effects = append(l.Effects[:n], l.Effects[n+1:]...)
	return nil
}

func (l *Layer) MoveEffect(n, newIndex int) error {
	if n < 0 || n >= len(l.Effects) || newIndex < 0 || newIndex >= len(l.Effects) {
		return fmt.Errorf("index out of bounds")
	}
	effect := l.Effects[n]
	l.Effects = append(l.Effects[:n], l.Effects[n+1:]...)
	l.Effects = append(l.Effects[:newIndex], append([]any{effect}, l.Effects[newIndex:]...)...)
	return nil
}

func (l *Layer) SetRect(r Rectangle) {
	l.Rect = r
}

func (l *Layer) ResizeRect(width, height float64) {
	l.Rect.Max.X = width
	l.Rect.Max.Y = height
}

func (l *Layer) MoveRect(x, y float64) {
	l.Rect.Min.X += x
	l.Rect.Min.Y += y

	l.Rect.Max.X += x
	l.Rect.Max.Y += y
}

func (l *Layer) Evaluate() error {
	// stub
	for _, s := range l.Sprites {
		s.Tick()
	}

	return nil

}

func (l *Layer) Draw(r *sdl.Renderer) error {
	for _, s := range l.Sprites {
		s.Draw(r)
	}

	return nil
}
