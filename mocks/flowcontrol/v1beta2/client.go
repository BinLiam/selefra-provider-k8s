package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	rest "k8s.io/client-go/rest"
)

type MockFlowcontrolV1beta2Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockFlowcontrolV1beta2InterfaceMockRecorder
}

type MockFlowcontrolV1beta2InterfaceMockRecorder struct {
	mock *MockFlowcontrolV1beta2Interface
}

func NewMockFlowcontrolV1beta2Interface(ctrl *gomock.Controller) *MockFlowcontrolV1beta2Interface {
	mock := &MockFlowcontrolV1beta2Interface{ctrl: ctrl}
	mock.recorder = &MockFlowcontrolV1beta2InterfaceMockRecorder{mock}
	return mock
}

func (m *MockFlowcontrolV1beta2Interface) EXPECT() *MockFlowcontrolV1beta2InterfaceMockRecorder {
	return m.recorder
}

func (m *MockFlowcontrolV1beta2Interface) FlowSchemas() v1beta2.FlowSchemaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.FlowSchemas)
	ret0, _ := ret[0].(v1beta2.FlowSchemaInterface)
	return ret0
}

func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) FlowSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.FlowSchemas, reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).FlowSchemas))
}

func (m *MockFlowcontrolV1beta2Interface) PriorityLevelConfigurations() v1beta2.PriorityLevelConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PriorityLevelConfigurations)
	ret0, _ := ret[0].(v1beta2.PriorityLevelConfigurationInterface)
	return ret0
}

func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) PriorityLevelConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PriorityLevelConfigurations, reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).PriorityLevelConfigurations))
}

func (m *MockFlowcontrolV1beta2Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).RESTClient))
}
