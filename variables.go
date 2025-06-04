package main

import (
	"fmt"
	"reflect"
	"strings"
)

var pl = fmt.Println

func main () {
	// explicit declaration
	var vName string = "Martin"
	// type inference + multiple intitialisation 
	var v1, v2 = 1.2 , 2.2
	var v3 = "Hello"
	// also legal
	v4 := 33
	// can only be reassigned like so
	v4 = true
	/*
		core types are ....
		float64, int, string, bool, rune ??
	*/
	// note - curly braces !!!
	sb := strings.Builder{}
	sb.WriteString(reflect.TypeOf(vName).String() + " ")
	sb.WriteString(reflect.TypeOf(v1).String() + " ")
	sb.WriteString(reflect.TypeOf(v2).String() + " ")
	sb.WriteString(reflect.TypeOf(v3).String() + " ")
	sb.WriteString(reflect.TypeOf(v4).String())

	pl(sb.String())
	
	
	
	


}
