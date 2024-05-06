/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

A variable of interface type stores a pair: the concrete value assigned to the variable,
and that value’s type descriptor. To be more precise, the value is the underlying concrete data
item that implements the interface and the type describes the full type of that item.

At the basic level, reflection is just a mechanism to examine the type and value pair
stored inside an interface variable. To get started, there are two types we need to know
about in package reflect: Type and Value. Those two types give access to the contents of
an interface variable, and two simple functions, called

reflect.TypeOf and reflect.ValueOf
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 1
	b := "lorem ipsum"
	c := 'c'
	r(a)
	r(b)
	r(c)
}

func r(i interface{}) {

	//  When we call reflect.TypeOf(x), x is first stored in an empty interface, which is then
	//  passed as the argument; reflect.TypeOf unpacks that empty interface to recover the type information.
	fmt.Println(reflect.TypeOf(i))
}
