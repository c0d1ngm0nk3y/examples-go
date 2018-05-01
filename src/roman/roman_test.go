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

func TestFromArabic1(t *testing.T) {
	assert.Equal(t, "I", FromArabic(1))
}

func TestFromArabic5(t *testing.T) {
	assert.Equal(t, "V", FromArabic(5))
}

func TestFromArabic10(t *testing.T) {
	assert.Equal(t, "X", FromArabic(10))
}

func TestFromArabic50(t *testing.T) {
	assert.Equal(t, "L", FromArabic(50))
}

func TestFromArabic100(t *testing.T) {
	assert.Equal(t, "C", FromArabic(100))
}

func TestFromArabic500(t *testing.T) {
	assert.Equal(t, "D", FromArabic(500))
}

func TestFromArabic1000(t *testing.T) {
	assert.Equal(t, "M", FromArabic(1000))
}

func TestFromArabic1050(t *testing.T) {
	assert.Equal(t, "ML", FromArabic(1050))
}

func TestFromArabic4(t *testing.T) {
	assert.Equal(t, "IV", FromArabic(4))
}

func TestFromArabic9(t *testing.T) {
	assert.Equal(t, "IX", FromArabic(9))
}

func TestFromArabic40(t *testing.T) {
	assert.Equal(t, "XL", FromArabic(40))
}

func TestFromArabic400(t *testing.T) {
	assert.Equal(t, "CD", FromArabic(400))
}
