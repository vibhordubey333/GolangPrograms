package common

import "fmt"

//Student struct represent a student
type Student struct {
	ID                  int
	FirstName, LastName string
}

//FullName return complete name of the student.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

//College struct represnt college.
type College struct {
	database map[int]Student //private
}

//Add Student.
func (c *College) Add(payload Student, reply *Student) error {

	//Checking if student already exist in DB.
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("Student with ID  '%d' already exist", payload.ID)
	}
	// Add student p in the database
	c.database[payload.ID] = payload
	//set reply value.
	*reply = payload

	// return nil error
	return nil
}

func (c *College) Get(payload int, reply *Student) error {

	//Get student with ID p from the database.
	result, ok := c.database[payload]
	if !ok {
		return fmt.Errorf("Student with ID '%d' does not exist", payload)
	}
	*reply = result

	return nil
}

//NewCollege returns a new instance of college.
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
