// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"

// BrokerSyncer is an autogenerated mock type for the BrokerSyncer type
type BrokerSyncer struct {
	mock.Mock
}

// Sync provides a mock function with given fields: name, maxSyncRetries
func (_m *BrokerSyncer) Sync(name string, maxSyncRetries int) error {
	ret := _m.Called(name, maxSyncRetries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(name, maxSyncRetries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
