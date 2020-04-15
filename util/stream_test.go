package util

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewStream(t *testing.T) {
	str := "abcd"
	s := NewStream(bytes.NewReader([]byte(str)), "$")
	assert.Equal(t, s.Next(), "a")
	assert.Equal(t, s.Next(), "b")

	assert.Equal(t, s.Peek(), "c")
	assert.Equal(t, s.Peek(), "c")

	s.PutBack("b")
	assert.Equal(t, s.Peek(), "b")
	assert.Equal(t, s.Next(), "b")
	assert.Equal(t, s.Next(), "c")

	assert.Equal(t, s.HasNext(), true, "hasnext failed")
	assert.Equal(t, s.Next(), "d")
	assert.Equal(t, s.Next(), "$")
	assert.Equal(t, s.Next(), "$")

	assert.Equal(t, s.HasNext(), false, "hasnext failed")
}
