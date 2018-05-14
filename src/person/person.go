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
