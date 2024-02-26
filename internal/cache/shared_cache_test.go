package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShardedCache_Get(t *testing.T) {
	mockCache := &MockShardedCache{}
	expectedValue := "theValue"
	mockCache.On("get", "theKey").Return(expectedValue, nil)

	val, err := mockCache.get("theKey")
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, val)

	mockCache.AssertExpectations(t)
}

func TestShardedCacheImpl_Set(t *testing.T) {
	mockCache := &MockShardedCache{}
	mockCache.On("set", "theKey", "theValue").Return().Once() // Настройте поведение мока
	mockCache.set("theKey", "theValue")

	mockCache.AssertExpectations(t)
}
