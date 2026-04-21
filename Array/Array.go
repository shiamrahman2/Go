package main
import "fmt"

func main(){
   var arr[2] int
    
   arr[0]=1
   arr[1]=3
   fmt.Println(arr)
   // shorten array defined
   arr1:=[3] int{
	3,4,5,
   }
   fmt.Println(arr1)
}