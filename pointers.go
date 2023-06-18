package main
import "fmt"

func main(){
	fmt.Println("pointers ka chapter hai")
	a := 5
	var ptr *int = &a
	fmt.Println(a, ptr);
}
