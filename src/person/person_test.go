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
