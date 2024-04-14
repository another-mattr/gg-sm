package types

import "reflect"

func GetZeroValue[T any]() any {
	return reflect.Zero(reflect.TypeOf((*T)(nil)).Elem()).Interface()
}