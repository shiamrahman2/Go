package main
import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}



func main(){
   // slice create
     /*
	   1->From an existing Array
	 */
   arr:=[6]string{"This","is","a","Go","Interview","Question"}

   s:=arr[1:4]
   
   /*
     slice have three Property
	  1->Ptr
	  2->length
	  3->Capacity
 
   */
   fmt.Println(arr)
   fmt.Println(s)
   fmt.Println(len(s))
   fmt.Println(cap(s))

   /*
    2-> from a slice
   */
   s1:=s[2:3]
   fmt.Println(s1)
   fmt.Println(len(s1))
   fmt.Println(cap(s1))

   /*
    3-> slice literals
   */

   s2:=[]int{1,2,3}
    fmt.Println(s2)
   fmt.Println(len(s2))
   fmt.Println(cap(s2))

  /*
    4-> using make building function
	Can be used liked array
  */

  s3:=make([]int,4)// (type,size)

    s3[0]=5
	s3[1]=4
    fmt.Println(s3)
   fmt.Println(len(s3))
   fmt.Println(cap(s3))

   /*
      5->Using make function
   */
   s4:=make([]int,3,5)//(type,size,capacity)
   fmt.Println(s4)
   fmt.Println(len(s4))
   fmt.Println(cap(s4))

     s4=s4[:5]// extend length max size of capacity
    fmt.Println(s4)
   fmt.Println(len(s4))
   fmt.Println(cap(s4))
   /*
     6-> empty or nil slice
   */
    var s5[] int

	  fmt.Println(s5)
   fmt.Println(len(s5))
   fmt.Println(cap(s5))

   s5=append(s5,1,3,4,5)
   fmt.Println(s5)
   fmt.Println(len(s5))
   fmt.Println(cap(s5))

   var x[] int

   x=append(x, 1)
   x=append(x, 2)
   x=append(x, 3)
   y:=x
   x=append(x, 4)

   fmt.Println(x)
   fmt.Println(y)

   y=append(y, 5)

   fmt.Println(x)
   fmt.Println(y)

   x[0]=10
    
   fmt.Println(x)
   fmt.Println(y)

   p := []int{1, 2, 3, 4, 5}

	p = append(p, 6)
	p = append(p, 7)

	q := p[4:]

	r := changeSlice(q)

	fmt.Println(p)
	fmt.Println(r)

}