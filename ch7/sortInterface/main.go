package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}

func main() {
	name := []string{"823", "234", "345", "sda", "ass"}
	tmp := StringSlice(name)
	sort.Sort(tmp)
	fmt.Println(tmp)
}
