package helper

import (
	"github.com/stretchr/testify/mock"
)

type Field interface {
	Index() int
	Name() string

	SetIndex(index int) Field
	SetName(name string) Field
}

// FakeField is an autogenerated mock type for the FakeField type
type FakeField struct {
	mock.Mock
}

// Index provides a mock function with given fields:
func (_m *FakeField) Index() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *FakeField) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetIndex provides a mock function with given fields: index
func (_m *FakeField) SetIndex(index int) Field {
	ret := _m.Called(index)

	var r0 Field
	if rf, ok := ret.Get(0).(func(int) Field); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Field)
		}
	}

	return r0
}

// SetName provides a mock function with given fields: name
func (_m *FakeField) SetName(name string) Field {
	ret := _m.Called(name)

	var r0 Field
	if rf, ok := ret.Get(0).(func(string) Field); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Field)
		}
	}

	return r0
}

type mockConstructorTestingTNewField interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeField creates a new instance of FakeField. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeField(t mockConstructorTestingTNewField) *FakeField {
	mock := &FakeField{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}