package main
import "fmt"
// global scope a and b which can access every part of programme
var (
	a=20
	b=30
)
func add (x int,y int){
	// x and y is local scope which can access only in function
	z:=x+y
	fmt.Println(z)
}
func main(){
	p:=23
	q:=34
	add(p,q)
	add(a,b)
}