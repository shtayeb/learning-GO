package main

import "fmt"

// global variable
var url = "http://shahryartayeb.com"

func birthday(age *int){
    *age = *age + 1
	*age++
	fmt.Printf("The pointer is %v and the value is %v\n",age,*age)
}

func main(){
    age := 22
    birthday(&age)
    fmt.Println(age)
}
