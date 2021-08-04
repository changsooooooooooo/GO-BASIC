package square

import "testing"
import "github.com/stretchr/testify/assert"

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, Square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(40, Square(8), "square(9) should be 81")
}
