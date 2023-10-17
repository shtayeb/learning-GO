## Type Definitions
- New Types
- Aliases

```go
    // Type Alias
    type distance = float32

    // New type based on 
    type distance float32

    // types have a constructor and conversion functions
    d := distance(34.5)

```

## Methods for a type
```go
type distance float32 // miles
type distanceKm float64

// Method for a type
func (km distanceKm) ToMiles() distance {
	return distance(km / 1.6)
}

func (miles distance) ToKm() distanceKm {
	return distanceKm(1.6 * miles)
}

func test() {
	d := distance(4.5)
	// using the type method
	km := d.ToKm()

	print(km.ToMiles())
}
```


## Complex types for definitions
- Structures
    - They kind of replace the class idea
    - It's a data type with strongly typed property
    - They have default contructor
    - You can add methods to it

- Interfaces
    - A definition of methods
    - You emulate polimorphism from OOP
    - Implicit implementation
    - We can embed interfaces in other interfaces


## Type - Structs
If you want to export struct remember to user TitleCase.
You will find two pre-built constructors, with and without name

```go
type User struct{
    id int
    name string
}

func main(){
    var u2 User
    u1 = User {id:1,name:"Frontend Masters"}
    u2 := User {2,"frontend Masters"}
}
```

## Type - Struct with Methods
Structs are functions attached to a type declared outside of the structure.

```go
func (u User) PrettyPrint() string {
    return string(u.id) + ": " + u.name
}

func main(){
    u2 := User {2,"Frontend Masters"}
    msg := u2.PrettyPrint()
}
```
