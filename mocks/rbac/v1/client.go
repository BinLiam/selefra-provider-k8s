package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rest "k8s.io/client-go/rest"
)

type MockRbacV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockRbacV1InterfaceMockRecorder
}

type MockRbacV1InterfaceMockRecorder struct {
	mock *MockRbacV1Interface
}

func NewMockRbacV1Interface(ctrl *gomock.Controller) *MockRbacV1Interface {
	mock := &MockRbacV1Interface{ctrl: ctrl}
	mock.recorder = &MockRbacV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockRbacV1Interface) EXPECT() *MockRbacV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockRbacV1Interface) ClusterRoleBindings() v1.ClusterRoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ClusterRoleBindings)
	ret0, _ := ret[0].(v1.ClusterRoleBindingInterface)
	return ret0
}

func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ClusterRoleBindings, reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoleBindings))
}

func (m *MockRbacV1Interface) ClusterRoles() v1.ClusterRoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ClusterRoles)
	ret0, _ := ret[0].(v1.ClusterRoleInterface)
	return ret0
}

func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ClusterRoles, reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoles))
}

func (m *MockRbacV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockRbacV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockRbacV1Interface)(nil).RESTClient))
}

func (m *MockRbacV1Interface) RoleBindings(arg0 string) v1.RoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RoleBindings, arg0)
	ret0, _ := ret[0].(v1.RoleBindingInterface)
	return ret0
}

func (mr *MockRbacV1InterfaceMockRecorder) RoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RoleBindings, reflect.TypeOf((*MockRbacV1Interface)(nil).RoleBindings), arg0)
}

func (m *MockRbacV1Interface) Roles(arg0 string) v1.RoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Roles, arg0)
	ret0, _ := ret[0].(v1.RoleInterface)
	return ret0
}

func (mr *MockRbacV1InterfaceMockRecorder) Roles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Roles, reflect.TypeOf((*MockRbacV1Interface)(nil).Roles), arg0)
}
