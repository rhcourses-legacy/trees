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

// addLeft adds a new element to the left of the current element.
func (e *Element[T]) addLeft() *Element[T] {
	e.left = NewElement[T]()
	return e.left
}

// addRight adds a new element to the right of the current element.
func (e *Element[T]) addRight() *Element[T] {
	e.right = NewElement[T]()
	return e.right
}

// SetValue sets the value of the element.
// If the element is empty, two new elements are added to the left and right.
func (e *Element[T]) SetValue(value T) {
	if e.IsEmpty() {
		e.addLeft()
		e.addRight()
	}
	e.value = value
}

// Value returns the Value of this element.
// Panics if the element is empty.
func (e *Element[T]) Value() T {
	if e.IsEmpty() {
		panic("Value method called on an empty node.")
	}
	return e.value
}

// InOrderValues returns a slice of values in the tree in in-order traversal.
func (e *Element[T]) InOrderValues() []T {
	if e.IsEmpty() {
		return []T{}
	}
	return append(append(e.left.InOrderValues(), e.value), e.right.InOrderValues()...)
}

// PreOrderValues returns a slice of values in the tree in pre-order traversal.
func (e *Element[T]) PreOrderValues() []T {
	if e.IsEmpty() {
		return []T{}
	}
	return append(append([]T{e.value}, e.left.PreOrderValues()...), e.right.PreOrderValues()...)
}

// PostOrderValues returns a slice of values in the tree in post-order traversal.
func (e *Element[T]) PostOrderValues() []T {
	if e.IsEmpty() {
		return []T{}
	}
	return append(append(e.left.PostOrderValues(), e.right.PostOrderValues()...), e.value)
}

// MermaidLines returns a slice of strings, each of which is a line in a Mermaid graph.
// The slice contains the node values and edges, but no header.
func (e *Element[T]) MermaidLines() []string {
	result := []string{}
	if e.IsEmpty() {
		return result
	}
	if !e.left.IsEmpty() {
		result = append(result, fmt.Sprintf("%v --> %v", e.value, e.left.value))
	}
	if !e.right.IsEmpty() {
		result = append(result, fmt.Sprintf("%v --> %v", e.value, e.right.value))
	}
	return append(append(result, e.left.MermaidLines()...), e.right.MermaidLines()...)
}
