package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type do string

func (d do) M() {
	fmt.Println(d)
}

func main() {
	do("233, 666").M()

	var d interface{} = do("233")
	var s interface{} = string("233")

	fmt.Println(d == s)
	fmt.Printf("%T %v \n", s, s)
	fmt.Printf("%T %v \n", d, d)

	fmt.Println(reflect.TypeOf(d))

	a1 := []int{1, 2, 3}
	// a2 := []int{1, 2, 3}
	fmt.Println(a1 == nil)

	var a3, a4 [3]int
	for i := 0; i < 3; i++ {
		a3[i] = i
		a4[i] = i
	}
	fmt.Println(a3 == a4)

	// int size

	fmt.Println(strconv.IntSize)

	var i = -1
	fmt.Println(i)

}
