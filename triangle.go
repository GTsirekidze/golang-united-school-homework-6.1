package golang_united_school_homework

import (
	"math"
)

type Triangle struct {
	Side float64
}

func (triangle Triangle) CalcPerimeter() float64 {
	return triangle.Side * 3
}

func (triangle Triangle) CalcArea() float64 {
	return triangle.Side * triangle.Side * math.Sqrt(3) / 4
}

func (triangle Triangle) ReturnIfCircle() bool {
	return false
}
