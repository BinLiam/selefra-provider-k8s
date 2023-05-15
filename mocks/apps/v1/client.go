package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	rest "k8s.io/client-go/rest"
)

type MockAppsV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAppsV1InterfaceMockRecorder
}

type MockAppsV1InterfaceMockRecorder struct {
	mock *MockAppsV1Interface
}

func NewMockAppsV1Interface(ctrl *gomock.Controller) *MockAppsV1Interface {
	mock := &MockAppsV1Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAppsV1Interface) EXPECT() *MockAppsV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAppsV1Interface) ControllerRevisions(arg0 string) v1.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ControllerRevisions, arg0)
	ret0, _ := ret[0].(v1.ControllerRevisionInterface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ControllerRevisions, reflect.TypeOf((*MockAppsV1Interface)(nil).ControllerRevisions), arg0)
}

func (m *MockAppsV1Interface) DaemonSets(arg0 string) v1.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DaemonSets, arg0)
	ret0, _ := ret[0].(v1.DaemonSetInterface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DaemonSets, reflect.TypeOf((*MockAppsV1Interface)(nil).DaemonSets), arg0)
}

func (m *MockAppsV1Interface) Deployments(arg0 string) v1.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Deployments, arg0)
	ret0, _ := ret[0].(v1.DeploymentInterface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Deployments, reflect.TypeOf((*MockAppsV1Interface)(nil).Deployments), arg0)
}

func (m *MockAppsV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAppsV1Interface)(nil).RESTClient))
}

func (m *MockAppsV1Interface) ReplicaSets(arg0 string) v1.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicaSets, arg0)
	ret0, _ := ret[0].(v1.ReplicaSetInterface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicaSets, reflect.TypeOf((*MockAppsV1Interface)(nil).ReplicaSets), arg0)
}

func (m *MockAppsV1Interface) StatefulSets(arg0 string) v1.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StatefulSets, arg0)
	ret0, _ := ret[0].(v1.StatefulSetInterface)
	return ret0
}

func (mr *MockAppsV1InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StatefulSets, reflect.TypeOf((*MockAppsV1Interface)(nil).StatefulSets), arg0)
}
