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

func TestTotalFloat64(t *testing.T) {
	type testCase struct {
		Name string

		Prices []float64

		ExpectedFloat64 float64
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualFloat64 := TotalFloat64(tc.Prices)

			assert.Equal(t, tc.ExpectedFloat64, actualFloat64)
		})
	}

	validate(t, &testCase{
		Name:            "TestTotalFloat64 with positive numbers",
		Prices:          []float64{10.99, 5.99, 3.99, 7.99},
		ExpectedFloat64: 28.96,
	})
}

func TestAverage(t *testing.T) {
	type testCase struct {
		Name string

		Numbers []int

		ExpectedFloat64 float64
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualFloat64 := Average(tc.Numbers)

			assert.Equal(t, tc.ExpectedFloat64, actualFloat64)
		})
	}

	validate(t, &testCase{
		Name:            "TestAverage with positive numbers",
		Numbers:         []int{1, 2, 3, 4, 5},
		ExpectedFloat64: 3,
	})
}

func TestAverageFloat64(t *testing.T) {
	type testCase struct {
		Name string

		Prices []float64

		ExpectedFloat64 float64
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualFloat64 := AverageFloat64(tc.Prices)

			assert.Equal(t, tc.ExpectedFloat64, actualFloat64)
		})
	}

	validate(t, &testCase{
		Name:            "TestAverageFloat64 with positive numbers",
		Prices:          []float64{10.99, 5.99, 3.99, 7.99},
		ExpectedFloat64: 7.24,
	})
}
