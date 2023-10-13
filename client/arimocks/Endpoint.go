// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// Endpoint is an autogenerated mock type for the Endpoint type
type Endpoint struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *Endpoint) Data(key *ari.Key) (*ari.EndpointData, error) {
	ret := _m.Called(key)

	var r0 *ari.EndpointData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.EndpointData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.EndpointData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *Endpoint) Get(key *ari.Key) *ari.EndpointHandle {
	ret := _m.Called(key)

	var r0 *ari.EndpointHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.EndpointHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.EndpointHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: filter
func (_m *Endpoint) List(filter *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(filter)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(*ari.Key) []*ari.Key); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByTech provides a mock function with given fields: tech, filter
func (_m *Endpoint) ListByTech(tech string, filter *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(tech, filter)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(string, *ari.Key) []*ari.Key); ok {
		r0 = rf(tech, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *ari.Key) error); ok {
		r1 = rf(tech, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
