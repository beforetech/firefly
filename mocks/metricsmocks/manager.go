// Code generated by mockery v2.46.0. DO NOT EDIT.

package metricsmocks

import (
	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"
	core "github.com/hyperledger/firefly/pkg/core"

	metrics "github.com/hyperledger/firefly/internal/metrics"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// AddTime provides a mock function with given fields: id
func (_m *Manager) AddTime(id string) {
	_m.Called(id)
}

// BlockchainContractDeployment provides a mock function with given fields:
func (_m *Manager) BlockchainContractDeployment() {
	_m.Called()
}

// BlockchainEvent provides a mock function with given fields: location, signature
func (_m *Manager) BlockchainEvent(location string, signature string) {
	_m.Called(location, signature)
}

// BlockchainQuery provides a mock function with given fields: location, methodName
func (_m *Manager) BlockchainQuery(location string, methodName string) {
	_m.Called(location, methodName)
}

// BlockchainTransaction provides a mock function with given fields: location, methodName
func (_m *Manager) BlockchainTransaction(location string, methodName string) {
	_m.Called(location, methodName)
}

// CountBatchPin provides a mock function with given fields: namespace
func (_m *Manager) CountBatchPin(namespace string) {
	_m.Called(namespace)
}

// DeleteTime provides a mock function with given fields: id
func (_m *Manager) DeleteTime(id string) {
	_m.Called(id)
}

// GetTime provides a mock function with given fields: id
func (_m *Manager) GetTime(id string) time.Time {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTime")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(string) time.Time); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// IsMetricsEnabled provides a mock function with given fields:
func (_m *Manager) IsMetricsEnabled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsMetricsEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MessageConfirmed provides a mock function with given fields: msg, eventType
func (_m *Manager) MessageConfirmed(msg *core.Message, eventType fftypes.FFEnum) {
	_m.Called(msg, eventType)
}

// MessageSubmitted provides a mock function with given fields: msg
func (_m *Manager) MessageSubmitted(msg *core.Message) {
	_m.Called(msg)
}

// NodeIdentityDXCertExpiry provides a mock function with given fields: namespace, expiry
func (_m *Manager) NodeIdentityDXCertExpiry(namespace string, expiry time.Time) {
	_m.Called(namespace, expiry)
}

// NodeIdentityDXCertMismatch provides a mock function with given fields: namespace, mismatch
func (_m *Manager) NodeIdentityDXCertMismatch(namespace string, mismatch metrics.NodeIdentityDXCertMismatchStatus) {
	_m.Called(namespace, mismatch)
}

// TransferConfirmed provides a mock function with given fields: transfer
func (_m *Manager) TransferConfirmed(transfer *core.TokenTransfer) {
	_m.Called(transfer)
}

// TransferSubmitted provides a mock function with given fields: transfer
func (_m *Manager) TransferSubmitted(transfer *core.TokenTransfer) {
	_m.Called(transfer)
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
