package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2(t *testing.T) {
	assert.Equal(t, true, IsPrime(2))
}

func Test1(t *testing.T) {
	assert.Equal(t, false, IsPrime(1))
}

func Test0(t *testing.T) {
	assert.Equal(t, false, IsPrime(0))
}

func Test4(t *testing.T) {
	assert.Equal(t, false, IsPrime(4))
}

func Test3(t *testing.T) {
	assert.Equal(t, true, IsPrime(3))
}

func Test9(t *testing.T) {
	assert.Equal(t, false, IsPrime(9))
}

func Test113(t *testing.T) {
	assert.Equal(t, true, IsPrime(113))
}

func Test169(t *testing.T) {
	assert.Equal(t, false, IsPrime(169))
}
