package main

import (
	"fmt" 
	"strconv"
)

func add(x, y int, s string) string{
	return strconv.Itoa(x + y) + s
}
func main(){
	fmt.Println(add(3, 6, "ladoo khaunga"))
}