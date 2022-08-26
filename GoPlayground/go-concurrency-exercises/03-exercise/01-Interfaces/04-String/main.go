package main

import "fmt"

type user struct {
	name  string
	email string
}

//Implement custom formating for user struct values. Using method defined in Stringer interface.
func (u user)String() string{
	return fmt.Sprintf("%s <%s>",u.name,u.email)
}
func main() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	fmt.Println(u)
}
//Output: Without Stringer interface
//{John Doe johndoe@example.com}

//Output after implementing Stringer interface
//John Doe <johndoe@example.com>