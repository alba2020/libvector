package vector

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Vector struct {
	coordinates []decimal.Decimal
	dimension   int
}

func NewVector(coordinates ...float64) Vector {
	v := Vector{dimension: len(coordinates)}
	cs := make([]decimal.Decimal, v.dimension)
	for i, c := range coordinates {
		cs[i] = decimal.NewFromFloat(c)
	}
	v.coordinates = cs
	return v
}

func (v Vector) Round(places int32) Vector {
	cs := make([]decimal.Decimal, len(v.coordinates))
	for i, c := range v.coordinates {
		cs[i] = c.Round(places)
	}
	v2 := Vector{
		dimension: v.dimension,
		coordinates: cs,
	}
	return v2
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector: %v", v.coordinates)
}

func (v Vector) Eq(other Vector) bool {
	if len(v.coordinates) != len(other.coordinates) {
		return false
	}
	for i, c := range v.coordinates {
		if ! c.Equal(other.coordinates[i]) {
			return false
		}
	}
	return true
}

// sum with another vector
func (v Vector) Sum(other Vector) Vector {
	if v.dimension != other.dimension {
		panic("bad dimensions")
	}

	sum := Vector{
		dimension: v.dimension,
		coordinates: make([]decimal.Decimal, v.dimension),
	}

	for i, c := range v.coordinates {
		sum.coordinates[i] = c.Add(other.coordinates[i])
	}
	return sum
}

// subtract another vector
func (v Vector) Sub(other Vector) Vector {
	if v.dimension != other.dimension {
		panic("bad dimensions")
	}

	sum := Vector{
		dimension: v.dimension,
		coordinates: make([]decimal.Decimal, v.dimension),
	}

	for i, c := range v.coordinates {
		sum.coordinates[i] = c.Sub(other.coordinates[i])
	}
	return sum
}

// multiply by number
func (v Vector) Mul(n float64) Vector {
	res := Vector{dimension: v.dimension}
	res.coordinates = make([]decimal.Decimal, v.dimension)
	for i, c := range v.coordinates {
		res.coordinates[i] = c.Mul(decimal.NewFromFloat(n))
	}
	return res
}

// magnitude of vector
func (v Vector) Mod() decimal.Decimal {
	var sum decimal.Decimal
	for _, c := range v.coordinates {
		sum = sum.Add(c.Mul(c))
	}
	return sum.Pow(decimal.NewFromFloat(0.5))
}

// // unit vector
// func (v Vector) Norm() Vector {
// 	mod := v.Mod()
// 	if mod == 0 {
// 		panic("division by zero")
// 	}
// 	return v.Mul(1 / mod)
// }

// // scalar product - a*b*cos(theta)
// func (v Vector) Dot(other Vector) float64 {
// 	var sum float64
// 	for i, x := range v.coordinates {
// 		sum += x * other.coordinates[i]
// 	}
// 	return sum
// }

// // angle in radians
// func (v Vector) Angle(other Vector) float64 {
// 	cos := v.Dot(other) / (v.Mod() * other.Mod())
// 	return math.Acos(cos)
// }
