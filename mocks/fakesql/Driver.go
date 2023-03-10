// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"
)

// FakeDriver is an autogenerated mock type for the FakeDriver type
type FakeDriver struct {
	mock.Mock
}

// Open provides a mock function with given fields: name
func (_m *FakeDriver) Open(name string) (driver.Conn, error) {
	ret := _m.Called(name)

	var r0 driver.Conn
	if rf, ok := ret.Get(0).(func(string) driver.Conn); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewDriver creates a new instance of FakeDriver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDriver(t mockConstructorTestingTNewDriver) *FakeDriver {
	mock := &FakeDriver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
