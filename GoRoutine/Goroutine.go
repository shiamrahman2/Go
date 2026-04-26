package main
import(
	"fmt"
	"time"
)
func printHello(num int){
	fmt.Println("Hello -",num)
}
func main(){
	fmt.Println("Hello Shiam")
    
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)
     
	fmt.Println("Hi Myself Shiam")

	time.Sleep(5*time.Second)

	fmt.Println("This Is Shiam")



}
