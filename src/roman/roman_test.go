package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArabicI(t *testing.T) {
	assert.Equal(t, 1, ToArabic("I"))
}

func TestToArabicV(t *testing.T) {
	assert.Equal(t, 5, ToArabic("V"))
}

func TestToArabicX(t *testing.T) {
	assert.Equal(t, 10, ToArabic("X"))
}

func TestToArabicL(t *testing.T) {
	assert.Equal(t, 50, ToArabic("L"))
}

func TestToArabicC(t *testing.T) {
	assert.Equal(t, 100, ToArabic("C"))
}

func TestToArabicD(t *testing.T) {
	assert.Equal(t, 500, ToArabic("D"))
}

func TestToArabicM(t *testing.T) {
	assert.Equal(t, 1000, ToArabic("M"))
}

func TestToArabicML(t *testing.T) {
	assert.Equal(t, 1050, ToArabic("ML"))
}

func TestToArabicIV(t *testing.T) {
	assert.Equal(t, 4, ToArabic("IV"))
}
