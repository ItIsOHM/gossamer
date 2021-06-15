// Code generated by mockery v2.8.0. DO NOT EDIT.

package sync

import (
	runtime "github.com/ChainSafe/gossamer/lib/runtime"
	mock "github.com/stretchr/testify/mock"
)

// MockBlockProducer is an autogenerated mock type for the BlockProducer type
type MockBlockProducer struct {
	mock.Mock
}

// Pause provides a mock function with given fields:
func (_m *MockBlockProducer) Pause() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Resume provides a mock function with given fields:
func (_m *MockBlockProducer) Resume() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRuntime provides a mock function with given fields: rt
func (_m *MockBlockProducer) SetRuntime(rt runtime.Instance) {
	_m.Called(rt)
}
