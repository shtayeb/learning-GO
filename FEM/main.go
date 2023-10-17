package main

import "shtb.dev/go/fem/data"

func main() {
	sh := data.Instructor{Id: 1, LastName: "Tayeb"}
	sh.FirstName = "Shahryar"

	// john := data.NewInstructor("John", "Doe")

	goCourse := data.Course{Id: 1, Name: "Go", Slug: "go", Legacy: false, Instructor: sh}

	print(goCourse.String())

	// print(sh.Print())
	// print(john)
}
