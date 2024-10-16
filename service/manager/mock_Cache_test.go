// Code generated by mockery v2.46.2. DO NOT EDIT.

package manager

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCache is an autogenerated mock type for the Cache type
type MockCache struct {
	mock.Mock
}

type MockCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCache) EXPECT() *MockCache_Expecter {
	return &MockCache_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key
func (_m *MockCache) Get(ctx context.Context, key string) (any, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 any
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (any, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) any); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCache_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockCache_Expecter) Get(ctx interface{}, key interface{}) *MockCache_Get_Call {
	return &MockCache_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *MockCache_Get_Call) Run(run func(ctx context.Context, key string)) *MockCache_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCache_Get_Call) Return(_a0 any, _a1 error) *MockCache_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_Get_Call) RunAndReturn(run func(context.Context, string) (any, error)) *MockCache_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockCache) GetAll(ctx context.Context) ([]any, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []any
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]any, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []any); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]any)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCache_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockCache_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCache_Expecter) GetAll(ctx interface{}) *MockCache_GetAll_Call {
	return &MockCache_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *MockCache_GetAll_Call) Run(run func(ctx context.Context)) *MockCache_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCache_GetAll_Call) Return(_a0 []any, _a1 error) *MockCache_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCache_GetAll_Call) RunAndReturn(run func(context.Context) ([]any, error)) *MockCache_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value
func (_m *MockCache) Set(ctx context.Context, key string, value any) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, any) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockCache_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value any
func (_e *MockCache_Expecter) Set(ctx interface{}, key interface{}, value interface{}) *MockCache_Set_Call {
	return &MockCache_Set_Call{Call: _e.mock.On("Set", ctx, key, value)}
}

func (_c *MockCache_Set_Call) Run(run func(ctx context.Context, key string, value any)) *MockCache_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(any))
	})
	return _c
}

func (_c *MockCache_Set_Call) Return(_a0 error) *MockCache_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_Set_Call) RunAndReturn(run func(context.Context, string, any) error) *MockCache_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
