package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type MockFlowcontrolV1alpha1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockFlowcontrolV1alpha1InterfaceMockRecorder
}

type MockFlowcontrolV1alpha1InterfaceMockRecorder struct {
	mock *MockFlowcontrolV1alpha1Interface
}

func NewMockFlowcontrolV1alpha1Interface(ctrl *gomock.Controller) *MockFlowcontrolV1alpha1Interface {
	mock := &MockFlowcontrolV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockFlowcontrolV1alpha1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockFlowcontrolV1alpha1Interface) EXPECT() *MockFlowcontrolV1alpha1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockFlowcontrolV1alpha1Interface) FlowSchemas() v1alpha1.FlowSchemaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.FlowSchemas)
	ret0, _ := ret[0].(v1alpha1.FlowSchemaInterface)
	return ret0
}

func (mr *MockFlowcontrolV1alpha1InterfaceMockRecorder) FlowSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.FlowSchemas, reflect.TypeOf((*MockFlowcontrolV1alpha1Interface)(nil).FlowSchemas))
}

func (m *MockFlowcontrolV1alpha1Interface) PriorityLevelConfigurations() v1alpha1.PriorityLevelConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PriorityLevelConfigurations)
	ret0, _ := ret[0].(v1alpha1.PriorityLevelConfigurationInterface)
	return ret0
}

func (mr *MockFlowcontrolV1alpha1InterfaceMockRecorder) PriorityLevelConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PriorityLevelConfigurations, reflect.TypeOf((*MockFlowcontrolV1alpha1Interface)(nil).PriorityLevelConfigurations))
}

func (m *MockFlowcontrolV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockFlowcontrolV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockFlowcontrolV1alpha1Interface)(nil).RESTClient))
}
