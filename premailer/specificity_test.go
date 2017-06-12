package premailer

import (
	"sort"
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

func TestSpecificitySort(t *testing.T) {
	undertest := make([]*styleRule, 4)
	for i := 0; i < 4; i++ {
		undertest[i] = &styleRule{}
	}
	specificity0 := makeSpecificity(1, 2, 100, "table")
	undertest[0].specificity = specificity0
	specificity1 := makeSpecificity(1, 2, 102, "table.red")
	undertest[1].specificity = specificity1
	specificity2 := makeSpecificity(1, 3, 103, "span[lang~=\"en-us\"]")
	undertest[2].specificity = specificity2
	specificity3 := makeSpecificity(0, 3, 104, "#example")
	undertest[3].specificity = specificity3

	// expected order
	/*
		expected3 := []int{0, 1, 0, 0, 0, 3, 104}
		expected0 := []int{1, 0, 0, 0, 1, 2, 100}
		expected2 := []int{1, 0, 0, 1, 1, 3, 103}
		expected1 := []int{1, 0, 1, 0, 1, 2, 102}
	*/
	sort.Sort(bySpecificity(undertest))

	assert.Equal(t, specificity3, undertest[0].specificity)
	assert.Equal(t, specificity0, undertest[1].specificity)
	assert.Equal(t, specificity2, undertest[2].specificity)
	assert.Equal(t, specificity1, undertest[3].specificity)
}

func TestSpecificitySortRuleSetIndex(t *testing.T) {
	undertest := make([]*styleRule, 2)
	for i := 0; i < 2; i++ {
		undertest[i] = &styleRule{}
	}
	specificity0 := makeSpecificity(1, 2, 102, "table")
	undertest[0].specificity = specificity0
	specificity1 := makeSpecificity(1, 1, 102, "table")
	undertest[1].specificity = specificity1

	sort.Sort(bySpecificity(undertest))

	assert.Equal(t, specificity1, undertest[0].specificity)
	assert.Equal(t, specificity0, undertest[1].specificity)
}

func TestSpecificitySortRuleIndex(t *testing.T) {
	undertest := make([]*styleRule, 2)
	for i := 0; i < 2; i++ {
		undertest[i] = &styleRule{}
	}
	specificity0 := makeSpecificity(1, 1, 102, "table")
	undertest[0].specificity = specificity0
	specificity1 := makeSpecificity(1, 1, 100, "table")
	undertest[1].specificity = specificity1

	sort.Sort(bySpecificity(undertest))

	assert.Equal(t, specificity1, undertest[0].specificity)
	assert.Equal(t, specificity0, undertest[1].specificity)
}
