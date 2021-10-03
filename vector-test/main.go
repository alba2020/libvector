package main

import (
	"fmt"
	"libvector/vector"
)

func basic() {
	v1 := vector.NewVector(3, 2, 1)
	v2 := vector.NewVector(3, 2, 1)
	v3 := vector.NewVector(4, 3, 2)

	fmt.Println("v1", v1)

	fmt.Println("v1 == v1?", v1.Eq(v1))
	fmt.Println("v1 == v2?", v1.Eq(v2))
	fmt.Println("v1 == v3?", v1.Eq(v3))
}

func sum() {
	v1 := vector.NewVector(1, 2, 3, 4)
	v2 := vector.NewVector(8, 7, 6, 5)

	sum1 := v1.Sum(v2)
	fmt.Println(sum1)

	sub1 := v2.Sub(v1)
	fmt.Println(sub1)

	mul := v2.Mul(3)
	fmt.Println(mul)
}

func ex1() {
	v1 := vector.NewVector(8.218, -9.341)
	v2 := vector.NewVector(-1.129, 2.111)
	fmt.Println(v1.Sum(v2))

	v3 := vector.NewVector(7.119, 8.215)
	v4 := vector.NewVector(-8.223, 0.878)
	fmt.Println(v3.Sub(v4))

	v5 := vector.NewVector(1.671, -1.012, -0.318)
	fmt.Println(v5.Mul(7.41))
}

func ex2() {
	v1 := vector.NewVector(3, 4)
	fmt.Println(v1, "mod", v1.Mod())
	v2 := vector.NewVector(-0.221, 7.437)
	fmt.Println(v2, "mod", v2.Mod())
	v3 := vector.NewVector(8.813, -1.331, -6.247)
	fmt.Println(v3, "mod", v3.Mod())

	v4 := vector.NewVector(5.581, -2.136)
	fmt.Println(v4, "norm", v4.Norm())
	v5 := vector.NewVector(1.996, 3.108, -4.554)
	fmt.Println(v5, "mod", v5.Norm())
}

func main() {
	// basic()
	// sum()
	// ex1()
	ex2()
}
