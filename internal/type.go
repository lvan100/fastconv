package internal

import "reflect"

func Ptr[T any](t T) *T {
	return &t
}

func TypeFor[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
