package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockExtensionsV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockExtensionsV1beta1InterfaceMockRecorder
}

type MockExtensionsV1beta1InterfaceMockRecorder struct {
	mock *MockExtensionsV1beta1Interface
}

func NewMockExtensionsV1beta1Interface(ctrl *gomock.Controller) *MockExtensionsV1beta1Interface {
	mock := &MockExtensionsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockExtensionsV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockExtensionsV1beta1Interface) EXPECT() *MockExtensionsV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockExtensionsV1beta1Interface) DaemonSets(arg0 string) v1beta1.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DaemonSets, arg0)
	ret0, _ := ret[0].(v1beta1.DaemonSetInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DaemonSets, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).DaemonSets), arg0)
}

func (m *MockExtensionsV1beta1Interface) Deployments(arg0 string) v1beta1.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Deployments, arg0)
	ret0, _ := ret[0].(v1beta1.DeploymentInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Deployments, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).Deployments), arg0)
}

func (m *MockExtensionsV1beta1Interface) Ingresses(arg0 string) v1beta1.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Ingresses, arg0)
	ret0, _ := ret[0].(v1beta1.IngressInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Ingresses, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).Ingresses), arg0)
}

func (m *MockExtensionsV1beta1Interface) NetworkPolicies(arg0 string) v1beta1.NetworkPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkPolicies, arg0)
	ret0, _ := ret[0].(v1beta1.NetworkPolicyInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) NetworkPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkPolicies, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).NetworkPolicies), arg0)
}

func (m *MockExtensionsV1beta1Interface) PodSecurityPolicies() v1beta1.PodSecurityPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PodSecurityPolicies)
	ret0, _ := ret[0].(v1beta1.PodSecurityPolicyInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) PodSecurityPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PodSecurityPolicies, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).PodSecurityPolicies))
}

func (m *MockExtensionsV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).RESTClient))
}

func (m *MockExtensionsV1beta1Interface) ReplicaSets(arg0 string) v1beta1.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicaSets, arg0)
	ret0, _ := ret[0].(v1beta1.ReplicaSetInterface)
	return ret0
}

func (mr *MockExtensionsV1beta1InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicaSets, reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).ReplicaSets), arg0)
}
