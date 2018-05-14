package person

import (
	"fmt"
)

//Person has a name
type Person struct {
	firstname string
	lastname  string
}

func (p Person) sayHello() string {
	s := fmt.Sprintf("I am %s %s", p.firstname, p.lastname)
	return s
}

func (p Person) getName() (string, string) {
	return p.firstname, p.lastname
}

//FormalPerson is also a Person
type FormalPerson struct {
	Person
}

func (fp FormalPerson) sayHello() string {
	s := fmt.Sprintf("I am %s, %s", fp.lastname, fp.firstname)
	return s
}

//ConfusedPerson is also a Person AND FormalPerson
type ConfusedPerson struct {
	Person
	FormalPerson
}
