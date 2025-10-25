package util

import "testing"


func TestArraysEqual(t *testing.T) {

	//Test: Simple test to make sure simple arrays that equal do equal

	equal := ArraysEqual([]string{
		"a", "b", "c",
	}, []string{
		"a", "b", "c",
	})

	if !equal { 
		t.Error(`["a", "b", "c"] should match ["a", "b", "c"]`)
	}

	//Test: Make sure length is taken into consideration

	equal = ArraysEqual([]int{
		1, 2, 3,
	}, []int{
		1, 2, 3, 4,
	})

	if equal { 
		t.Error(`[1, 2, 3] should not match [1, 2, 3, 4]`)
	}

	equal = ArraysEqual([]int{
		1, 2, 3, 4,
	}, []int{
		1, 2, 3,
	})

	if equal { 
		t.Error(`[1, 2, 3, 4] should not match [1, 2, 3]`)
	}

	//Test: If value is somehow not taken into consideration somehow I dont even know...
	equal = ArraysEqual([]int{
		1, 2, 3,
	}, []int{
		4, 5, 6,
	})

	if equal { 
		t.Error(`[1, 2, 3] should not match [4, 5, 6]`)
	}
}
