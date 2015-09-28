package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		in1, in2, want Matrix
	}{
		Matrix{{Element{48}, Element{39}}, {Element{37}, Element{95}}},

		Matrix{{Element{65}, Element{82}}, {Element{29}, Element{87}}},

		Matrix{{Element{4251}, Element{7392}}, {Element{5160}, Element{11299}}},
	}

	for _, c := range cases {
		got := multiply(c.in1, c.in2, nil)
		if got != c.want {
			t.Errorf("Multiply(%v, %v) == %v, want %v", c.in1, c.in2, got, c.want)
		}
	}

}
