// Code generated by mockery v2.33.0. DO NOT EDIT.

package ports

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUserRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockUserRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockUserRepository_Delete_Call {
	return &MockUserRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockUserRepository_Delete_Call) Run(run func(ctx context.Context, id string)) *MockUserRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUserRepository_Delete_Call) Return(_a0 error) *MockUserRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockUserRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, id, output
func (_m *MockUserRepository) Read(ctx context.Context, id string, output interface{}) error {
	ret := _m.Called(ctx, id, output)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, id, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockUserRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - output interface{}
func (_e *MockUserRepository_Expecter) Read(ctx interface{}, id interface{}, output interface{}) *MockUserRepository_Read_Call {
	return &MockUserRepository_Read_Call{Call: _e.mock.On("Read", ctx, id, output)}
}

func (_c *MockUserRepository_Read_Call) Run(run func(ctx context.Context, id string, output interface{})) *MockUserRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockUserRepository_Read_Call) Return(_a0 error) *MockUserRepository_Read_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Read_Call) RunAndReturn(run func(context.Context, string, interface{}) error) *MockUserRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadAll provides a mock function with given fields: ctx, output
func (_m *MockUserRepository) ReadAll(ctx context.Context, output interface{}) error {
	ret := _m.Called(ctx, output)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_ReadAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadAll'
type MockUserRepository_ReadAll_Call struct {
	*mock.Call
}

// ReadAll is a helper method to define mock.On call
//   - ctx context.Context
//   - output interface{}
func (_e *MockUserRepository_Expecter) ReadAll(ctx interface{}, output interface{}) *MockUserRepository_ReadAll_Call {
	return &MockUserRepository_ReadAll_Call{Call: _e.mock.On("ReadAll", ctx, output)}
}

func (_c *MockUserRepository_ReadAll_Call) Run(run func(ctx context.Context, output interface{})) *MockUserRepository_ReadAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *MockUserRepository_ReadAll_Call) Return(_a0 error) *MockUserRepository_ReadAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_ReadAll_Call) RunAndReturn(run func(context.Context, interface{}) error) *MockUserRepository_ReadAll_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, data, output
func (_m *MockUserRepository) Save(ctx context.Context, data interface{}, output interface{}) error {
	ret := _m.Called(ctx, data, output)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, data, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockUserRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - data interface{}
//   - output interface{}
func (_e *MockUserRepository_Expecter) Save(ctx interface{}, data interface{}, output interface{}) *MockUserRepository_Save_Call {
	return &MockUserRepository_Save_Call{Call: _e.mock.On("Save", ctx, data, output)}
}

func (_c *MockUserRepository_Save_Call) Run(run func(ctx context.Context, data interface{}, output interface{})) *MockUserRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(interface{}))
	})
	return _c
}

func (_c *MockUserRepository_Save_Call) Return(_a0 error) *MockUserRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Save_Call) RunAndReturn(run func(context.Context, interface{}, interface{}) error) *MockUserRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, data, output
func (_m *MockUserRepository) Update(ctx context.Context, id string, data interface{}, output interface{}) error {
	ret := _m.Called(ctx, id, data, output)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, interface{}) error); ok {
		r0 = rf(ctx, id, data, output)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - data interface{}
//   - output interface{}
func (_e *MockUserRepository_Expecter) Update(ctx interface{}, id interface{}, data interface{}, output interface{}) *MockUserRepository_Update_Call {
	return &MockUserRepository_Update_Call{Call: _e.mock.On("Update", ctx, id, data, output)}
}

func (_c *MockUserRepository_Update_Call) Run(run func(ctx context.Context, id string, data interface{}, output interface{})) *MockUserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(interface{}))
	})
	return _c
}

func (_c *MockUserRepository_Update_Call) Return(_a0 error) *MockUserRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Update_Call) RunAndReturn(run func(context.Context, string, interface{}, interface{}) error) *MockUserRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
