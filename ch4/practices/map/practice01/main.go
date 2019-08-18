//map查询键是否存在
package main

import (
	"fmt"
)

func main() {
	var ages = map[string]int{"A":0}

	age, ok := ages["bob"]
	fmt.Println("bob是否存在 :", ok)
	fmt.Println(age)

	x := map[string]int{"A":0}
	y := map[string]int{"B":42}

	fmt.Println(equal(x, y))

	fmt.Println(missingEqual(x, y))

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x{
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func missingEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x{
		if xv != y[k] {
			return false
		}
	}
	return true
}