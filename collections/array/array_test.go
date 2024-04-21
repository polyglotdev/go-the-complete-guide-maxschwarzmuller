package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotal(t *testing.T) {
	type testCase struct {
		Name string

		Numbers []int

		ExpectedInt int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualInt := Total(tc.Numbers)

			assert.Equal(t, tc.ExpectedInt, actualInt)
		})
	}

	validate(t, &testCase{
		Name:        "TestTotal with positive numbers",
		Numbers:     []int{1, 2, 3, 4, 5},
		ExpectedInt: 15,
	})
}
