// Code generated by mockery v2.46.0. DO NOT EDIT.

package tokenmocks

import (
	context "context"

	tokens "github.com/hyperledger/firefly/pkg/tokens"
	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// TokenPoolCreated provides a mock function with given fields: ctx, plugin, pool
func (_m *Callbacks) TokenPoolCreated(ctx context.Context, plugin tokens.Plugin, pool *tokens.TokenPool) error {
	ret := _m.Called(ctx, plugin, pool)

	if len(ret) == 0 {
		panic("no return value specified for TokenPoolCreated")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tokens.Plugin, *tokens.TokenPool) error); ok {
		r0 = rf(ctx, plugin, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokensApproved provides a mock function with given fields: plugin, approval
func (_m *Callbacks) TokensApproved(plugin tokens.Plugin, approval *tokens.TokenApproval) error {
	ret := _m.Called(plugin, approval)

	if len(ret) == 0 {
		panic("no return value specified for TokensApproved")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(tokens.Plugin, *tokens.TokenApproval) error); ok {
		r0 = rf(plugin, approval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokensTransferred provides a mock function with given fields: plugin, transfer
func (_m *Callbacks) TokensTransferred(plugin tokens.Plugin, transfer *tokens.TokenTransfer) error {
	ret := _m.Called(plugin, transfer)

	if len(ret) == 0 {
		panic("no return value specified for TokensTransferred")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(tokens.Plugin, *tokens.TokenTransfer) error); ok {
		r0 = rf(plugin, transfer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCallbacks creates a new instance of Callbacks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCallbacks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Callbacks {
	mock := &Callbacks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
