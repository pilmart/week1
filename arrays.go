package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	// Array length declared upfront
	var myArray1 = [3]int{1, 2, 3}
	myArray2 := [5]string{"a", "b", "c", "d", "e"}

	pl(myArray1)
	pl(myArray2)

	var myArrayInferred = [...]string{"x", "y", "z"}
	pl("length of inferred array", len(myArrayInferred))

	myArrayInferred1 := [...]string{"x", "y", "z", "a", "b"}
	pl("length of inferred array", len(myArrayInferred1))

}
