package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Surface struct {
	Surface *sdl.Surface
}

// Funkce pro načtení SDL_Surface ze souboru, který je již uložen v cache
func LoadSurfaceFromCache(cache *FileCache, path string) (*Surface, error) {
	// Získej soubor z cache
	data, err := cache.LoadFile(path)
	if err != nil {
		return nil, err
	}

	// Načti povrch (Surface) z paměti (byte array)
	rwOps, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}
	surface, err := img.LoadRW(rwOps, false)
	if err != nil {
		return nil, err
	}

	return &Surface{Surface: surface}, nil
}

// Uvolnění SDL_Surface
func (s *Surface) Free() {
	if s.Surface != nil {
		s.Surface.Free() // Probably frees also cache!
	}
}
