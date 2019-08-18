package main

import "fmt"

var m = make(map[string]int)

func main() {
	list := []string{"123", "234", "345", "456"}
	//fmt.Println(k(list))

	Add(list)
	fmt.Println(Count(list))
	Add(list)
	fmt.Println(Count(list))
	Add(list)
	fmt.Println(Count(list))
	Add(list)
	fmt.Println(Count(list))

}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string)  {
	m[k(list)]++
	fmt.Println(k(list), m[k(list)])
}

func Count(list []string) int {
	return m[k(list)]
}
