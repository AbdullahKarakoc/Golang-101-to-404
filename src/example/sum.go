package main

import "fmt"

func sum(a, b any) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

}

func mainSum() {
	sum(1, 2)
	sum(1, 3.4)
}
