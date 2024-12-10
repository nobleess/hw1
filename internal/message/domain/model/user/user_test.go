package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	u := New("Login")
	assert.Equal(t, u.Login, "Login")
	assert.Equal(t, u.id, u.ID())
	assert.Equal(t, u.id, ID(1))
}
