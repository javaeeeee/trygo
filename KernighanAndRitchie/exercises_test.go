package KernighanAndRitchie

import (
	"errors"
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
