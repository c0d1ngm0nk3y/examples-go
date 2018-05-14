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
