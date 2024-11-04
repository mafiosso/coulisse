package main

import (
	"io/ioutil"
	"sync"
)

// Definice struktury pro FileCache
type FileCache struct {
	cache       map[string][]byte
	order       []string // Seznam klíčů pro správu FIFO
	mu          sync.Mutex
	maxSize     int64 // Maximální velikost cache
	currentSize int64 // Aktuální velikost cache
}

// Funkce pro vytvoření nového FileCache
func NewFileCache(maxSize int64) *FileCache {
	return &FileCache{
		cache:   make(map[string][]byte),
		order:   []string{},
		maxSize: maxSize,
	}
}

// Funkce pro načtení obsahu souboru a uložení do cache
func (fc *FileCache) LoadFile(path string) ([]byte, error) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	// Zkontroluj, jestli je soubor již v cache
	if data, exists := fc.cache[path]; exists {
		return data, nil
	}

	// Načti soubor
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Zjisti velikost nového záznamu
	newSize := int64(len(data))

	// Pokud překročíme maximální velikost, odstraň nejstarší záznamy
	for fc.currentSize+newSize > fc.maxSize && len(fc.order) > 0 {
		fc.removeOldest()
	}

	// Ulož data do cache
	fc.cache[path] = data
	fc.order = append(fc.order, path) // Přidej klíč do pořadí
	fc.currentSize += newSize         // Aktualizuj aktuální velikost

	return data, nil
}

// Funkce pro získání obsahu souboru z cache
func (fc *FileCache) GetFile(path string) ([]byte, bool) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	data, exists := fc.cache[path]
	return data, exists
}

// Funkce pro odstranění souboru z cache
func (fc *FileCache) RemoveFile(path string) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	// Zkontroluj, jestli soubor existuje v cache
	if data, exists := fc.cache[path]; exists {
		// Odeber soubor z cache a aktualizuj velikost
		delete(fc.cache, path)
		fc.currentSize -= int64(len(data))

		// Odeber klíč z pořadí
		for i, key := range fc.order {
			if key == path {
				fc.order = append(fc.order[:i], fc.order[i+1:]...)
				break
			}
		}
		return true
	}
	return false
}

// Soukromá funkce pro odstranění nejstaršího záznamu (používá se při překročení limitu)
func (fc *FileCache) removeOldest() {
	if len(fc.order) > 0 {
		oldestKey := fc.order[0]
		if data, exists := fc.cache[oldestKey]; exists {
			delete(fc.cache, oldestKey)
			fc.currentSize -= int64(len(data))
			fc.order = fc.order[1:] // Odeber nejstarší klíč
		}
	}
}

// Funkce pro vyčištění celé cache
func (fc *FileCache) ClearCache() {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	// Vyprázdni mapu cache a seznam klíčů
	fc.cache = make(map[string][]byte)
	fc.order = []string{}
	fc.currentSize = 0
}
