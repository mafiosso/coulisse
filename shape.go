package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	Min Point
	Max Point
}

func (r Rectangle) ToNative() sdl.Rect {
	return sdl.Rect{X: int32(r.Min.X), Y: int32(r.Min.Y), W: int32(r.Max.X - r.Min.X), H: int32(r.Max.Y - r.Min.Y)}
}

// Funkce pro kontrolu, zda je bod uvnitř obdélníku
func (r Rectangle) Contains(p Point) bool {
	return p.X >= r.Min.X && p.X <= r.Max.X && p.Y >= r.Min.Y && p.Y <= r.Max.Y
}

// Funkce pro detekci kolize mezi dvěma obdélníky
func (r1 Rectangle) Intersects(r2 Rectangle) (bool, []Point) {
	// Zjistíme, jestli se obdélníky překrývají
	if r1.Max.X < r2.Min.X || r1.Min.X > r2.Max.X || r1.Max.Y < r2.Min.Y || r1.Min.Y > r2.Max.Y {
		return false, nil
	}

	// Seznam rohů r1
	corners := []Point{
		{r1.Min.X, r1.Min.Y}, // levý dolní roh
		{r1.Max.X, r1.Min.Y}, // pravý dolní roh
		{r1.Min.X, r1.Max.Y}, // levý horní roh
		{r1.Max.X, r1.Max.Y}, // pravý horní roh
	}

	// Seznam bodů uvnitř r2
	var pointsInside []Point
	for _, corner := range corners {
		if r2.Contains(corner) {
			pointsInside = append(pointsInside, corner)
		}
	}

	return true, pointsInside
}

type PointInterface interface {
	GetX() float64
	GetY() float64
	Move(dx, dy float64)
	DistanceTo(other Point) float64
}

func (p *Point) GetX() float64 {
	return p.X
}

func (p *Point) GetY() float64 {
	return p.Y
}

func (p *Point) Move(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

func (p *Point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(other.X-p.X, 2) + math.Pow(other.Y-p.Y, 2))
}

func BoundingBox(points []Point) Rectangle {
	if len(points) == 0 {
		return Rectangle{}
	}

	minX := math.Inf(1) // Nastavení na nekonečno
	minY := math.Inf(1)
	maxX := math.Inf(-1) // Nastavení na -nekonečno
	maxY := math.Inf(-1)

	// Procházení bodů a hledání minimálních a maximálních hodnot
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return Rectangle{
		Min: Point{X: minX, Y: minY},
		Max: Point{X: maxX, Y: maxY},
	}
}

type Vector struct {
	X float64
	Y float64
}

func (p Point) AddVector(v Vector) Point {
	return Point{
		X: p.X + v.X,
		Y: p.Y + v.Y,
	}
}

// Metoda pro sečtení dvou vektorů
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

// Metoda pro násobení vektoru skalárem
func (v Vector) Scale(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

// Metoda pro výpočet délky vektoru
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Normalize() Vector {
	length := v.Length()
	if length == 0 {
		return Vector{X: 0, Y: 0}
	}
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}

// Metoda pro výpočet kolmého vektoru
func (v Vector) Perpendicular() Vector {
	return Vector{
		X: -v.Y, // Změníme znaménko Y
		Y: v.X,  // X se stane Y
	}
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Funkce pro výpočet úhlu mezi dvěma vektory v radiánech
func AngleBetween(v1, v2 Vector) float64 {
	dotProduct := v1.Dot(v2)
	length1 := v1.Length()
	length2 := v2.Length()

	// Kontrola délky vektorů, aby nedošlo k dělení nulou
	if length1 == 0 || length2 == 0 {
		return 0 // nebo můžeš vrátit chybu, pokud jsou vektory nulové
	}

	// Vypočítání kosinu úhlu
	cosTheta := dotProduct / (length1 * length2)

	// Ochrana proti hodnotám mimo interval [-1, 1] pro arccos
	if cosTheta < -1 {
		cosTheta = -1
	} else if cosTheta > 1 {
		cosTheta = 1
	}

	// Vypočítání úhlu
	return math.Acos(cosTheta) // úhel v radiánech
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

// Funkce pro převod stupňů na radiány
func DegreesToRadians(degrees float64) float64 {
	return degrees * (3.141592653589793 / math.Pi)
}

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]float64
}

// Konstruktor pro vytvoření matice
func NewMatrix(rows, columns int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, columns)
	}
	return &Matrix{Rows: rows, Columns: columns, Data: data}
}

// Metoda pro sčítání dvou matic
func (m1 *Matrix) Add(m2 *Matrix) *Matrix {
	if m1.Rows != m2.Rows || m1.Columns != m2.Columns {
		return nil // nebo můžeš vrátit chybu
	}
	result := NewMatrix(m1.Rows, m1.Columns)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Columns; j++ {
			result.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
		}
	}
	return result
}

// Metoda pro odčítání dvou matic
func (m1 *Matrix) Subtract(m2 *Matrix) *Matrix {
	if m1.Rows != m2.Rows || m1.Columns != m2.Columns {
		return nil // nebo můžeš vrátit chybu
	}
	result := NewMatrix(m1.Rows, m1.Columns)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Columns; j++ {
			result.Data[i][j] = m1.Data[i][j] - m2.Data[i][j]
		}
	}
	return result
}

