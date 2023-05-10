package bintree

import "strings"

// BinTree is a generic binary tree.
type BinTree[T any] struct {
	root *Element[T]
}

// NewBinTree returns a new empty binary tree.
func NewBinTree[T any]() *BinTree[T] {
	return &BinTree[T]{
		root: NewElement[T](),
	}
}

// IsEmpty returns true if the tree is empty.
func (t *BinTree[T]) IsEmpty() bool {
	return t.root.IsEmpty()
}

// MermaidString returns a Mermaid representation of the tree.
func (t *BinTree[T]) MermaidString() string {
	mermaidlines := strings.Join(t.root.MermaidLines(), "\n")
	return "graph TD" + "\n" + mermaidlines
}

// SetValue expects a path and a value.
// The path is a string of L and R characters, indicating the path to the element.
// For example, "LLR" means the right child of the left child of the left child.
// Non-existing elements are created as needed.
func (t *BinTree[T]) SetValue(path string, value T) {
	e := t.root
	for _, c := range path {
		if c == 'L' {
			e = e.addLeft()
		} else if c == 'R' {
			e = e.addRight()
		} else {
			panic("Invalid path character: " + string(c))
		}
	}
	e.SetValue(value)
}

// Value returns the value of the element at the given path.
// Panics if the element is empty or if the path is invalid.
func (t *BinTree[T]) Value(path string) T {
	e := t.root
	for _, c := range path {
		if c == 'L' {
			e = e.left
		} else if c == 'R' {
			e = e.right
		} else {
			panic("Invalid path character: " + string(c))
		}
	}
	return e.Value()
}
