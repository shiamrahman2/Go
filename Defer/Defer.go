package main

import "fmt"
/*
  Rule:

return value → copy first, then defer
named return → defer can change final result

*/
func calculator()(result int){
    result=0;
	fmt.Println("First-",result)

	show:=func(){
		result=result+10
		fmt.Println("Defer-",result)
	}

	defer show()

	result=5

	fmt.Println("Second",result)

	return
}

func calc() int{
    result:=0;
	fmt.Println("First-",result)

	show:=func(){
		result=result+10
		fmt.Println("Defer-",result)
	}

	defer show()

	result=5

	fmt.Println("Second",result)

	return result
}

func main(){
	/*
	i:=0

	fmt.Println("First-",i)

	defer fmt.Println("second-",i)

	i++

	fmt.Println("Third-",i)

	defer fmt.Println("Fourth-",i)
	*/
	a:=calculator()
	fmt.Println(a)
	b:=calc()
	fmt.Println(b)

}