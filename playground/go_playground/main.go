package main

import (
	"reflect"
)

func main() {
	i := reflect.ValueOf(3)
	i.Set(reflect.ValueOf(&i).Elem())
}
