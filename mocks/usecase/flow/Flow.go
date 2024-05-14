// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "review-chatbot/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// Flow is an autogenerated mock type for the Flow type
type Flow struct {
	mock.Mock
}

type Flow_Expecter struct {
	mock *mock.Mock
}

func (_m *Flow) EXPECT() *Flow_Expecter {
	return &Flow_Expecter{mock: &_m.Mock}
}

// Answer provides a mock function with given fields: step, userAnswer
func (_m *Flow) Answer(step int, userAnswer string) (int, string) {
	ret := _m.Called(step, userAnswer)

	if len(ret) == 0 {
		panic("no return value specified for Answer")
	}

	var r0 int
	var r1 string
	if rf, ok := ret.Get(0).(func(int, string) (int, string)); ok {
		return rf(step, userAnswer)
	}
	if rf, ok := ret.Get(0).(func(int, string) int); ok {
		r0 = rf(step, userAnswer)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, string) string); ok {
		r1 = rf(step, userAnswer)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// Flow_Answer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Answer'
type Flow_Answer_Call struct {
	*mock.Call
}

// Answer is a helper method to define mock.On call
//   - step int
//   - userAnswer string
func (_e *Flow_Expecter) Answer(step interface{}, userAnswer interface{}) *Flow_Answer_Call {
	return &Flow_Answer_Call{Call: _e.mock.On("Answer", step, userAnswer)}
}

func (_c *Flow_Answer_Call) Run(run func(step int, userAnswer string)) *Flow_Answer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *Flow_Answer_Call) Return(_a0 int, _a1 string) *Flow_Answer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Flow_Answer_Call) RunAndReturn(run func(int, string) (int, string)) *Flow_Answer_Call {
	_c.Call.Return(run)
	return _c
}

// Ask provides a mock function with given fields: step
func (_m *Flow) Ask(step int) string {
	ret := _m.Called(step)

	if len(ret) == 0 {
		panic("no return value specified for Ask")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(step)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Flow_Ask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ask'
type Flow_Ask_Call struct {
	*mock.Call
}

// Ask is a helper method to define mock.On call
//   - step int
func (_e *Flow_Expecter) Ask(step interface{}) *Flow_Ask_Call {
	return &Flow_Ask_Call{Call: _e.mock.On("Ask", step)}
}

func (_c *Flow_Ask_Call) Run(run func(step int)) *Flow_Ask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Flow_Ask_Call) Return(_a0 string) *Flow_Ask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Flow_Ask_Call) RunAndReturn(run func(int) string) *Flow_Ask_Call {
	_c.Call.Return(run)
	return _c
}

// FinalStep provides a mock function with given fields:
func (_m *Flow) FinalStep() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FinalStep")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Flow_FinalStep_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FinalStep'
type Flow_FinalStep_Call struct {
	*mock.Call
}

// FinalStep is a helper method to define mock.On call
func (_e *Flow_Expecter) FinalStep() *Flow_FinalStep_Call {
	return &Flow_FinalStep_Call{Call: _e.mock.On("FinalStep")}
}

func (_c *Flow_FinalStep_Call) Run(run func()) *Flow_FinalStep_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Flow_FinalStep_Call) Return(_a0 int) *Flow_FinalStep_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Flow_FinalStep_Call) RunAndReturn(run func() int) *Flow_FinalStep_Call {
	_c.Call.Return(run)
	return _c
}

// GetHistory provides a mock function with given fields: ctx, customerID, orderID
func (_m *Flow) GetHistory(ctx context.Context, customerID string, orderID string) (entity.Chat, error) {
	ret := _m.Called(ctx, customerID, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetHistory")
	}

	var r0 entity.Chat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entity.Chat, error)); ok {
		return rf(ctx, customerID, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entity.Chat); ok {
		r0 = rf(ctx, customerID, orderID)
	} else {
		r0 = ret.Get(0).(entity.Chat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, customerID, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Flow_GetHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHistory'
type Flow_GetHistory_Call struct {
	*mock.Call
}

// GetHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - customerID string
//   - orderID string
func (_e *Flow_Expecter) GetHistory(ctx interface{}, customerID interface{}, orderID interface{}) *Flow_GetHistory_Call {
	return &Flow_GetHistory_Call{Call: _e.mock.On("GetHistory", ctx, customerID, orderID)}
}

func (_c *Flow_GetHistory_Call) Run(run func(ctx context.Context, customerID string, orderID string)) *Flow_GetHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Flow_GetHistory_Call) Return(_a0 entity.Chat, _a1 error) *Flow_GetHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Flow_GetHistory_Call) RunAndReturn(run func(context.Context, string, string) (entity.Chat, error)) *Flow_GetHistory_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *Flow) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Flow_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type Flow_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *Flow_Expecter) ID() *Flow_ID_Call {
	return &Flow_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *Flow_ID_Call) Run(run func()) *Flow_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Flow_ID_Call) Return(_a0 string) *Flow_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Flow_ID_Call) RunAndReturn(run func() string) *Flow_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Flow) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Flow_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Flow_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Flow_Expecter) Name() *Flow_Name_Call {
	return &Flow_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Flow_Name_Call) Run(run func()) *Flow_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Flow_Name_Call) Return(_a0 string) *Flow_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Flow_Name_Call) RunAndReturn(run func() string) *Flow_Name_Call {
	_c.Call.Return(run)
	return _c
}

// SaveHistory provides a mock function with given fields: ctx, step, customerID, orderID, history
func (_m *Flow) SaveHistory(ctx context.Context, step int, customerID string, orderID string, history string) error {
	ret := _m.Called(ctx, step, customerID, orderID, history)

	if len(ret) == 0 {
		panic("no return value specified for SaveHistory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) error); ok {
		r0 = rf(ctx, step, customerID, orderID, history)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flow_SaveHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveHistory'
type Flow_SaveHistory_Call struct {
	*mock.Call
}

// SaveHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - step int
//   - customerID string
//   - orderID string
//   - history string
func (_e *Flow_Expecter) SaveHistory(ctx interface{}, step interface{}, customerID interface{}, orderID interface{}, history interface{}) *Flow_SaveHistory_Call {
	return &Flow_SaveHistory_Call{Call: _e.mock.On("SaveHistory", ctx, step, customerID, orderID, history)}
}

func (_c *Flow_SaveHistory_Call) Run(run func(ctx context.Context, step int, customerID string, orderID string, history string)) *Flow_SaveHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *Flow_SaveHistory_Call) Return(_a0 error) *Flow_SaveHistory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Flow_SaveHistory_Call) RunAndReturn(run func(context.Context, int, string, string, string) error) *Flow_SaveHistory_Call {
	_c.Call.Return(run)
	return _c
}

// NewFlow creates a new instance of Flow. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlow(t interface {
	mock.TestingT
	Cleanup(func())
}) *Flow {
	mock := &Flow{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
