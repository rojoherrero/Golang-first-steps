package models

// Person defines a real person data
type Person struct {
	Name     string
	Dob      string
	Gender   string
	Single   bool
	Employed bool
}

// Student defines a student
type Student struct {
	Person
	Speciality string
}
