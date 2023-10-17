package main

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "USA"
}

func init() {
	fmt.Println("Test")
}

// You can have multiple init() in the same file

func main(){
	fmt.Print("Main")
}