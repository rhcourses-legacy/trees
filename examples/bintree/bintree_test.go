package bintree

import (
	"strings"
	"testing"

	"github.com/rhcourses/trees/testhelpers"
)

func TestBinTree_NewBinTree(t *testing.T) {
	tree := NewBinTree[string]()
	if !tree.IsEmpty() {
		t.Error("NewBinTree should return an empty tree")
	}
}

func TestBinTree_SetValue(t *testing.T) {
	tree := NewBinTree[string]()
	tree.SetValue("L", "foo")

	if tree.IsEmpty() {
		t.Error("Tree should not be empty")
	}
	valueL := tree.Value("L")
	testhelpers.AssertEqual(t, "foo", valueL)

	tree.SetValue("R", "bar")
	valueR := tree.Value("R")
	testhelpers.AssertEqual(t, "bar", valueR)
}

func TestBinTree_MermaidString(t *testing.T) {
	tree := NewBinTree[string]()
	tree.SetValue("", "root")
	tree.SetValue("L", "L")
	tree.SetValue("R", "R")
	tree.SetValue("LL", "LL")
	tree.SetValue("LR", "LR")
	tree.SetValue("RL", "RL")
	tree.SetValue("RR", "RR")

	expectedstrings := []string{
		"graph TD",
		"root --> L",
		"root --> R",
		"L --> LL",
		"L --> LR",
		"R --> RL",
		"R --> RR",
	}
	expected := strings.Join(expectedstrings, "\n")
	actual := tree.MermaidString()

	testhelpers.AssertEqual(t, expected, actual)
}
