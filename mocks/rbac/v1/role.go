package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/rbac/v1"
	v12 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type MockRolesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockRolesGetterMockRecorder
}

type MockRolesGetterMockRecorder struct {
	mock *MockRolesGetter
}

func NewMockRolesGetter(ctrl *gomock.Controller) *MockRolesGetter {
	mock := &MockRolesGetter{ctrl: ctrl}
	mock.recorder = &MockRolesGetterMockRecorder{mock}
	return mock
}

func (m *MockRolesGetter) EXPECT() *MockRolesGetterMockRecorder {
	return m.recorder
}

func (m *MockRolesGetter) Roles(arg0 string) v12.RoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Roles, arg0)
	ret0, _ := ret[0].(v12.RoleInterface)
	return ret0
}

func (mr *MockRolesGetterMockRecorder) Roles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Roles, reflect.TypeOf((*MockRolesGetter)(nil).Roles), arg0)
}

type MockRoleInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockRoleInterfaceMockRecorder
}

type MockRoleInterfaceMockRecorder struct {
	mock *MockRoleInterface
}

func NewMockRoleInterface(ctrl *gomock.Controller) *MockRoleInterface {
	mock := &MockRoleInterface{ctrl: ctrl}
	mock.recorder = &MockRoleInterfaceMockRecorder{mock}
	return mock
}

func (m *MockRoleInterface) EXPECT() *MockRoleInterfaceMockRecorder {
	return m.recorder
}

func (m *MockRoleInterface) Apply(arg0 context.Context, arg1 *v11.RoleApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockRoleInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockRoleInterface) Create(arg0 context.Context, arg1 *v1.Role, arg2 v10.CreateOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockRoleInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockRoleInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRoleInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockRoleInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockRoleInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRoleInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockRoleInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockRoleInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockRoleInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockRoleInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockRoleInterface)(nil).List), arg0, arg1)
}

func (m *MockRoleInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Role, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockRoleInterface)(nil).Patch), varargs...)
}

func (m *MockRoleInterface) Update(arg0 context.Context, arg1 *v1.Role, arg2 v10.UpdateOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockRoleInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockRoleInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockRoleInterface)(nil).Watch), arg0, arg1)
}
