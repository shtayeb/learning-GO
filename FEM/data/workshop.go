package data

import "time"

type Workshop struct {
	// Embedding - remove the identifier
	Course
	Name string // This one precedes the Name from Course
	Date time.Time
}

func (c Workshop) Signup() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor

	return w
}
