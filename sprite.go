package main

import "github.com/veandco/go-sdl2/sdl"

// Definice struktury pro sprite
type Sprite struct {
	Rect          Rectangle // Velikost spritu (bounding box)
	Movement      Vector
	Accelleration Vector
	Matrix        *Matrix // Transformace spritu
	Texture       any     // Todo attach a texture
	Audio         any
	D             Drawer
}

type Spriter interface {
	Tick()
	Collide(*Sprite) []Point
	Destroy()
	ApplyPhysics(float32)
	Draw(*sdl.Renderer) error // draw itself to a render target
}

// Funkce pro vytvoření nového spritu
func NewSprite(width, height float64) *Sprite {
	s := Sprite{
		Rect: Rectangle{
			Min: Point{X: 0, Y: 0},
			Max: Point{X: width, Y: height},
		},
		Matrix: IdentityMatrix(3),
	}

	return &s
}

func (s *Sprite) Tick() {
	// stub
}

func (s *Sprite) Transform() {
	// stub
}

func (s *Sprite) Collide(ss *Sprite) []Point {
	b, pts := s.Rect.Intersects(ss.Rect)
	if b {
		return pts
	}

	return nil
}

func (s *Sprite) Destroy() {
	//stub
}

func (s *Sprite) ApplyPhysics(dt float64) {
	// todo better physics
	s.Movement.X += dt * s.Accelleration.X
	s.Movement.Y += dt * s.Accelleration.Y

	dv := s.Movement.Scale(dt)

	s.Rect.Min = s.Rect.Min.AddVector(dv)
	s.Rect.Max = s.Rect.Min.AddVector(dv)
}

func (s *Sprite) Draw(r *sdl.Renderer) error {
	s.D = &FillDraw{Dst: s.Rect, Color: sdl.Color{R: 255, G: 0, B: 255, A: 255}}
	return s.D.Draw(r)
}
