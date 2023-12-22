package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	var n int64;
	fmt.Println("Enter a decimal number: ")
	fmt.Scan(&n)
	binary := strconv.FormatInt(n, 2);
	
	fmt.Println(binary)
	
}