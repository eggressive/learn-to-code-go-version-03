package main

import "fmt"

const (
	_ = iota // ignore the first value by assigning it to blank identifier
	a = 1 << iota // 1 shifted left by 0
	b			// 1 shifted left by 1
	c		 	// 1 shifted left by 2
	d			// 1 shifted left by 3
	e			// 1 shifted left by 4
	f			// 1 shifted left by 5
	g			// 1 shifted left by 6
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Printf("%d \t %b\n", a, a)
	fmt.Printf("%d \t %b\n", b, b)
	fmt.Printf("%d \t %b\n", c, c)
	fmt.Printf("%d \t %b\n", d, d)
	fmt.Printf("%d \t %b\n", e, e)
	fmt.Printf("%d \t %b\n", f, f)
	fmt.Printf("%d \t %b\n", g, g)
}
