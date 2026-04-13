package main
import "fmt"

func main(){
	/*
	  add can't be invoke first without defined because it's a local scope
	*/
	// add(4,5)
   add:=func (a int,b int){
	c:=a+b
	fmt.Println(c)
   }
   add(5,6)
}
func init(){
	fmt.Println("I'm init function i'm call first")
}