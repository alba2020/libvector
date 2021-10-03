package vector

import (
	"fmt"
	"math"

	"github.com/alba2020/libvector/util"
)

type Vector struct {
	coordinates []float64
	dimension   int
}

func NewVector(coordinates ...float64) Vector {
	v := Vector{coordinates: coordinates}
	v.dimension = len(coordinates)
	return v
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector: %v", v.coordinates)
}

func (v Vector) Eq(other Vector) bool {
	return util.Equals(v.coordinates, other.coordinates)
}

// sum with another vector
func (v Vector) Sum(other Vector) Vector {
	sum := Vector{}
	if v.dimension != other.dimension {
		panic("bad dimensions")
	}
	sum.dimension = v.dimension
	sum.coordinates = make([]float64, sum.dimension)
	for i := 0; i < v.dimension; i++ {
		sum.coordinates[i] = v.coordinates[i] +
			other.coordinates[i]
	}
	return sum
}

// subtract another vector
func (v Vector) Sub(other Vector) Vector {
	sum := Vector{}
	if v.dimension != other.dimension {
		panic("Bad dimensions")
	}
	sum.dimension = v.dimension
	sum.coordinates = make([]float64, sum.dimension)
	for i := 0; i < v.dimension; i++ {
		sum.coordinates[i] = v.coordinates[i] -
			other.coordinates[i]
	}
	return sum
}

// multiply by number
func (v Vector) Mul(n float64) Vector {
	res := Vector{dimension: v.dimension}
	res.coordinates = make([]float64, v.dimension)
	for i := 0; i < v.dimension; i++ {
		res.coordinates[i] = v.coordinates[i] * n
	}
	return res
}

// magnitude of vector
func (v Vector) Mod() float64 {
	var sum float64
	for _, c := range v.coordinates {
		sum += c * c
	}
	// sum1 := _sum(_map(v.coordinates, func(x float64) float64{
	//      return x * x
	// }))
	return math.Sqrt(sum)
}

// unit vector
func (v Vector) Norm() Vector {
	mod := v.Mod()
	if mod == 0 {
		panic("division by zero")
	}
	return v.Mul(1 / mod)
}

// scalar product - a*b*cos(theta)
func (v Vector) Dot(other Vector) float64 {
	var sum float64
	for i, x := range v.coordinates {
		sum += x * other.coordinates[i]
	}
	return sum
}

// angle in radians
func (v Vector) Angle(other Vector) float64 {
	cos := v.Dot(other) / (v.Mod() * other.Mod())
	return math.Acos(cos)
}
