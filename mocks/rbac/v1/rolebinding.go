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

type MockRoleBindingsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingsGetterMockRecorder
}

type MockRoleBindingsGetterMockRecorder struct {
	mock *MockRoleBindingsGetter
}

func NewMockRoleBindingsGetter(ctrl *gomock.Controller) *MockRoleBindingsGetter {
	mock := &MockRoleBindingsGetter{ctrl: ctrl}
	mock.recorder = &MockRoleBindingsGetterMockRecorder{mock}
	return mock
}

func (m *MockRoleBindingsGetter) EXPECT() *MockRoleBindingsGetterMockRecorder {
	return m.recorder
}

func (m *MockRoleBindingsGetter) RoleBindings(arg0 string) v12.RoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RoleBindings, arg0)
	ret0, _ := ret[0].(v12.RoleBindingInterface)
	return ret0
}

func (mr *MockRoleBindingsGetterMockRecorder) RoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RoleBindings, reflect.TypeOf((*MockRoleBindingsGetter)(nil).RoleBindings), arg0)
}

type MockRoleBindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingInterfaceMockRecorder
}

type MockRoleBindingInterfaceMockRecorder struct {
	mock *MockRoleBindingInterface
}

func NewMockRoleBindingInterface(ctrl *gomock.Controller) *MockRoleBindingInterface {
	mock := &MockRoleBindingInterface{ctrl: ctrl}
	mock.recorder = &MockRoleBindingInterfaceMockRecorder{mock}
	return mock
}

func (m *MockRoleBindingInterface) EXPECT() *MockRoleBindingInterfaceMockRecorder {
	return m.recorder
}

func (m *MockRoleBindingInterface) Apply(arg0 context.Context, arg1 *v11.RoleBindingApplyConfiguration, arg2 v10.ApplyOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockRoleBindingInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) Create(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.CreateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockRoleBindingInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRoleBindingInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockRoleBindingInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRoleBindingInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockRoleBindingInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockRoleBindingInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockRoleBindingInterface)(nil).List), arg0, arg1)
}

func (m *MockRoleBindingInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockRoleBindingInterface)(nil).Patch), varargs...)
}

func (m *MockRoleBindingInterface) Update(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.UpdateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockRoleBindingInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockRoleBindingInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoleBindingInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockRoleBindingInterface)(nil).Watch), arg0, arg1)
}
