package main

import "testing"

func testReport(T *testing.T) {
	cases := []struct {
		name     string
		input    map[string]int
		expected []struct{}
	}{
		{
			name: "Test an unordered map without duplicates",
			input: map[string]int{
				"D": 5,
				"C": 3,
				"B": 4,
				"A": 1,
			},
		},
	}

}
