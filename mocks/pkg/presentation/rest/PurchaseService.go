// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/heronvitor/pkg/entities"
	mock "github.com/stretchr/testify/mock"
)

// PurchaseService is an autogenerated mock type for the PurchaseService type
type PurchaseService struct {
	mock.Mock
}

// CreatePurchase provides a mock function with given fields: purchase
func (_m *PurchaseService) CreatePurchase(purchase entities.Purchase) (entities.Purchase, error) {
	ret := _m.Called(purchase)

	if len(ret) == 0 {
		panic("no return value specified for CreatePurchase")
	}

	var r0 entities.Purchase
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Purchase) (entities.Purchase, error)); ok {
		return rf(purchase)
	}
	if rf, ok := ret.Get(0).(func(entities.Purchase) entities.Purchase); ok {
		r0 = rf(purchase)
	} else {
		r0 = ret.Get(0).(entities.Purchase)
	}

	if rf, ok := ret.Get(1).(func(entities.Purchase) error); ok {
		r1 = rf(purchase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPurchaseInCurrency provides a mock function with given fields: id, currency
func (_m *PurchaseService) GetPurchaseInCurrency(id string, currency string) (*entities.PurchaseInCurrency, error) {
	ret := _m.Called(id, currency)

	if len(ret) == 0 {
		panic("no return value specified for GetPurchaseInCurrency")
	}

	var r0 *entities.PurchaseInCurrency
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*entities.PurchaseInCurrency, error)); ok {
		return rf(id, currency)
	}
	if rf, ok := ret.Get(0).(func(string, string) *entities.PurchaseInCurrency); ok {
		r0 = rf(id, currency)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.PurchaseInCurrency)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPurchaseService creates a new instance of PurchaseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPurchaseService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PurchaseService {
	mock := &PurchaseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
