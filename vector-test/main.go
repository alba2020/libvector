package main

import (
	"fmt"

	"github.com/alba2020/libvector/util"
	"github.com/alba2020/libvector/vector"

	"github.com/shopspring/decimal"
)

func ex2() {
	v3 := vector.NewVector(8.813, -1.331, -6.247)
	fmt.Println(v3, "mod", v3.Mod())

	v4 := vector.NewVector(5.581, -2.136)
	fmt.Println(v4, "norm", v4.Norm())
	v5 := vector.NewVector(1.996, 3.108, -4.554)
	fmt.Println(v5, "mod", v5.Norm())
}

func ex3() {
	v1 := vector.NewVector(7.887, 4.138)
	w1 := vector.NewVector(-8.802, 6.776)
	fmt.Println(v1.Dot(w1)) // -41.382

	v2 := vector.NewVector(-5.955, -4.904, -1.874)
	w2 := vector.NewVector(-4.496, -8.755, 7.103)
	fmt.Println(v2.Dot(w2)) // 56.397

	v3 := vector.NewVector(3.183, -7.627)
	w3 := vector.NewVector(-2.668, 5.319)
	fmt.Println("rad", v3.Angle(w3)) // 3.072

	v4 := vector.NewVector(7.35, 0.221, 5.188)
	w4 := vector.NewVector(2.751, 8.259, 3.985)
	fmt.Println("deg", util.ToDegrees(v4.Angle(w4))) // 60.276
}

func main() {
	// basic()
	// sum()
	// ex1()
	// ex2()
	// ex3()
	x := decimal.NewFromFloat32(12.2312312)
	fmt.Println(x)
	fmt.Println(x.Round(3))
}
