## Routine
Is a light weight thread

Usage
```go
func main(){
    // Creates a new routine if we use go keyword to any functin keyword
    go printMessage("Go is great")
}
```
By default a main go routine is created

> A GO app ends when the main go routine is ended

## GoRoutines and Channels
- A Go routine is the GO way of using threads
- We open a Go Routine just by invoking any function with a `go` prefix
- `go functionCall()`

- Go routines can communites through channels, a special type of variable
- A channel contains a value of any kind
- A routine define a value for a channel and other routine can wait for that value
- Channels can be buffered or not

```go
var m1 chan string

m2 := make(chan string)

m2 <- "hello"

message := <- m2
```

