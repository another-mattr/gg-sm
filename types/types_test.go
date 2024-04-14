package types_test

import (
	"reflect"
	"testing"

	"github.com/another-mattr/gg-sm/types" // replace with your actual package path
)

type ComplexType struct {
	A int
	B string
	C *ComplexType
	D []int
}

func TestZeroValueOf(t *testing.T) {
	tests := []struct {
		name     string
		want     interface{}
		typeFunc func() interface{}
	}{
		{"int", int(0), types.GetZeroValue[int]},
		{"float64", float64(0), types.GetZeroValue[float64]},
		{"string", "", types.GetZeroValue[string]},
		{"ComplexType", ComplexType{}, types.GetZeroValue[ComplexType]},
		// add more types as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.typeFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
