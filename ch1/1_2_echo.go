package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		s := strconv.Itoa(i) + " " + os.Args[i]
		fmt.Println(s)
	}
}
