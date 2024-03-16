package main

import "fmt"

// Numeros is a custom type
type Numeros struct {
	nums []int
}

func sum1(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	return sum1(n[1:]...) + n[0]
}

func (n *Numeros) Sum() int {
	if n == nil {
		return 0
	}
	return sum1(n.nums...)
}

func main() {
	n1 := Numeros{nums: []int{1, 2, 3}}
	fmt.Println(n1.Sum()) // 6

	n2 := Numeros{}
	fmt.Println(n2.Sum()) // 0
}
