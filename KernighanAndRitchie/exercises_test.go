package KernighanAndRitchie

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestConvertFahrenheitToCelcius(t *testing.T) {
	cases := []struct {
		low, high, step int
		want            []float64
		err             error
	}{
		{32, 50, 9, []float64{0., 5., 10.}, nil},
		{32, 50, 0, nil, errors.New(ErrorStepZero)},
		{50, 32, 9, nil, errors.New(ErrorLowBiggerThanHigh)},
		{32, 52, 9, nil, errors.New(ErrorNotMultiple)},
	}
	for _, c := range cases {
		got, err := ConvertFahrenheitToCelcius(c.low, c.high, c.step)
		if got == nil && !reflect.DeepEqual(err, c.err) {
			t.Error(err.Error)
		} else if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ConvertFahrenheitToCelcius(%d,%d,%d) == %v, want %v",
				c.low, c.high, c.step, got, c.want)
		}
	}
}

func TestPower(t *testing.T) {
	cases := []struct {
		n    float64
		m    int
		want float64
	}{
		{1., 0, 1.},
		{1., 1, 1},
		{2., 2, 4.},
		{2., -1, .5},
		{0., -1, math.Inf(1)},
	}
	for _, c := range cases {
		got := Power(c.n, c.m)
		if got != c.want {
			t.Errorf("Power(%g,%d)==%g, want %g", c.n, c.m, got, c.want)
		}
	}
}

func TestCopyTextFile(t *testing.T) {
}
