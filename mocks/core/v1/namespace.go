package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type MockNamespacesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockNamespacesGetterMockRecorder
}

type MockNamespacesGetterMockRecorder struct {
	mock *MockNamespacesGetter
}

func NewMockNamespacesGetter(ctrl *gomock.Controller) *MockNamespacesGetter {
	mock := &MockNamespacesGetter{ctrl: ctrl}
	mock.recorder = &MockNamespacesGetterMockRecorder{mock}
	return mock
}

func (m *MockNamespacesGetter) EXPECT() *MockNamespacesGetterMockRecorder {
	return m.recorder
}

func (m *MockNamespacesGetter) Namespaces() v12.NamespaceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Namespaces)
	ret0, _ := ret[0].(v12.NamespaceInterface)
	return ret0
}

func (mr *MockNamespacesGetterMockRecorder) Namespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Namespaces, reflect.TypeOf((*MockNamespacesGetter)(nil).Namespaces))
}

type MockNamespaceInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockNamespaceInterfaceMockRecorder
}

type MockNamespaceInterfaceMockRecorder struct {
	mock *MockNamespaceInterface
}

func NewMockNamespaceInterface(ctrl *gomock.Controller) *MockNamespaceInterface {
	mock := &MockNamespaceInterface{ctrl: ctrl}
	mock.recorder = &MockNamespaceInterfaceMockRecorder{mock}
	return mock
}

func (m *MockNamespaceInterface) EXPECT() *MockNamespaceInterfaceMockRecorder {
	return m.recorder
}

func (m *MockNamespaceInterface) Apply(arg0 context.Context, arg1 *v11.NamespaceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockNamespaceInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) ApplyStatus(arg0 context.Context, arg1 *v11.NamespaceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockNamespaceInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) Create(arg0 context.Context, arg1 *v1.Namespace, arg2 v10.CreateOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockNamespaceInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNamespaceInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockNamespaceInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) Finalize(arg0 context.Context, arg1 *v1.Namespace, arg2 v10.UpdateOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Finalize, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Finalize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Finalize, reflect.TypeOf((*MockNamespaceInterface)(nil).Finalize), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockNamespaceInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NamespaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.NamespaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockNamespaceInterface)(nil).List), arg0, arg1)
}

func (m *MockNamespaceInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockNamespaceInterface)(nil).Patch), varargs...)
}

func (m *MockNamespaceInterface) Update(arg0 context.Context, arg1 *v1.Namespace, arg2 v10.UpdateOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockNamespaceInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Namespace, arg2 v10.UpdateOptions) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockNamespaceInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockNamespaceInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNamespaceInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockNamespaceInterface)(nil).Watch), arg0, arg1)
}
