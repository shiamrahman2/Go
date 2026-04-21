package main
import "fmt"

type User struct{
	Name string
	Age int
}

// Receiver Function

func(usr User) printDetail(){
   fmt.Println("Name=",usr.Name)
   fmt.Println("Age=",usr.Age)
}

func (usr User) printUserDetail(a int){
	fmt.Println("User Age After",a,"Year late=",usr.Age+a)
}
func main(){
    var user1 User

	user1=User{
		Name:"Shiam",
		Age:22,
	}
    user1.printDetail()
	user1.printUserDetail(10)
	user2:=User{
		Name:"Sami",
		Age:24,
	}
	user2.printDetail()
	user2.printUserDetail(15)
}