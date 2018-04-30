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

func TestFactors4(t *testing.T) {
	expected := []int{2, 2}
	assert.Equal(t, expected, GetPrimeFactors(4))
}

func TestFactors9(t *testing.T) {
	expected := []int{3, 3}
	assert.Equal(t, expected, GetPrimeFactors(9))
}

func TestFactors420(t *testing.T) {
	expected := []int{2, 2, 3, 5, 7}
	assert.Equal(t, expected, GetPrimeFactors(420))
}
