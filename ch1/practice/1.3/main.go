package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now().UnixNano() - start.UnixNano())

	start = time.Now()
	fmt.Println(start)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now().UnixNano() - start.UnixNano())
}
