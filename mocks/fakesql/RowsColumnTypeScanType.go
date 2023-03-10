// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"

	reflect "reflect"
)

// RowsColumnTypeScanType is an autogenerated mock type for the RowsColumnTypeScanType type
type RowsColumnTypeScanType struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *RowsColumnTypeScanType) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ColumnTypeScanType provides a mock function with given fields: index
func (_m *RowsColumnTypeScanType) ColumnTypeScanType(index int) reflect.Type {
	ret := _m.Called(index)

	var r0 reflect.Type
	if rf, ok := ret.Get(0).(func(int) reflect.Type); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reflect.Type)
		}
	}

	return r0
}

// Columns provides a mock function with given fields:
func (_m *RowsColumnTypeScanType) Columns() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Next provides a mock function with given fields: dest
func (_m *RowsColumnTypeScanType) Next(dest []driver.Value) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func([]driver.Value) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRowsColumnTypeScanType interface {
	mock.TestingT
	Cleanup(func())
}

// NewRowsColumnTypeScanType creates a new instance of RowsColumnTypeScanType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRowsColumnTypeScanType(t mockConstructorTestingTNewRowsColumnTypeScanType) *RowsColumnTypeScanType {
	mock := &RowsColumnTypeScanType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
