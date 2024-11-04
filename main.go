package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}

	// Příklad bodu
	point := Point{X: 2, Y: 3}

	// Příklad transformační matice (např. posun o (1, 2))
	translationMatrix := NewMatrix(3, 3)
	translationMatrix.Data[0][0] = 1
	translationMatrix.Data[0][1] = 0
	translationMatrix.Data[0][2] = 1 // posun X
	translationMatrix.Data[1][0] = 0
	translationMatrix.Data[1][1] = 1
	translationMatrix.Data[1][2] = 2 // posun Y
	translationMatrix.Data[2][0] = 0
	translationMatrix.Data[2][1] = 0
	translationMatrix.Data[2][2] = 1

	// Transformace bodu
	transformedPoint := TransformPoint(point, translationMatrix)

	// Výpis výsledku
	fmt.Printf("Původní bod: (%f, %f)\n", point.X, point.Y)
	fmt.Printf("Transformovaný bod: (%f, %f)\n", transformedPoint.X, transformedPoint.Y)
}
