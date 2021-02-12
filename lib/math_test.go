package lib

import "testing"

// https://golang.org/pkg/testing/
// testing framework Ginkgo、Gomega

var Debug bool = true

func TestAverageSuccess(t *testing.T) {
	if Debug {
		t.Skip("Skip Reason")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}

func TestAverageSkip(t *testing.T) {
	t.Skip("Skip Reason")
	v := Average([]int{1, 2, 3, 4, 5, 6, 7})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
