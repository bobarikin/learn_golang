package main

import "fmt"

type Point struct {
	X, Y int
}

func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func main() {
	someCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	result := doSomething(someCallback, "plus")
	fmt.Println(result)

	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(movePoint(p, 1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)
}