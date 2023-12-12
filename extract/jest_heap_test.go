package extract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJestHeapExtract(t *testing.T) {
	// Arrange
	extractor := NewJestHeapExtrator()
	input := "PASS js/027.test.js (399 MB heap size)"

	// Act
	result, err := extractor.Extract(input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, float64(399), result.Value)
}
