package main
import "fmt"
/*
processOperation is a Higher Order Function
  because
     1->take a function as input parameter
 other function would be Higher Order if
    2->return a function 
	3->follow both 1 & 2

*/

func processOperation(a int ,b int, op func(x int,y int)){
     op(a,b)
}

func subs(x int ,y int){
	z:=x-y
	fmt.Println(z)
}
/*
   return add function 
*/
func operation() func(x int,y int){
	return add
}

func add(x int,y int){
	z:=x+y
	fmt.Println(z)
}
func main(){
      processOperation(19,5,subs)
	  sum:=operation()
	  sum(4,5)
}