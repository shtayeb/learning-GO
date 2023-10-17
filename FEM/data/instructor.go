package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(name string, lastname string) Instructor {
	return Instructor{FirstName: name,LastName: lastname}
}

func (instructor Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)",instructor.LastName,instructor.FirstName,instructor.Score)
}