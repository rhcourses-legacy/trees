package bintree

import "fmt"

// Element is a node in a binary tree.
// This is a generic type, so it can store any type of value,
// as long as it satisfies the constraints defined by ValueType.
type Element[T any] struct {
	value T
	left  *Element[T]
	right *Element[T]
}

// NewElement returns a new empty element.
func NewElement[T any]() *Element[T] {
	return &Element[T]{}
}

// IsEmpty returns true if the element is empty.
func (e *Element[T]) IsEmpty() bool {
	return e.left == nil && e.right == nil
}

// String returns the string representation of the element.
func (e *Element[T]) String() string {
	if e.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%v", e.value)
}
