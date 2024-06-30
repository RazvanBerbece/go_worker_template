package unit

import (
	"testing"
)

func TestExample(t *testing.T) {

	cases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{1, 1}, 2},
		{[]int{1, 2}, 3},
	}

	for _, c := range cases {
		if output := c.input[0] + c.input[1]; c.expectedOutput != output {
			t.Errorf("incorrect actual output %d: expected %d", output, c.expectedOutput)
		}
	}

}
