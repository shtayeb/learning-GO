## Basic Rules
- We use .go files
- Code Blocks in {}
- No styling freedom
- We do have semi-colon to separate sentences
- Case-sensitive
- Strongly typed
- NOT and object-oriented language
- No classes, no exceptions
- We have on file acting as the entry point with a main function
- A folder is a package
- Packages can have simple names or URLS(github.com/fem/my-library)
- Within one go file, we can have
    - Variables
    - Functions
    - Type declarations
    - Method declarations

## Modules and CLI
- A module is a group of packages
- It's our project
- It contains a go.mod file with configuration and metadata
- CLI manipulates the module
    - go mod init
    - go build
    - go run
    - go test
    - go get

## Create a GO module
```shell
go mod init shtb.dev/go/io
```
This will create a `go.mod` file inside the folder with bellow contents
```go
module shtb.dev/go/io

go 1.21.3
```
To run the module we run
```
go run .
```
It will scan the folder and look for a file `main.go` and a function named `main`

## Workspaces and CLI
- A workspace is a new kind of multi-module app concept from 1.18
- It contains a go.work file with configuration and metadata including whick module to use

- CLI manipulates the workspace
    - `go work init`

## Defining variables
var [VARIABLE NAME] [TYPE]

- Data types goes after indetifier
- Variables have `nil` by default
- Constants can only be bool,string or numbers
- Strings uses double quotes

```go
var x int
var name string
var price = 2 // type is taken from value

const y = 2

var z int = 2

var text string
text = "Hello!"
```

### Variable initialization shortcut
It only works inside functions.
Type is taken from the value.
```go
otherText := "Bye!"
```

### Variable scopes
- Global
- Function
- Block

### Immutable and Constant
We can set value of an immutable in runtime.
Value of contants are set in compilation time.

## Built-in Data Types
- string
- integer values: int, int8, int16, int32, int64, uint, uint8, uint16,uint32,uint64
- floating point values: float32, float64
- bool

### Boolean operators: 
==, !=, <, >, >=,<=,&&,||,!

> We can work with pointers!

## Package
- A package is a group of files in the same folder.
- We can define the name at the top.
- We can import other packages.

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello Go!")
}
```

> Go compiler only sees package not files

> If the name starts with uppercase it is public if it is lowercase it is private to that package
> PrintData() - public - will be exported from package
> printData() - private


## Visibility

### What we write in a module
- if it's camelCase, it's private
- if it's TitleCase, it's public and exported

### Variables and lambda functions can be:
- module scoped
- function scoped
- block scoped

```go
package main

import "fmt"

func notExported(){

}

func Exported(){

}
```

## Numbers
```go
id := 4
price := 45.4

priceAsInt := int(price)
idAsFloat := float32(id)
```

## Strings
Multiline strings by default
```go
str1 := "This is just a text"

str2 := "Every text 
is multiline 
by default
"
```

## Collections
- Arrays: fixed length - `[5]int`
- Slices: similar to a dynamic length array, but they are actually chucks of arrays - `[]int`

- Maps: key/value dictinaries 
    var VAR_NAME map[TYPE_0F_KEY]TYPE_OF_VALUE
    ```go
    var Codes map[int]string

    // To get length
    qty := len(Codes)
    ```


> You can have more that on function named init()

- Collections are not objects, so we use global functions to work with them,
such as `len()` or `cap()`


```go
// Arrays
var Countries [3]string
Countries[0] = "Panama"

Prices := [2]int {100, 150} // Like a constructor

// Slices
names := []string {"Mary","John"}
names := append(names,"carol")
println(len(names))
```

## Functions
- Similar to other languages
- It can receive arguments
- Arguments can have default values
- The last argument can be of variable length
### The tricky part
    - Functions can return more than one value (at once)
    - Functions can return labeled variables
    - Functions receive arguments always by value

```go
func add(a int, b int) int {
    return a+b;
}

func addAndSubstract(a int,b int)(int,int){
    return a+b,a-b
}

addition, substraction := addAndSubstract(4,2)
_, substraction := addAndSubstract(4,2)
addition, _ := addAndSubstract(4,2)
```

### Functions receiving references
> One package one main()

```go
func birthday(age *int){
    age = *age + 1
}

func main(){
    age := 22
    birthday(&age)
    fmt.Prinln(age)
}
````

## Functin Curiosity
- `panic("Message!")` // It will close your app
- `defer fmt.Println('bye')` // It will delay execution to the end. It uses stack data structure
- `init()` //

## Errors Design Pattern

We don't have exceptions in GO, this is the typical design pattern when an action may trigger an error.

## Control Structure
- if - else
    ```go
    if user != nil{

    } else {

    }

    if message := "hello"; user != nil{
        // message is only available inside the if and else
    } else {

    }
    ```
- switch(reloaded!)
    This is a simple switch operation. No break is needed; you can fallthrough to the next case, though
    ```go
    switch day{
        case "Monday":
            fmt.Println("It is Monday")
        case "Saturday":
            fallthrough
        case "Sunday":
            fmt.Println("This is sunday")
        default:
            fmt.Println("It is another working day")
    }
    ```

    Switch with condition
    It can replace large ifs
    ```go
    switch {
        case user == nil:

        case user.active == false:

        case user.deleted == true:

        default:
    }
    ```
- for
    ```go
    for i:=0; i<len(collection); i++{

    }

    // For range, similar to "for in" in JS
    for index:=range collection {

    }

    // For range, similar to "foreach"
    for key, value := range map {

    }
    ```

    We can emulate a while by using a boolean expression
    ```go
    endOfGame := false
    for endOfGame {
        // process Game Loop
    }

    for count < 10 {
        count += 1;
    }

    for {
        // Infinite loop
    }
    ```


- There is no while or do-while
- No parenthesis are needed for boolean conditions or values
- Only one type of equality operators == 
- Other operators != < > <= >=

