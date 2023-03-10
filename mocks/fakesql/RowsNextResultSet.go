// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"
)

// RowsNextResultSet is an autogenerated mock type for the RowsNextResultSet type
type RowsNextResultSet struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *RowsNextResultSet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Columns provides a mock function with given fields:
func (_m *RowsNextResultSet) Columns() []string {
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

// HasNextResultSet provides a mock function with given fields:
func (_m *RowsNextResultSet) HasNextResultSet() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields: dest
func (_m *RowsNextResultSet) Next(dest []driver.Value) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func([]driver.Value) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextResultSet provides a mock function with given fields:
func (_m *RowsNextResultSet) NextResultSet() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRowsNextResultSet interface {
	mock.TestingT
	Cleanup(func())
}

// NewRowsNextResultSet creates a new instance of RowsNextResultSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRowsNextResultSet(t mockConstructorTestingTNewRowsNextResultSet) *RowsNextResultSet {
	mock := &RowsNextResultSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
