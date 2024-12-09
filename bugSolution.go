```go
func main() {
    var x int = 10
    var y *int = &x
    fmt.Println(*y)
    *y = 20
    fmt.Println(x)
    y = nil
    if y != nil {
        fmt.Println(*y) 
    } else {
        fmt.Println("y is nil") // Handle the nil pointer
    }
}
```