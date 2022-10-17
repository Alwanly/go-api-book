package redis

import (
	"net/url"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestCacheURL(t *testing.T) {

	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	urlStr := "redis://" + s.Addr()
	u, err := url.Parse(urlStr)
	assert.Nil(t, err)
	assert.NotNil(t, u)

	cache, err := Initialize()
	assert.Nil(t, err)
	assert.NotNil(t, cache)

	rc, ok := cache.(*Cache)
	assert.True(t, ok)
	assert.NotNil(t, rc)
}
