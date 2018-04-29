package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactors1(t *testing.T) {
	expected := []int(nil)
	assert.Equal(t, expected, GetPrimeFactors(1))
}

func TestFactors2(t *testing.T) {
	expected := []int{2}
	assert.Equal(t, expected, GetPrimeFactors(2))
}

func TestFactors3(t *testing.T) {
	expected := []int{3}
	assert.Equal(t, expected, GetPrimeFactors(3))
}
