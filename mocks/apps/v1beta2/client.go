package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	rest "k8s.io/client-go/rest"
)

type MockAppsV1beta2Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAppsV1beta2InterfaceMockRecorder
}

type MockAppsV1beta2InterfaceMockRecorder struct {
	mock *MockAppsV1beta2Interface
}

func NewMockAppsV1beta2Interface(ctrl *gomock.Controller) *MockAppsV1beta2Interface {
	mock := &MockAppsV1beta2Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1beta2InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAppsV1beta2Interface) EXPECT() *MockAppsV1beta2InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAppsV1beta2Interface) ControllerRevisions(arg0 string) v1beta2.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ControllerRevisions, arg0)
	ret0, _ := ret[0].(v1beta2.ControllerRevisionInterface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ControllerRevisions, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).ControllerRevisions), arg0)
}

func (m *MockAppsV1beta2Interface) DaemonSets(arg0 string) v1beta2.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DaemonSets, arg0)
	ret0, _ := ret[0].(v1beta2.DaemonSetInterface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DaemonSets, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).DaemonSets), arg0)
}

func (m *MockAppsV1beta2Interface) Deployments(arg0 string) v1beta2.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Deployments, arg0)
	ret0, _ := ret[0].(v1beta2.DeploymentInterface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Deployments, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).Deployments), arg0)
}

func (m *MockAppsV1beta2Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).RESTClient))
}

func (m *MockAppsV1beta2Interface) ReplicaSets(arg0 string) v1beta2.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicaSets, arg0)
	ret0, _ := ret[0].(v1beta2.ReplicaSetInterface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicaSets, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).ReplicaSets), arg0)
}

func (m *MockAppsV1beta2Interface) StatefulSets(arg0 string) v1beta2.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StatefulSets, arg0)
	ret0, _ := ret[0].(v1beta2.StatefulSetInterface)
	return ret0
}

func (mr *MockAppsV1beta2InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StatefulSets, reflect.TypeOf((*MockAppsV1beta2Interface)(nil).StatefulSets), arg0)
}
