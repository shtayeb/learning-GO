## Interfaces in GO
Interfaces are implicite in GO.
You dont have to manully implement it.

```go
type expense interface {
    calculate() float64
    cost() float64
}
```


## Type assetions in GO
```go
type shape interface {
    area() float64
}

type circle interface{
    radius() float64
}

// Get a circle from a shape type
c, ok := s.(circle)
```

## Clean Interfaces
- Keep interfaces small
- Interfaces should have no knowledge of satisfying types

```go
type car interface {
    Color() string
    Speed() int

    // isFireTrack() NOT GOOD PATTERN
}

// GOOD
type firetrack interface {
    car
    HoseLength() int
}
```

## Interfaces are not classes
- Interfaces are not classes, they are slimmer
- Interfaces arent heirarchical
- interfaces define function signature but not behaviour


## The Error Interface
```go
type error interface {
    Error() string
}
```


## Formating strings
```go
name := "Shahryar"
fmt.Sprintf("Hello %s, I am %s", "World",name)

// "%f" -> float
// "%.2f" -> float

// "%d" -> int
// "%s" -> string

```

## The Error interface
Custome error types
```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("error here")
}
```

## The Errors Package
```go
var err = errors.New("something went wrong")
```

## Loops in GO
```go
for i:=0; i <0; i++{

}
```
We can omit any section of the for loop. like condition
There is no while loop
```go
for CONDITION {
    // acts as while loop
}
```

## Arrays
Are fixes in sizes
```go
var nums [3]int

nums := [3]int{1,2,3}
```

## Slices
- Dynamic sized
- They are built on top of arrays
- Slice are passed by reference

```go
nums := [5]int{1,2,3,4,5}
numSlice := nums[1:3] // [1,2]
numSlice := nums[:] // [1,2,3,4,5]
```

### make
```go
// make([]T,len,cap)
myslice := make([]int,5,10)

myslice := []string{"I","am","here"}

len(myslice)
cap(myslice)
```

## Variadic function
```go
func sum(nums ...int) int {
    // nums is just a slice
    for i := 0; i < len(nums); i++ {

    }
}

sum(1,2,3,4)
```

## Spread operator
```go
printStrings(strings ...string){

}


name := []string{"john","doe"}
printStrings(names...)
```

## Add items to slice
### Append
```go
append(slice,newThing)
append(slice,newThing,thing2)
append(slice,things...)
```

## Slice of Slices
```go
rows := [][]int{}
```

## range
```go
for INDEX,ELEMETN := range SLICE{

}
```

## map
- Are passes as reference
```go
ages := map[string]int{
    "John":34,
    "shahryar":21,
}
```

### Mutations
```go
// insert an element
m[key] = elem

// Get an element
elem = m[key]

// Delete an element
delete(map,key)

// check if a key exist
elem, ok := m[key]
```

- Only comparable types can be used as keys in a map (string,int,struct)

# Advanced Functions

## First class

## Higher order function

## Currying -> func that gets a func and return a func

## Defer keyword
Excute some function at the end of a function execution.

## Closures

## Anonumous functions
```go
test := func(){

}
```

# Pointers
```go
x := 2
&x // address 0xc030

// get value from an address
*x // 2
```

### Pointer type
```go
var p *int
```

## value
```go
// will not update the user struct
func (u User) (){
    u.name = "test"
}

// Will update
unc (u *User) (){
    u.name = "test"
}

```


# Local Development
## Package naming
- Packages are directory level
## Modules
- Modules are collection of packages


### Clean packages
- Hide internal logic
= Dont change API
- Dont export functions from the main package
- Packages shouldnt know about dependents


# Channels and Concurency
Will spawn a go routine
```go
go doSomething()
```
- Channel are a way to safely send value to go routines
## Buffered channels

## Select
is like a switch for channels
```go
select {
    case:
    case:
    default:
}
```

## Read-only and write-only channels

## Channel Review
- A send to a nil channel blocks forever
    ```go
    var c chan string // c is nil
    c <- "let's get started" // blocks
    ```
- A receive from a nil chanel blocks forever
    ```go
    var c chan string // c is nil
    fmt.Println(<-c) // blcoks
    ```
- A send to a closed channel panics
    ```go
    var c = make(chan int,100)
    close(c)

    c <- 1 // panic: send on closed channel
    ```
- A receive from a closed channel returns the zero value immediatly
    ```go
    var c = make(chan int,100)
    close(c)
    fmt.Print(<-c) // 0
    ```

## Mutex in GO (Mutual Exclusion)
Mutex allow us to lock access to data. This ensures that we can control which goroutine can access certain data at which time.

```go
func protected(){
    mux.Lock()
    defer mux.Unlock()

    // The rest of the function is protected
}
```

### Maps are not thread-safe

## RW Mutex
`sync.RWMutex`

# Generics

```go
func splitIntSlice(s []int)([]int, []int){
    mid := len(s)/2
     return s[:mid],s[mid:]
}

func splitAnySlice[T any](s []T) ([]T,[]T){
    mid := len(s)/2
    return s[:mid],s[mid:]
}
```

### Why generics?
- Generic reduce repetitive code
- Generics are used more often in libraries and packages


