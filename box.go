package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("shapesCapacity Out Of Range")
	}
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}
	newBox := *b
	shape := b.shapes[i]
	newBox.shapes = append(newBox.shapes[:i], newBox.shapes[i+1:]...)
	*b = newBox
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}
	b.shapes[i] = shape

	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	ans := float64(0)
	for _, item := range b.shapes {
		ans += item.CalcPerimeter()
	}
	return ans
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	ans := float64(0)
	for _, item := range b.shapes {
		ans += item.CalcArea()
	}
	return ans
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	isCircle := false
	Error := errors.New("circles are not in list")
	for i, item := range b.shapes {
		if item.ReturnIfCircle() {
			newBox := *b
			newBox.shapes = append(newBox.shapes[:i], newBox.shapes[i+1:]...)
			*b = newBox
			isCircle = true
			i--
		}
	}

	if isCircle {
		Error = nil
	}
	return Error
}
