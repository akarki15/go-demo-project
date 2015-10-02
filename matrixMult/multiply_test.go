package main

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		in1, in2, want Matrix
	}{
		{
			Matrix{Val: [][]Element{{Element{48}, Element{39}}, {Element{37}, Element{95}}}},
			Matrix{Val: [][]Element{{Element{65}, Element{82}}, {Element{29}, Element{87}}}},
			Matrix{Val: [][]Element{{Element{4251}, Element{7329}}, {Element{5160}, Element{11299}}}},
		},
		{
			Matrix{Val: [][]Element{{Element{23}, Element{32}}, {Element{43}, Element{23}}}},
			Matrix{Val: [][]Element{{Element{89}, Element{84}}, {Element{76}, Element{64}}}},
			Matrix{Val: [][]Element{{Element{4479}, Element{3980}}, {Element{5575}, Element{5084}}}},
		},
	}

	for _, c := range cases {

		matChannel := make(chan Matrix)
		go multiply(c.in1, c.in2, matChannel, 0, 0, 0, 0)

		got := <-matChannel
		ans, err := got.IsEqual(c.want)
		if err != nil {
			fmt.Println(err)
		} else if !ans {
			t.Errorf("Multiply(%v, %v) == %v, want %v", c.in1, c.in2, got, c.want)
		}

	}
}
