// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import mock "github.com/stretchr/testify/mock"

// decimalDecompose is an autogenerated mock type for the decimalDecompose type
type decimalDecompose struct {
	mock.Mock
}

// Decompose provides a mock function with given fields: buf
func (_m *decimalDecompose) Decompose(buf []byte) (byte, bool, []byte, int32) {
	ret := _m.Called(buf)

	var r0 byte
	if rf, ok := ret.Get(0).(func([]byte) byte); ok {
		r0 = rf(buf)
	} else {
		r0 = ret.Get(0).(byte)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func([]byte) bool); ok {
		r1 = rf(buf)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 []byte
	if rf, ok := ret.Get(2).(func([]byte) []byte); ok {
		r2 = rf(buf)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]byte)
		}
	}

	var r3 int32
	if rf, ok := ret.Get(3).(func([]byte) int32); ok {
		r3 = rf(buf)
	} else {
		r3 = ret.Get(3).(int32)
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTnewDecimalDecompose interface {
	mock.TestingT
	Cleanup(func())
}

// newDecimalDecompose creates a new instance of decimalDecompose. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newDecimalDecompose(t mockConstructorTestingTnewDecimalDecompose) *decimalDecompose {
	mock := &decimalDecompose{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
