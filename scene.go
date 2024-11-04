package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	Layers   map[string]*Layer
	Order    []string // Udržuje pořadí vrstev podle názvu
	Clear    Drawer
	Renderer *sdl.Renderer
}

func NewScene(wnd *sdl.Window) Scene {
	width, height := wnd.GetSize()
	r, e := sdl.CreateRenderer(wnd, -1, sdl.RENDERER_ACCELERATED)

	if e != nil {
		panic("Cannot acquire Renderer.")
	}

	ret := Scene{
		Layers:   make(map[string]*Layer),
		Order:    []string{},
		Clear:    FillDraw{Dst: Rectangle{Max: Point{X: float64(width), Y: float64(height)}}},
		Renderer: r,
	}

	return ret
}
func (s *Scene) AddLayer(name string, layer *Layer) error {
	if _, exists := s.Layers[name]; exists {
		return fmt.Errorf("layer with name %s already exists", name)
	}
	s.Layers[name] = layer
	s.Order = append(s.Order, name) // Přidáme název do seznamu pro zachování pořadí
	return nil
}

func (s *Scene) RemoveLayer(name string) error {
	if _, exists := s.Layers[name]; !exists {
		return fmt.Errorf("layer with name %s not found", name)
	}
	delete(s.Layers, name)

	// Aktualizace pořadí - odstraníme název ze seznamu
	for i, layerName := range s.Order {
		if layerName == name {
			s.Order = append(s.Order[:i], s.Order[i+1:]...)
			break
		}
	}
	return nil
}

func (s *Scene) MoveLayer(name string, newIndex int) error {
	if _, exists := s.Layers[name]; !exists {
		return fmt.Errorf("layer with name %s not found", name)
	}

	if newIndex < 0 || newIndex >= len(s.Order) {
		return fmt.Errorf("new index %d out of bounds", newIndex)
	}

	// Najdeme aktuální index
	var currentIndex int
	for i, layerName := range s.Order {
		if layerName == name {
			currentIndex = i
			break
		}
	}

	// Přesuneme název vrstvy na novou pozici
	s.Order = append(s.Order[:currentIndex], s.Order[currentIndex+1:]...)
	s.Order = append(s.Order[:newIndex], append([]string{name}, s.Order[newIndex:]...)...)

	return nil
}

func (s *Scene) GetLayer(name string) (*Layer, error) {
	if layer, exists := s.Layers[name]; exists {
		return layer, nil
	}
	return nil, fmt.Errorf("layer with name %s not found", name)
}

func (s *Scene) IterateLayersInOrder() []*Layer {
	var orderedLayers []*Layer
	for _, name := range s.Order {
		if layer, exists := s.Layers[name]; exists {
			orderedLayers = append(orderedLayers, layer)
		}
	}
	return orderedLayers
}

func (s *Scene) Evaluate() error {
	layers_ord := s.IterateLayersInOrder()

	for _, l := range layers_ord {
		e := l.Evaluate()
		if e != nil {
			return e
		}
	}
	return nil
}
