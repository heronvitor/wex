// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/heronvitor/pkg/entities"
	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetPurchaseByID provides a mock function with given fields: id
func (_m *TransactionRepository) GetPurchaseByID(id string) (*entities.Purchase, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPurchaseByID")
	}

	var r0 *entities.Purchase
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Purchase, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Purchase); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Purchase)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePurchase provides a mock function with given fields: _a0
func (_m *TransactionRepository) SavePurchase(_a0 entities.Purchase) (entities.Purchase, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SavePurchase")
	}

	var r0 entities.Purchase
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Purchase) (entities.Purchase, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entities.Purchase) entities.Purchase); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entities.Purchase)
	}

	if rf, ok := ret.Get(1).(func(entities.Purchase) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}