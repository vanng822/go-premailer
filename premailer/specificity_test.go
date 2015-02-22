package premailer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSpecificitySelectorType(t *testing.T) {
	spec := makeSpecificity(1, 2, 100, "table")
	expected := []int{1, 0, 0, 0, 1, 2, 100}
	assert.Equal(t, expected, spec.importantOrders())
}

func TestSpecificitySelectorClass(t *testing.T) {
	// class
	spec := makeSpecificity(1, 2, 102, "table.red")
	expected := []int{1, 0, 1, 0, 1, 2, 102}
	assert.Equal(t, expected, spec.importantOrders())
}

func TestSpecificitySelectorAttr(t *testing.T) {
	// Attribute
	spec := makeSpecificity(1, 3, 103, "span[lang~=\"en-us\"]")
	expected := []int{1, 0, 0, 1, 1, 3, 103}
	assert.Equal(t, expected, spec.importantOrders())
}


func TestSpecificitySelectorId(t *testing.T) {
	// id 
	spec := makeSpecificity(0, 3, 104, "#example")
	expected := []int{0, 1, 0, 0, 0, 3, 104}
	assert.Equal(t, expected, spec.importantOrders())
}