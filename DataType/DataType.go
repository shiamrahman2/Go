package main
import "fmt"
func main(){
fmt.Printf("%d\n", 10)   // decimal
fmt.Printf("%b\n", 10)   // binary
fmt.Printf("%o\n", 10)   // octal
fmt.Printf("%x\n", 10)   // hex
fmt.Printf("%f\n", 3.1416)     // default
fmt.Printf("%.2f\n", 3.1416)   // 2 decimal
fmt.Printf("%e\n", 3.1416)     // scientific
fmt.Printf("%s\n", "Hello")
fmt.Printf("%q\n", "Hello") // with quotes
fmt.Printf("%t\n", true)
fmt.Printf("%c\n", 'A')
fmt.Printf("%U\n", 'A') // Unicode
fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", "Go")
fmt.Printf("%T\n", 10)       // int
fmt.Printf("%T\n", "hello")  // string
    x := 10
    y := 3.14
    str := "Go"
    flag := true

    fmt.Printf("Int: %d\n", x)
    fmt.Printf("Float: %.2f\n", y)
    fmt.Printf("String: %s\n", str)
    fmt.Printf("Bool: %t\n", flag)
    fmt.Printf("Pointer: %p\n", &x)
    fmt.Printf("Type: %T\n", x)
}