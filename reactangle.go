package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rectangle Rectangle) CalcPerimeter() float64 {
	return (rectangle.Height + rectangle.Weight) * 2
}

func (rectangle Rectangle) CalcArea() float64 {
	return rectangle.Height * rectangle.Weight
}

func (rectangle Rectangle) ReturnIfCircle() bool {
	return false
}
