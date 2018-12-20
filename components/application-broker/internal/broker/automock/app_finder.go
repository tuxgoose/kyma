// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import internal "github.com/kyma-project/kyma/components/application-broker/internal"
import mock "github.com/stretchr/testify/mock"

// AppFinder is an autogenerated mock type for the AppFinder type
type AppFinder struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *AppFinder) FindAll() ([]*internal.Application, error) {
	ret := _m.Called()

	var r0 []*internal.Application
	if rf, ok := ret.Get(0).(func() []*internal.Application); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internal.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneByServiceID provides a mock function with given fields: id
func (_m *AppFinder) FindOneByServiceID(id internal.ApplicationServiceID) (*internal.Application, error) {
	ret := _m.Called(id)

	var r0 *internal.Application
	if rf, ok := ret.Get(0).(func(internal.ApplicationServiceID) *internal.Application); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ApplicationServiceID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: name
func (_m *AppFinder) Get(name internal.ApplicationName) (*internal.Application, error) {
	ret := _m.Called(name)

	var r0 *internal.Application
	if rf, ok := ret.Get(0).(func(internal.ApplicationName) *internal.Application); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ApplicationName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}