package main
import "fmt"
func printArr(number *[3]int){
	fmt.Println(number)
}
func main(){
	x:=10
    fmt.Println("X->",x)
	p:=&x// &-> address of
	*p=30 // *->value at address
	fmt.Println("X->",x)

	arr:=[3] int {1,2,3,}
	printArr(&arr)

}