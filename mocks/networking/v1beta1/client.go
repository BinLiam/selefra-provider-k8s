package v1beta1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockNetworkingV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkingV1beta1InterfaceMockRecorder
}

type MockNetworkingV1beta1InterfaceMockRecorder struct {
	mock *MockNetworkingV1beta1Interface
}

func NewMockNetworkingV1beta1Interface(ctrl *gomock.Controller) *MockNetworkingV1beta1Interface {
	mock := &MockNetworkingV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockNetworkingV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockNetworkingV1beta1Interface) EXPECT() *MockNetworkingV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockNetworkingV1beta1Interface) IngressClasses() v1beta1.IngressClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.IngressClasses)
	ret0, _ := ret[0].(v1beta1.IngressClassInterface)
	return ret0
}

func (mr *MockNetworkingV1beta1InterfaceMockRecorder) IngressClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.IngressClasses, reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).IngressClasses))
}

func (m *MockNetworkingV1beta1Interface) Ingresses(arg0 string) v1beta1.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Ingresses, arg0)
	ret0, _ := ret[0].(v1beta1.IngressInterface)
	return ret0
}

func (mr *MockNetworkingV1beta1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Ingresses, reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).Ingresses), arg0)
}

func (m *MockNetworkingV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockNetworkingV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockNetworkingV1beta1Interface)(nil).RESTClient))
}
