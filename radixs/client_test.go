package radixs_test

import (
	"testing"

	"github.com/Akagi201/utilgo/radixs"
	"github.com/Akagi201/utilgo/tests"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c := radixs.New("localhost:6379")
	k := tests.RandStr()
	v := tests.RandStr()

	assert := assert.New(t)

	r, err := c.Cmd("SET", k, v).Str()
	assert.Nil(err)
	assert.Equal("OK", r)

	r, err = c.Cmd("GET", k).Str()
	assert.Nil(err)
	assert.Equal(v, r)

	ri, err := c.Cmd("DEL", k).Int()
	assert.Nil(err)
	assert.Equal(1, ri)
}