// Metoda pro násobení matice s jinou maticí
func (m1 *Matrix) Multiply(m2 *Matrix) *Matrix {
	if m1.Columns != m2.Rows {
		return nil // nebo můžeš vrátit chybu
	}
	result := NewMatrix(m1.Rows, m2.Columns)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m2.Columns; j++ {
			for k := 0; k < m1.Columns; k++ {
				result.Data[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}
	return result
}

// Metoda pro transpozici matice
func (m *Matrix) Transpose() *Matrix {
	result := NewMatrix(m.Columns, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			result.Data[j][i] = m.Data[i][j]
		}
	}
	return result
}

func IdentityMatrix(size int) *Matrix {
	result := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		result.Data[i][i] = 1
	}
	return result
}

func RotationMatrix(theta float64) *Matrix {
	result := NewMatrix(2, 2)
	result.Data[0][0] = math.Cos(theta)
	result.Data[0][1] = -math.Sin(theta)
	result.Data[1][0] = math.Sin(theta)
	result.Data[1][1] = math.Cos(theta)
	return result
}

// Funkce pro vytvoření posunové matice
func TranslationMatrix(tX, tY float64) *Matrix {
	result := NewMatrix(3, 3)
	result.Data[0][0] = 1
	result.Data[0][1] = 0
	result.Data[0][2] = tX
	result.Data[1][0] = 0
	result.Data[1][1] = 1
	result.Data[1][2] = tY
	result.Data[2][0] = 0
	result.Data[2][1] = 0
	result.Data[2][2] = 1
	return result
}

// Funkce pro vytvoření škálovací matice
func ScalingMatrix(sX, sY float64) *Matrix {
	result := NewMatrix(3, 3)
	result.Data[0][0] = sX
	result.Data[0][1] = 0
	result.Data[0][2] = 0
	result.Data[1][0] = 0
	result.Data[1][1] = sY
	result.Data[1][2] = 0
	result.Data[2][0] = 0
	result.Data[2][1] = 0
	result.Data[2][2] = 1
	return result
}

// Metoda pro výpis matice
func (m *Matrix) Print() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			fmt.Printf("%f ", m.Data[i][j])
		}
		fmt.Println()
	}
}

// Funkce pro výpočet inverzní matice pomocí Gaussovy eliminace
func (m *Matrix) Inverse() (*Matrix, error) {
	if m.Rows != m.Columns {
		return nil, fmt.Errorf("matice musí být čtvercová")
	}

	// Vytvoření rozšířené matice (m, I)
	augmented := NewMatrix(m.Rows, 2*m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			augmented.Data[i][j] = m.Data[i][j]
		}
		augmented.Data[i][i+m.Columns] = 1 // Jednotková matice na pravé straně
	}

	// Gaussova eliminace
	for i := 0; i < m.Rows; i++ {
		// Hledání maximálního prvku v aktuálním sloupci
		maxRow := i
		for k := i + 1; k < m.Rows; k++ {
			if math.Abs(augmented.Data[k][i]) > math.Abs(augmented.Data[maxRow][i]) {
				maxRow = k
			}
		}

		// Výměna maximálního řádku s aktuálním řádkem
		if maxRow != i {
			augmented.Data[i], augmented.Data[maxRow] = augmented.Data[maxRow], augmented.Data[i]
		}

		// Normalizace aktuálního řádku
		divisor := augmented.Data[i][i]
		if divisor == 0 {
			return nil, fmt.Errorf("matice nemá inverzi")
		}
		for j := 0; j < 2*m.Columns; j++ {
			augmented.Data[i][j] /= divisor
		}

		// Eliminuje ostatní řádky
		for k := 0; k < m.Rows; k++ {
			if k != i {
				factor := augmented.Data[k][i]
				for j := 0; j < 2*m.Columns; j++ {
					augmented.Data[k][j] -= factor * augmented.Data[i][j]
				}
			}
		}
	}

	// Extrakce inverzní matice
	inverse := NewMatrix(m.Rows, m.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			inverse.Data[i][j] = augmented.Data[i][j+m.Columns]
		}
	}

	return inverse, nil
}

// Funkce pro výpočet přechodové matice mezi dvěma maticemi
func TransitionMatrix(from, to *Matrix) (*Matrix, error) {
	if from.Rows != from.Columns || to.Rows != to.Columns {
		return nil, fmt.Errorf("oba vstupy musí být čtvercové matice")
	}

	// Můžeme zde implementovat logiku pro výpočet přechodové matice
	// Například můžeme použít inverzní matici
	inverseFrom, err := from.Inverse()
	if err != nil {
		return nil, err
	}

	transition := inverseFrom.Multiply(to)
	return transition, nil
}

func TransformPoint(p Point, m *Matrix) Point {
	if m.Rows != 3 || m.Columns != 3 {
		panic("Transformační matice musí být 3x3")
	}

	// Převod bodu na homogenní souřadnice
	x := p.X
	y := p.Y
	w := 1.0 // Homogenní souřadnice

	// Transformace bodu
	transformedX := m.Data[0][0]*x + m.Data[0][1]*y + m.Data[0][2]*w
	transformedY := m.Data[1][0]*x + m.Data[1][1]*y + m.Data[1][2]*w

	return Point{X: transformedX, Y: transformedY}
}
