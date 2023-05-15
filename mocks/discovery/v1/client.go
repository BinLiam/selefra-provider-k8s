package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	rest "k8s.io/client-go/rest"
)

type MockDiscoveryV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoveryV1InterfaceMockRecorder
}

type MockDiscoveryV1InterfaceMockRecorder struct {
	mock *MockDiscoveryV1Interface
}

func NewMockDiscoveryV1Interface(ctrl *gomock.Controller) *MockDiscoveryV1Interface {
	mock := &MockDiscoveryV1Interface{ctrl: ctrl}
	mock.recorder = &MockDiscoveryV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockDiscoveryV1Interface) EXPECT() *MockDiscoveryV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockDiscoveryV1Interface) EndpointSlices(arg0 string) v1.EndpointSliceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EndpointSlices, arg0)
	ret0, _ := ret[0].(v1.EndpointSliceInterface)
	return ret0
}

func (mr *MockDiscoveryV1InterfaceMockRecorder) EndpointSlices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EndpointSlices, reflect.TypeOf((*MockDiscoveryV1Interface)(nil).EndpointSlices), arg0)
}

func (m *MockDiscoveryV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockDiscoveryV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockDiscoveryV1Interface)(nil).RESTClient))
}
