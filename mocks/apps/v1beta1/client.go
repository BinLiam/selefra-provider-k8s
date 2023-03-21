package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockAppsV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAppsV1beta1InterfaceMockRecorder
}

type MockAppsV1beta1InterfaceMockRecorder struct {
	mock *MockAppsV1beta1Interface
}

func NewMockAppsV1beta1Interface(ctrl *gomock.Controller) *MockAppsV1beta1Interface {
	mock := &MockAppsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAppsV1beta1Interface) EXPECT() *MockAppsV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAppsV1beta1Interface) ControllerRevisions(arg0 string) v1beta1.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ControllerRevisions, arg0)
	ret0, _ := ret[0].(v1beta1.ControllerRevisionInterface)
	return ret0
}

func (mr *MockAppsV1beta1InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ControllerRevisions, reflect.TypeOf((*MockAppsV1beta1Interface)(nil).ControllerRevisions), arg0)
}

func (m *MockAppsV1beta1Interface) Deployments(arg0 string) v1beta1.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Deployments, arg0)
	ret0, _ := ret[0].(v1beta1.DeploymentInterface)
	return ret0
}

func (mr *MockAppsV1beta1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Deployments, reflect.TypeOf((*MockAppsV1beta1Interface)(nil).Deployments), arg0)
}

func (m *MockAppsV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAppsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAppsV1beta1Interface)(nil).RESTClient))
}

func (m *MockAppsV1beta1Interface) StatefulSets(arg0 string) v1beta1.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StatefulSets, arg0)
	ret0, _ := ret[0].(v1beta1.StatefulSetInterface)
	return ret0
}

func (mr *MockAppsV1beta1InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StatefulSets, reflect.TypeOf((*MockAppsV1beta1Interface)(nil).StatefulSets), arg0)
}
