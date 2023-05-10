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

func TestElement_SetValue(t *testing.T) {
	e := NewElement[string]()
	e.SetValue("foo")

	if e.IsEmpty() {
		t.Error("Element should not be empty")
	}
	if e.Value() != "foo" {
		t.Error("Element should have value 'foo', but got", e.Value())
	}
}
