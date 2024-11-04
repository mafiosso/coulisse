package main

import "github.com/veandco/go-sdl2/sdl"

type Drawer interface {
	Draw(dest *sdl.Renderer) error
}

type BltDraw struct {
	Src Rectangle
	Dst Rectangle
	S   Surface
	T   *sdl.Texture
}

func (l *BltDraw) Draw(renderer *sdl.Renderer) error {
	// Pokud textura ještě neexistuje, převedeme surface na texturu
	if l.T == nil {
		texture, err := renderer.CreateTextureFromSurface(l.S.Surface)
		if err != nil {
			return err
		}
		l.T = texture
	}

	// Nastavíme vykreslovací oblast podle l.Rect a vykreslíme texturu do rendereru
	rct := l.Dst.ToNative()
	err := renderer.Copy(l.T, nil, &rct)

	if err != nil {
		return err
	}
	return nil
}

type FillDraw struct {
	Dst   Rectangle
	Color sdl.Color
}

func (f FillDraw) Draw(renderer *sdl.Renderer) error {
	// Nastavení barvy (RGBA)
	err := renderer.SetDrawColor(f.Color.R, f.Color.G, f.Color.B, 255)
	if err != nil {
		return err
	}

	// Vyčištění rendereru (a tedy vyplnění barvou)
	err = renderer.Clear()
	if err != nil {
		return err
	}

	return nil
}

type PolyDraw struct {
	Src []Point
	Dst []Point // Src and Dst must be of same dimensionality
	S   Surface
}
