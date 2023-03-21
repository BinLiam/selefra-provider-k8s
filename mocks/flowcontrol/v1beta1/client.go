package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockFlowcontrolV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockFlowcontrolV1beta1InterfaceMockRecorder
}

type MockFlowcontrolV1beta1InterfaceMockRecorder struct {
	mock *MockFlowcontrolV1beta1Interface
}

func NewMockFlowcontrolV1beta1Interface(ctrl *gomock.Controller) *MockFlowcontrolV1beta1Interface {
	mock := &MockFlowcontrolV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockFlowcontrolV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockFlowcontrolV1beta1Interface) EXPECT() *MockFlowcontrolV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockFlowcontrolV1beta1Interface) FlowSchemas() v1beta1.FlowSchemaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.FlowSchemas)
	ret0, _ := ret[0].(v1beta1.FlowSchemaInterface)
	return ret0
}

func (mr *MockFlowcontrolV1beta1InterfaceMockRecorder) FlowSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.FlowSchemas, reflect.TypeOf((*MockFlowcontrolV1beta1Interface)(nil).FlowSchemas))
}

func (m *MockFlowcontrolV1beta1Interface) PriorityLevelConfigurations() v1beta1.PriorityLevelConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PriorityLevelConfigurations)
	ret0, _ := ret[0].(v1beta1.PriorityLevelConfigurationInterface)
	return ret0
}

func (mr *MockFlowcontrolV1beta1InterfaceMockRecorder) PriorityLevelConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PriorityLevelConfigurations, reflect.TypeOf((*MockFlowcontrolV1beta1Interface)(nil).PriorityLevelConfigurations))
}

func (m *MockFlowcontrolV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockFlowcontrolV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockFlowcontrolV1beta1Interface)(nil).RESTClient))
}
