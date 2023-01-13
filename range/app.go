package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int;
	Y int;
}

func main() {
	// arr := []string{"a", "b", "c"}
	// for i, el := range arr {
	// 	fmt.Println(i)
	// 	fmt.Println(el)
	// }

	// for _, el := range arr {
	// 	fmt.Println(el)
	// }

	pointsMap := map[string]int{
		"x": 123,
		"y": 234,
	}

	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)

	fmt.Println(p1)

	for k, v := range pointsMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}