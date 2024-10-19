package main

import (
	"fmt"
	"os"
)

func main() {
	file1, file2 := os.Args[1], os.Args[2]
	fmt.Println(file1)
	fmt.Println(file2)
}