package vector

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestEq(t *testing.T) {
	v1 := NewVector(1, 2, 3.14)
	v2 := NewVector(1, 2, 3.14)
	v3 := NewVector(3.14, 2, 1)

	if ! v1.Eq(v1) {
		t.Error("v1 != v1")
	}

	if ! v1.Eq(v2) {
		t.Error("v1 != v2")
	}

	if v1.Eq(v3) {
		t.Errorf("v1 %v must not equal v3 %v", v1, v3)
	}
}

func TestRound(t *testing.T) {
	v1 := NewVector(1.234567)
	if ! v1.Round(2).Eq(NewVector(1.23)) {
		t.Error("round error")
	}
	rounded := v1.Round(3)
	if ! rounded.Eq(NewVector(1.235)) {
		t.Errorf("round error %v", rounded)
	}
}

func TestSum2D(t *testing.T) {
	v1 := NewVector(8.218, -9.341)
	v2 := NewVector(-1.129, 2.111)
	sum := v1.Sum(v2)
    if ! sum.Round(3).Eq(NewVector(7.089, -7.230)) {
        t.Errorf("Wrong sum %v", sum)
    }
}

func TestSub2D(t *testing.T) {
	v1 := NewVector(7.119, 8.215)
	v2 := NewVector(-8.223, 0.878)
	res := v1.Sub(v2)
    if ! res.Round(3).Eq(NewVector(15.342, 7.337)) {
        t.Errorf("Wrong res %v", res)
    }
}

func TestMul3D(t *testing.T) {
	v := NewVector(1.671, -1.012, -0.318)
	res := v.Mul(7.41)
    if ! res.Round(3).Eq(NewVector(12.382, -7.499, -2.356)) {
        t.Errorf("Wrong res %v", res)
    }
}

func TestMod2D(t *testing.T) {
	v := NewVector(-0.221, 7.437)
	res := v.Mod().Round(3)
    if ! res.Equal(decimal.NewFromFloat(7.44)) {
        t.Errorf("Wrong res %v", res)
    }
}
