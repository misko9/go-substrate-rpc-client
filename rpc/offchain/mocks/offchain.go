// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	offchain "github.com/ComposableFi/go-substrate-rpc-client/v4/rpc/offchain"
	types "github.com/ComposableFi/go-substrate-rpc-client/v4/types"
	mock "github.com/stretchr/testify/mock"
)

// Offchain is an autogenerated mock type for the Offchain type
type Offchain struct {
	mock.Mock
}

// LocalStorageGet provides a mock function with given fields: kind, key
func (_m *Offchain) LocalStorageGet(kind offchain.StorageKind, key []byte) (*types.StorageDataRaw, error) {
	ret := _m.Called(kind, key)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(offchain.StorageKind, []byte) *types.StorageDataRaw); ok {
		r0 = rf(kind, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(offchain.StorageKind, []byte) error); ok {
		r1 = rf(kind, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalStorageSet provides a mock function with given fields: kind, key, value
func (_m *Offchain) LocalStorageSet(kind offchain.StorageKind, key []byte, value []byte) error {
	ret := _m.Called(kind, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(offchain.StorageKind, []byte, []byte) error); ok {
		r0 = rf(kind, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewOffchainT interface {
	mock.TestingT
	Cleanup(func())
}

// NewOffchain creates a new instance of Offchain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOffchain(t NewOffchainT) *Offchain {
	mock := &Offchain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
