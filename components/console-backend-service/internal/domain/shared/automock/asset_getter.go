// Code generated by mockery v1.0.0. DO NOT EDIT.
package automock

import mock "github.com/stretchr/testify/mock"

import v1alpha2 "github.com/kyma-project/kyma/components/asset-store-controller-manager/pkg/apis/assetstore/v1alpha2"

// AssetGetter is an autogenerated mock type for the AssetGetter type
type AssetGetter struct {
	mock.Mock
}

// ListForDocsTopicByType provides a mock function with given fields: namespace, docsTopicName, types
func (_m *AssetGetter) ListForDocsTopicByType(namespace string, docsTopicName string, types []string) ([]*v1alpha2.Asset, error) {
	ret := _m.Called(namespace, docsTopicName, types)

	var r0 []*v1alpha2.Asset
	if rf, ok := ret.Get(0).(func(string, string, []string) []*v1alpha2.Asset); ok {
		r0 = rf(namespace, docsTopicName, types)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha2.Asset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(namespace, docsTopicName, types)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
