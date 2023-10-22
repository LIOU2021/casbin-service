package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEmpty(t *testing.T) {
	var str string
	var num int
	var slice []int

	assert.True(t, IsEmpty(str))
	assert.True(t, IsEmpty(num))
	assert.True(t, IsEmpty(slice))

	str = "Hello"
	num = 42
	slice = []int{1, 2, 3}

	assert.False(t, IsEmpty(str))
	assert.False(t, IsEmpty(num))
	assert.False(t, IsEmpty(slice))
}
