package main

import "fmt"
const p=15
 
var a=10
/*
Closure in Golang - Detailed Explanation

This program demonstrates how closures work in Go and how they interact
with variables from different scopes.

1. Global and Constant Variables:
   - 'a' is a global variable (value = 10)
   - 'p' is a constant (value = 15)
   These can be accessed inside any function, including closures.

2. outer() Function:
   - This function defines two local variables:
        money := 100
        age   := 30
   - It prints the value of 'age'
   - Then it defines an anonymous function 'show'

3. Closure (Anonymous Function):
   - The function 'show' is a closure because it captures and uses
     the variable 'money' from its outer scope.
   - Inside 'show':
        money = money + a + p
     It updates 'money' using:
        local variable (money)
        global variable (a)
        constant (p)
   - It prints the updated value of 'money'

4. Returning the Closure:
   - outer() returns the 'show' function
   - Even after outer() finishes execution, the returned function
     still has access to 'money'
   - This is possible because closures "remember" their environment

5. Behavior in main():
   - incre1 := outer()
        → creates a new closure with its own 'money'
   - Calling incre1() multiple times:
        → modifies and retains the same 'money' value

   - incre2 := outer()
        → creates a completely new closure
        → has a separate 'money' variable (independent of incre1)

6. Key Concept:
   - Each call to outer() creates a NEW closure with its own state
   - Closures capture variables by reference, not by value
   - This allows them to maintain state between function calls

7. init() Function:
   - Runs automatically before main()
   - Used here to print a starting message

Summary:
A closure in Go is a function that captures variables from its surrounding
scope and retains access to them even after the outer function has returned.
Each closure instance maintains its own independent state.
*/
func outer() func(){
	money:=100
	age:=30
	fmt.Println("Age->",age)
	show:=func(){
		money=money+a+p
		fmt.Println(money)
	}
	return show
}
func main(){
    incre1:=outer()
	incre1()
	incre1()
	incre2:=outer()
	incre2()
	incre2()
}
func init(){
	fmt.Println("=====Bank=====")
}