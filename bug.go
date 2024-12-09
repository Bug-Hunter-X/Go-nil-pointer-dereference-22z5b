```go
func main() {
    var x int = 10
    var y *int = &x
    fmt.Println(*y)
    *y = 20
    fmt.Println(x)
    y = nil
    fmt.Println(*y) // This will cause a runtime error
}
```