package main

import (
	"fmt"
	"strconv"
)

type Numero int

func (n Numero) String() string {
	return strconv.Itoa(int(n))
}

func (n Numero) Format(f fmt.State, verb rune) {
	val := n.String()
	if verb == 81 {
		val = "**" + val + "**"
	}
	fmt.Fprint(f, val)
}

func mainformatter() {
	a := Numero(1)
	fmt.Printf("number is: %Q\n", a)

}
