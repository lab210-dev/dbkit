// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"
)

// ColumnConverter is an autogenerated mock type for the ColumnConverter type
type ColumnConverter struct {
	mock.Mock
}

// ColumnConverter provides a mock function with given fields: idx
func (_m *ColumnConverter) ColumnConverter(idx int) driver.ValueConverter {
	ret := _m.Called(idx)

	var r0 driver.ValueConverter
	if rf, ok := ret.Get(0).(func(int) driver.ValueConverter); ok {
		r0 = rf(idx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.ValueConverter)
		}
	}

	return r0
}

type mockConstructorTestingTNewColumnConverter interface {
	mock.TestingT
	Cleanup(func())
}

// NewColumnConverter creates a new instance of ColumnConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewColumnConverter(t mockConstructorTestingTNewColumnConverter) *ColumnConverter {
	mock := &ColumnConverter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
