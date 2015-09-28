package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		in1, in2, want Matrix
	}{
		{
			Matrix{{Element{48}, Element{39}}, {Element{37}, Element{95}}},
			Matrix{{Element{65}, Element{82}}, {Element{29}, Element{87}}},
			Matrix{{Element{4251}, Element{7329}}, {Element{5160}, Element{11299}}},
		},
		{
			Matrix{{Element{23}, Element{32}}, {Element{43}, Element{23}}},
			Matrix{{Element{89}, Element{84}}, {Element{76}, Element{64}}},
			Matrix{{Element{4479}, Element{3980}}, {Element{5575}, Element{5084}}},
		},
	}

	for _, c := range cases {

		matChannel := make(chan Matrix)
		go multiply(c.in1, c.in2, matChannel)

		got := <-matChannel
		if !got.IsEqual(c.want) {
			t.Errorf("Multiply(%v, %v) == %v, want %v", c.in1, c.in2, got, c.want)
		}
	}
}
