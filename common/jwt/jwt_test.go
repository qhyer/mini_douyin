package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT(t *testing.T) {
	uid := int64(1)
	token, err := CreateTokenByID(uid)
	assert.Nil(t, err)
	uid2, err := ParseTokenToID(token)
	assert.Nil(t, err)
	assert.Equal(t, uid, uid2)
}
