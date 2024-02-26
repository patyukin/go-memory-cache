// Code generated by mockery v2.42.0. DO NOT EDIT.

package cache

import mock "github.com/stretchr/testify/mock"

// MockShardedCache is an autogenerated mock type for the ShardedCache type
type MockShardedCache struct {
	mock.Mock
}

type MockShardedCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShardedCache) EXPECT() *MockShardedCache_Expecter {
	return &MockShardedCache_Expecter{mock: &_m.Mock}
}

// get provides a mock function with given fields: key
func (_m *MockShardedCache) get(key string) (string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShardedCache_get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'get'
type MockShardedCache_get_Call struct {
	*mock.Call
}

// get is a helper method to define mock.On call
//   - key string
func (_e *MockShardedCache_Expecter) get(key interface{}) *MockShardedCache_get_Call {
	return &MockShardedCache_get_Call{Call: _e.mock.On("get", key)}
}

func (_c *MockShardedCache_get_Call) Run(run func(key string)) *MockShardedCache_get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockShardedCache_get_Call) Return(_a0 string, _a1 error) *MockShardedCache_get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShardedCache_get_Call) RunAndReturn(run func(string) (string, error)) *MockShardedCache_get_Call {
	_c.Call.Return(run)
	return _c
}

// set provides a mock function with given fields: key, value
func (_m *MockShardedCache) set(key string, value string) {
	_m.Called(key, value)
}

// MockShardedCache_set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'set'
type MockShardedCache_set_Call struct {
	*mock.Call
}

// set is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MockShardedCache_Expecter) set(key interface{}, value interface{}) *MockShardedCache_set_Call {
	return &MockShardedCache_set_Call{Call: _e.mock.On("set", key, value)}
}

func (_c *MockShardedCache_set_Call) Run(run func(key string, value string)) *MockShardedCache_set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockShardedCache_set_Call) Return() *MockShardedCache_set_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockShardedCache_set_Call) RunAndReturn(run func(string, string)) *MockShardedCache_set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockShardedCache creates a new instance of MockShardedCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShardedCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShardedCache {
	mock := &MockShardedCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
