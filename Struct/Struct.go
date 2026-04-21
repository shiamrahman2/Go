package main
import "fmt"

type User struct{
	Name string
	Age int
}

func main(){
    var user1 User

	user1=User{
		Name:"Shiam",
		Age:22,
	}
    fmt.Println("Name->",user1.Name)
	fmt.Println("Age->",user1.Age)
	user2:=User{
		Name:"Sami",
		Age:24,
	}
	fmt.Println("Name->",user2.Name)
	fmt.Println("Age->",user2.Age)
}
/*
This program demonstrates how to use struct in Go.

1. A struct named 'User' is defined with two fields:
   - Name (string)
   - Age (int)

2. In the main function:
   - A variable 'user1' is declared and then initialized using struct literal.
   - Another variable 'user2' is directly initialized using shorthand syntax.

3. The program prints the Name and Age of both users.

This example shows:
- Struct declaration
- Struct initialization (two ways)
- Accessing struct fields
*/