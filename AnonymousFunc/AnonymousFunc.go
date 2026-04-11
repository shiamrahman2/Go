package main
import "fmt"

func main(){
	// anonymous function assign in variable add
	add:=func(x int ,y int )(int ){
		z:=x+y
		return z
	}
	fmt.Println(add(5,6))
	// Immediately Invoked Function Expression
	func(){
		fmt.Println("I'am Invoke after defined so i am IIFE")
	}()
}