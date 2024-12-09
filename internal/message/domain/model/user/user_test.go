package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	u := New("Login")
	assert.Equal(t, u.Login, "Login")
	assert.Equal(t, u.id, u.ID())
	assert.Equal(t, u.id, ID(1))
}
