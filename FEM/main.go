package main

import (
	"fmt"

	"shtb.dev/go/fem/data"
)

func main() {
	sh := data.Instructor{Id: 1, LastName: "Tayeb"}
	sh.FirstName = "Shahryar"

	// john := data.NewInstructor("John", "Doe")

	goCourse := data.Course{Id: 1, Name: "Go", Slug: "go", Legacy: false, Instructor: sh}
	swiftCourse := data.Course{Id: 2, Name: "Swift", Slug: "swift", Legacy: false, Instructor: sh}

	// print(goCourse.String())

	// print(sh.Print())
	// print(john)

	var courses [2]data.Signable

	courses[0] = goCourse
	courses[1] = swiftCourse

	for _, course := range courses {
		fmt.Print(course)
	}

	// It recives anything
	var things []interface{}

	things = append(things, "34")
	things = append(things, 1)
	things = append(things, true)
	things = append(things, swiftCourse)

	// Casted to type Workshop
	things[2].(data.Workshop).Signup()

	print(things)

}
