package main

import (
	"reflect"
	"testing"
)

func TestReport(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]int
		expected []Page
	}{
		{
			name: "Test an unordered map without duplicates",
			input: map[string]int{
				"https://example.com/page1": 5,
				"https://example.com/page3": 3,
				"https://example.com/page2": 4,
				"https://example.com/page7": 1,
			},
			expected: []Page{
				{
					url:   "https://example.com/page1",
					count: 5,
				},
				{
					url:   "https://example.com/page2",
					count: 4,
				},
				{
					url:   "https://example.com/page3",
					count: 3,
				},
				{
					url:   "https://example.com/page7",
					count: 1,
				},
			},
		},
		{
			name: "Test an unordered map with duplicates",
			input: map[string]int{
				"https://example.com/page1": 5,
				"https://example.com/page3": 5,
				"https://example.com/page2": 4,
				"https://example.com/page7": 1,
			},
			expected: []Page{
				{
					url:   "https://example.com/page1",
					count: 5,
				},
				{
					url:   "https://example.com/page3",
					count: 5,
				},
				{
					url:   "https://example.com/page2",
					count: 4,
				},
				{
					url:   "https://example.com/page7",
					count: 1,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortMap(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Test failed")
			}
		})
	}

}
