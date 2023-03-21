package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	rest "k8s.io/client-go/rest"
)

type MockNetworkingV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkingV1InterfaceMockRecorder
}

type MockNetworkingV1InterfaceMockRecorder struct {
	mock *MockNetworkingV1Interface
}

func NewMockNetworkingV1Interface(ctrl *gomock.Controller) *MockNetworkingV1Interface {
	mock := &MockNetworkingV1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockNetworkingV1Interface) EXPECT() *MockNetworkingV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockNetworkingV1Interface) IngressClasses() v1.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.IngressClasses)
	ret0, _ := ret[0].(v1.IngressClassInterface)
	return ret0
}

func (mr *MockNetworkingV1InterfaceMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.IngressClasses, reflect.TypeOf((*MockNetworkingV1Interface)(nil).IngressClasses))
}

func (m *MockNetworkingV1Interface) Ingresses(arg0 string) v1.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Ingresses, arg0)
	ret0, _ := ret[0].(v1.IngressInterface)
	return ret0
}

func (mr *MockNetworkingV1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Ingresses, reflect.TypeOf((*MockNetworkingV1Interface)(nil).Ingresses), arg0)
}

func (m *MockNetworkingV1Interface) NetworkPolicies(arg0 string) v1.NetworkPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkPolicies, arg0)
	ret0, _ := ret[0].(v1.NetworkPolicyInterface)
	return ret0
}

func (mr *MockNetworkingV1InterfaceMockRecorder) NetworkPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkPolicies, reflect.TypeOf((*MockNetworkingV1Interface)(nil).NetworkPolicies), arg0)
}

func (m *MockNetworkingV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockNetworkingV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockNetworkingV1Interface)(nil).RESTClient))
}
