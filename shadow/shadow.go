package main
import "fmt"
var a=20
func main() {
 age:=30
 if age>18{
	a:=47
	fmt.Println("shadow a->",a)
 }
 fmt.Println("Real a->",a)
}
