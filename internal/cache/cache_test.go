package cache_test

import (
	"github.com/patyukin/go-memory-cache/internal/cache/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Get(t *testing.T) {
	mockCache := &mocks.Cache{}
	mockCache.On("Get", "someKey").Return("someValue", nil)
	value, err := mockCache.Get("someKey")
	if err != nil {
		t.FailNow()
	}

	mockCache.AssertCalled(t, "Get", "someKey")
	assert.Equal(t, "someValue", value)
}

func TestCache_Set(t *testing.T) {
	mockCache := &mocks.Cache{}
	mockCache.On("Set", "someKey", "someValue").Return(nil)
	mockCache.On("Get", "someKey").Return("someValue", nil)
	mockCache.Set("someKey", "someValue")
	value, err := mockCache.Get("someKey")
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, "someValue", value)
	mockCache.AssertCalled(t, "Set", "someKey", "someValue")
	mockCache.AssertCalled(t, "Get", "someKey")
}
