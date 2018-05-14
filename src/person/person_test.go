package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonSaysHello(t *testing.T) {
	p := Person{"Ralf", "Pannemans"}
	hello := p.sayHello()
	assert.Equal(t, "I am Ralf Pannemans", hello)
}
func TestPersonGetName(t *testing.T) {
	p := Person{"Ralf", "Pannemans"}
	first, last := p.getName()
	assert.Equal(t, "Ralf", first)
	assert.Equal(t, "Pannemans", last)
}

func TestPersonGetNameNoNames(t *testing.T) {
	p := Person{}
	first, last := p.getName()
	assert.Equal(t, "", first)
	assert.Equal(t, "", last)
}

func TestFormalPersonGetNameSimpleConstr(t *testing.T) {
	fp := new(FormalPerson)
	fp.Person = Person{"Ralf", "Pannemans"}
	first, last := fp.getName()
	assert.Equal(t, "Ralf", first)
	assert.Equal(t, "Pannemans", last)
}
func TestFormalPersonGetNameEnhancedConstr(t *testing.T) {
	fp := FormalPerson{Person: Person{"Ralf", "Pannemans"}}
	first, last := fp.getName()
	assert.Equal(t, "Ralf", first)
	assert.Equal(t, "Pannemans", last)
}
func TestFormalPersonSaysHello(t *testing.T) {
	fp := FormalPerson{Person: Person{"Ralf", "Pannemans"}}
	hello := fp.sayHello()
	assert.Equal(t, "I am Pannemans, Ralf", hello)
}

func TestConfusedPersonSaysHello(t *testing.T) {
	p := Person{"Ralf", "Pannemans"}
	fp := FormalPerson{Person: p}
	cp := ConfusedPerson{FormalPerson: fp, Person: p}

	hello := cp.FormalPerson.sayHello()
	assert.Equal(t, "I am Pannemans, Ralf", hello)
}
