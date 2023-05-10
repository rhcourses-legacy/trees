package bintree

import "testing"

func TestNewElement_IsEmpty(t *testing.T) {
	e := NewElement[string]()
	if !e.IsEmpty() {
		t.Error("New element should be empty")
	}
}

func TestEmptyElement_EmptyString(t *testing.T) {
	e := NewElement[string]()
	if e.String() != "" {
		t.Error("Empty element should return empty string, but got", e.String())
	}
}
