package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	rest "k8s.io/client-go/rest"
)

type MockAutoscalingV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAutoscalingV1InterfaceMockRecorder
}

type MockAutoscalingV1InterfaceMockRecorder struct {
	mock *MockAutoscalingV1Interface
}

func NewMockAutoscalingV1Interface(ctrl *gomock.Controller) *MockAutoscalingV1Interface {
	mock := &MockAutoscalingV1Interface{ctrl: ctrl}
	mock.recorder = &MockAutoscalingV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAutoscalingV1Interface) EXPECT() *MockAutoscalingV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAutoscalingV1Interface) HorizontalPodAutoscalers(arg0 string) v1.HorizontalPodAutoscalerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.HorizontalPodAutoscalers, arg0)
	ret0, _ := ret[0].(v1.HorizontalPodAutoscalerInterface)
	return ret0
}

func (mr *MockAutoscalingV1InterfaceMockRecorder) HorizontalPodAutoscalers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.HorizontalPodAutoscalers, reflect.TypeOf((*MockAutoscalingV1Interface)(nil).HorizontalPodAutoscalers), arg0)
}

func (m *MockAutoscalingV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAutoscalingV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAutoscalingV1Interface)(nil).RESTClient))
}
