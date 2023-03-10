// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	context "context"

	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"
)

// QueryerContext is an autogenerated mock type for the QueryerContext type
type QueryerContext struct {
	mock.Mock
}

// QueryContext provides a mock function with given fields: ctx, query, args
func (_m *QueryerContext) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	ret := _m.Called(ctx, query, args)

	var r0 driver.Rows
	if rf, ok := ret.Get(0).(func(context.Context, string, []driver.NamedValue) driver.Rows); ok {
		r0 = rf(ctx, query, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []driver.NamedValue) error); ok {
		r1 = rf(ctx, query, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQueryerContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryerContext creates a new instance of QueryerContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryerContext(t mockConstructorTestingTNewQueryerContext) *QueryerContext {
	mock := &QueryerContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
