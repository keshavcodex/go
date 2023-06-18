package main

import (
	"fmt"
	"strconv"
)
func concetation(x, y int, s string) string{
	return strconv.Itoa(x + y)+ s
}
func main(){
	fmt.Println(concetation(1, 3, ", this is four"))
}