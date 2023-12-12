package extract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingExtract(t *testing.T) {
	// Arrange
	extractor := NewPingExtrator()
	input := "64 bytes from 8.8.8.8: icmp_seq=1 ttl=117 time=284 ms"

	// Act
	result, err := extractor.Extract(input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, float64(284), result.Value)
}
