package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/rbac/v1"
	v12 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type MockClusterRolesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRolesGetterMockRecorder
}

type MockClusterRolesGetterMockRecorder struct {
	mock *MockClusterRolesGetter
}

func NewMockClusterRolesGetter(ctrl *gomock.Controller) *MockClusterRolesGetter {
	mock := &MockClusterRolesGetter{ctrl: ctrl}
	mock.recorder = &MockClusterRolesGetterMockRecorder{mock}
	return mock
}

func (m *MockClusterRolesGetter) EXPECT() *MockClusterRolesGetterMockRecorder {
	return m.recorder
}

func (m *MockClusterRolesGetter) ClusterRoles() v12.ClusterRoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ClusterRoles)
	ret0, _ := ret[0].(v12.ClusterRoleInterface)
	return ret0
}

func (mr *MockClusterRolesGetterMockRecorder) ClusterRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ClusterRoles, reflect.TypeOf((*MockClusterRolesGetter)(nil).ClusterRoles))
}

type MockClusterRoleInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleInterfaceMockRecorder
}

type MockClusterRoleInterfaceMockRecorder struct {
	mock *MockClusterRoleInterface
}

func NewMockClusterRoleInterface(ctrl *gomock.Controller) *MockClusterRoleInterface {
	mock := &MockClusterRoleInterface{ctrl: ctrl}
	mock.recorder = &MockClusterRoleInterfaceMockRecorder{mock}
	return mock
}

func (m *MockClusterRoleInterface) EXPECT() *MockClusterRoleInterfaceMockRecorder {
	return m.recorder
}

func (m *MockClusterRoleInterface) Apply(arg0 context.Context, arg1 *v11.ClusterRoleApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockClusterRoleInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) Create(arg0 context.Context, arg1 *v1.ClusterRole, arg2 v10.CreateOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockClusterRoleInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClusterRoleInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockClusterRoleInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClusterRoleInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockClusterRoleInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockClusterRoleInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ClusterRoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ClusterRoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockClusterRoleInterface)(nil).List), arg0, arg1)
}

func (m *MockClusterRoleInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockClusterRoleInterface)(nil).Patch), varargs...)
}

func (m *MockClusterRoleInterface) Update(arg0 context.Context, arg1 *v1.ClusterRole, arg2 v10.UpdateOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockClusterRoleInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockClusterRoleInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockClusterRoleInterface)(nil).Watch), arg0, arg1)
}
