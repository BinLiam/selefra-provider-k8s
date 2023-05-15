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

type MockClusterRoleBindingsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleBindingsGetterMockRecorder
}

type MockClusterRoleBindingsGetterMockRecorder struct {
	mock *MockClusterRoleBindingsGetter
}

func NewMockClusterRoleBindingsGetter(ctrl *gomock.Controller) *MockClusterRoleBindingsGetter {
	mock := &MockClusterRoleBindingsGetter{ctrl: ctrl}
	mock.recorder = &MockClusterRoleBindingsGetterMockRecorder{mock}
	return mock
}

func (m *MockClusterRoleBindingsGetter) EXPECT() *MockClusterRoleBindingsGetterMockRecorder {
	return m.recorder
}

func (m *MockClusterRoleBindingsGetter) ClusterRoleBindings() v12.ClusterRoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ClusterRoleBindings)
	ret0, _ := ret[0].(v12.ClusterRoleBindingInterface)
	return ret0
}

func (mr *MockClusterRoleBindingsGetterMockRecorder) ClusterRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ClusterRoleBindings, reflect.TypeOf((*MockClusterRoleBindingsGetter)(nil).ClusterRoleBindings))
}

type MockClusterRoleBindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleBindingInterfaceMockRecorder
}

type MockClusterRoleBindingInterfaceMockRecorder struct {
	mock *MockClusterRoleBindingInterface
}

func NewMockClusterRoleBindingInterface(ctrl *gomock.Controller) *MockClusterRoleBindingInterface {
	mock := &MockClusterRoleBindingInterface{ctrl: ctrl}
	mock.recorder = &MockClusterRoleBindingInterfaceMockRecorder{mock}
	return mock
}

func (m *MockClusterRoleBindingInterface) EXPECT() *MockClusterRoleBindingInterfaceMockRecorder {
	return m.recorder
}

func (m *MockClusterRoleBindingInterface) Apply(arg0 context.Context, arg1 *v11.ClusterRoleBindingApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) Create(arg0 context.Context, arg1 *v1.ClusterRoleBinding, arg2 v10.CreateOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ClusterRoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ClusterRoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).List), arg0, arg1)
}

func (m *MockClusterRoleBindingInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Patch), varargs...)
}

func (m *MockClusterRoleBindingInterface) Update(arg0 context.Context, arg1 *v1.ClusterRoleBinding, arg2 v10.UpdateOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockClusterRoleBindingInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockClusterRoleBindingInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Watch), arg0, arg1)
}
