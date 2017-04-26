package lib

import "testing"
import "github.com/stretchr/testify/assert"

func TestValid(t *testing.T) {
	f := GoTemplateFormatter{"Stars: {{.Stars}}"}
	r, err := f.format(Stats{"id", 42})

	assert.Equal(t, "Stars: 42", r)
	assert.Nil(t, err)
}
