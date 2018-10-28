package main

import (
	"fmt"
	"math"
)

func main()  {
	var num float64
	fmt.Print(" Input First float point number: ")
fmt.Scan(&num)
var num2 float64
fmt.Print(" Input Second float point number: ")
fmt.Scan(&num2)
fmt.Println(math.Trunc(num))
fmt.Println(math.Trunc(num2))
}