package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	rest "k8s.io/client-go/rest"
)

type MockAutoscalingV2beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalingV2beta1InterfaceMockRecorder
}

type MockAutoscalingV2beta1InterfaceMockRecorder struct {
	mock *MockAutoscalingV2beta1Interface
}

func NewMockAutoscalingV2beta1Interface(ctrl *gomock.Controller) *MockAutoscalingV2beta1Interface {
	mock := &MockAutoscalingV2beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAutoscalingV2beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAutoscalingV2beta1Interface) EXPECT() *MockAutoscalingV2beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAutoscalingV2beta1Interface) HorizontalPodAutoscalers(arg0 string) v2beta1.HorizontalPodAutoscalerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.HorizontalPodAutoscalers, arg0)
	ret0, _ := ret[0].(v2beta1.HorizontalPodAutoscalerInterface)
	return ret0
}

func (mr *MockAutoscalingV2beta1InterfaceMockRecorder) HorizontalPodAutoscalers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.HorizontalPodAutoscalers, reflect.TypeOf((*MockAutoscalingV2beta1Interface)(nil).HorizontalPodAutoscalers), arg0)
}

func (m *MockAutoscalingV2beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAutoscalingV2beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAutoscalingV2beta1Interface)(nil).RESTClient))
}
