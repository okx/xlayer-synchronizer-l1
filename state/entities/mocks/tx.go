// Code generated by mockery. DO NOT EDIT.

package mock_entities

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// Tx is an autogenerated mock type for the Tx type
type Tx struct {
	mock.Mock
}

type Tx_Expecter struct {
	mock *mock.Mock
}

func (_m *Tx) EXPECT() *Tx_Expecter {
	return &Tx_Expecter{mock: &_m.Mock}
}

// AddCommitCallback provides a mock function with given fields: _a0
func (_m *Tx) AddCommitCallback(_a0 entities.TxCallbackType) {
	_m.Called(_a0)
}

// Tx_AddCommitCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddCommitCallback'
type Tx_AddCommitCallback_Call struct {
	*mock.Call
}

// AddCommitCallback is a helper method to define mock.On call
//   - _a0 entities.TxCallbackType
func (_e *Tx_Expecter) AddCommitCallback(_a0 interface{}) *Tx_AddCommitCallback_Call {
	return &Tx_AddCommitCallback_Call{Call: _e.mock.On("AddCommitCallback", _a0)}
}

func (_c *Tx_AddCommitCallback_Call) Run(run func(_a0 entities.TxCallbackType)) *Tx_AddCommitCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.TxCallbackType))
	})
	return _c
}

func (_c *Tx_AddCommitCallback_Call) Return() *Tx_AddCommitCallback_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tx_AddCommitCallback_Call) RunAndReturn(run func(entities.TxCallbackType)) *Tx_AddCommitCallback_Call {
	_c.Call.Return(run)
	return _c
}

// AddRollbackCallback provides a mock function with given fields: _a0
func (_m *Tx) AddRollbackCallback(_a0 entities.TxCallbackType) {
	_m.Called(_a0)
}

// Tx_AddRollbackCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRollbackCallback'
type Tx_AddRollbackCallback_Call struct {
	*mock.Call
}

// AddRollbackCallback is a helper method to define mock.On call
//   - _a0 entities.TxCallbackType
func (_e *Tx_Expecter) AddRollbackCallback(_a0 interface{}) *Tx_AddRollbackCallback_Call {
	return &Tx_AddRollbackCallback_Call{Call: _e.mock.On("AddRollbackCallback", _a0)}
}

func (_c *Tx_AddRollbackCallback_Call) Run(run func(_a0 entities.TxCallbackType)) *Tx_AddRollbackCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.TxCallbackType))
	})
	return _c
}

func (_c *Tx_AddRollbackCallback_Call) Return() *Tx_AddRollbackCallback_Call {
	_c.Call.Return()
	return _c
}

func (_c *Tx_AddRollbackCallback_Call) RunAndReturn(run func(entities.TxCallbackType)) *Tx_AddRollbackCallback_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: ctx
func (_m *Tx) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tx_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type Tx_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Tx_Expecter) Commit(ctx interface{}) *Tx_Commit_Call {
	return &Tx_Commit_Call{Call: _e.mock.On("Commit", ctx)}
}

func (_c *Tx_Commit_Call) Run(run func(ctx context.Context)) *Tx_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Tx_Commit_Call) Return(_a0 error) *Tx_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Commit_Call) RunAndReturn(run func(context.Context) error) *Tx_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: ctx
func (_m *Tx) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tx_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type Tx_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Tx_Expecter) Rollback(ctx interface{}) *Tx_Rollback_Call {
	return &Tx_Rollback_Call{Call: _e.mock.On("Rollback", ctx)}
}

func (_c *Tx_Rollback_Call) Run(run func(ctx context.Context)) *Tx_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Tx_Rollback_Call) Return(_a0 error) *Tx_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Tx_Rollback_Call) RunAndReturn(run func(context.Context) error) *Tx_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// NewTx creates a new instance of Tx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tx {
	mock := &Tx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
