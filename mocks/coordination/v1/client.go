package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	rest "k8s.io/client-go/rest"
)

type MockCoordinationV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinationV1InterfaceMockRecorder
}

type MockCoordinationV1InterfaceMockRecorder struct {
	mock *MockCoordinationV1Interface
}

func NewMockCoordinationV1Interface(ctrl *gomock.Controller) *MockCoordinationV1Interface {
	mock := &MockCoordinationV1Interface{ctrl: ctrl}
	mock.recorder = &MockCoordinationV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockCoordinationV1Interface) EXPECT() *MockCoordinationV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockCoordinationV1Interface) Leases(arg0 string) v1.LeaseInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Leases, arg0)
	ret0, _ := ret[0].(v1.LeaseInterface)
	return ret0
}

func (mr *MockCoordinationV1InterfaceMockRecorder) Leases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Leases, reflect.TypeOf((*MockCoordinationV1Interface)(nil).Leases), arg0)
}

func (m *MockCoordinationV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockCoordinationV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockCoordinationV1Interface)(nil).RESTClient))
}
