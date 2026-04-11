package main
import "fmt"
var a=23
func main(){
	fmt.Println("Main Function Execute after init function")
	a:=40
	fmt.Println("a value after init call->",a)
}
/*
🔥 Key Characteristics
✔ No parameters
✔ No return value
✔ Automatically executed
✔ Cannot be called explicitly
✔ Used for initialization tasks

*/
func init(){
	fmt.Println("Init Function Execute First")
	fmt.Println("a value before main call->",a)
}