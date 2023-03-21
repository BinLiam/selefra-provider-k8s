package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	openapi_v2 "github.com/google/gnostic/openapiv2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	version "k8s.io/apimachinery/pkg/version"
	discovery "k8s.io/client-go/discovery"
	openapi "k8s.io/client-go/openapi"
	rest "k8s.io/client-go/rest"
)

type MockDiscoveryInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockDiscoveryInterfaceMockRecorder
}

type MockDiscoveryInterfaceMockRecorder struct {
	mock *MockDiscoveryInterface
}

func NewMockDiscoveryInterface(ctrl *gomock.Controller) *MockDiscoveryInterface {
	mock := &MockDiscoveryInterface{ctrl: ctrl}
	mock.recorder = &MockDiscoveryInterfaceMockRecorder{mock}
	return mock
}

func (m *MockDiscoveryInterface) EXPECT() *MockDiscoveryInterfaceMockRecorder {
	return m.recorder
}

func (m *MockDiscoveryInterface) OpenAPISchema() (*openapi_v2.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.OpenAPISchema)
	ret0, _ := ret[0].(*openapi_v2.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) OpenAPISchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.OpenAPISchema, reflect.TypeOf((*MockDiscoveryInterface)(nil).OpenAPISchema))
}

func (m *MockDiscoveryInterface) OpenAPIV3() openapi.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.OpenAPIV)
	ret0, _ := ret[0].(openapi.Client)
	return ret0
}

func (mr *MockDiscoveryInterfaceMockRecorder) OpenAPIV3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.OpenAPIV, reflect.TypeOf((*MockDiscoveryInterface)(nil).OpenAPIV3))
}

func (m *MockDiscoveryInterface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockDiscoveryInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockDiscoveryInterface)(nil).RESTClient))
}

func (m *MockDiscoveryInterface) ServerGroups() (*v1.APIGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerGroups)
	ret0, _ := ret[0].(*v1.APIGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerGroups, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerGroups))
}

func (m *MockDiscoveryInterface) ServerGroupsAndResources() ([]*v1.APIGroup, []*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerGroupsAndResources)
	ret0, _ := ret[0].([]*v1.APIGroup)
	ret1, _ := ret[1].([]*v1.APIResourceList)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerGroupsAndResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerGroupsAndResources, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerGroupsAndResources))
}

func (m *MockDiscoveryInterface) ServerPreferredNamespacedResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerPreferredNamespacedResources)
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerPreferredNamespacedResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerPreferredNamespacedResources, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerPreferredNamespacedResources))
}

func (m *MockDiscoveryInterface) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerPreferredResources)
	ret0, _ := ret[0].([]*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerPreferredResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerPreferredResources, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerPreferredResources))
}

func (m *MockDiscoveryInterface) ServerResourcesForGroupVersion(arg0 string) (*v1.APIResourceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerResourcesForGroupVersion, arg0)
	ret0, _ := ret[0].(*v1.APIResourceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerResourcesForGroupVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerResourcesForGroupVersion, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerResourcesForGroupVersion), arg0)
}

func (m *MockDiscoveryInterface) ServerVersion() (*version.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServerVersion)
	ret0, _ := ret[0].(*version.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDiscoveryInterfaceMockRecorder) ServerVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServerVersion, reflect.TypeOf((*MockDiscoveryInterface)(nil).ServerVersion))
}

func (m *MockDiscoveryInterface) WithLegacy() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.WithLegacy)
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

func (mr *MockDiscoveryInterfaceMockRecorder) WithLegacy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.WithLegacy, reflect.TypeOf((*MockDiscoveryInterface)(nil).WithLegacy))
}
