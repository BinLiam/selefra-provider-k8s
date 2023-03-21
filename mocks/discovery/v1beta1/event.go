package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockEventsV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockEventsV1beta1InterfaceMockRecorder
}

type MockEventsV1beta1InterfaceMockRecorder struct {
	mock *MockEventsV1beta1Interface
}

func NewMockEventsV1beta1Interface(ctrl *gomock.Controller) *MockEventsV1beta1Interface {
	mock := &MockEventsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockEventsV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockEventsV1beta1Interface) EXPECT() *MockEventsV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockEventsV1beta1Interface) Events(arg0 string) v1beta1.EventInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Events, arg0)
	ret0, _ := ret[0].(v1beta1.EventInterface)
	return ret0
}

func (mr *MockEventsV1beta1InterfaceMockRecorder) Events(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Events, reflect.TypeOf((*MockEventsV1beta1Interface)(nil).Events), arg0)
}

func (m *MockEventsV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockEventsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockEventsV1beta1Interface)(nil).RESTClient))
}
