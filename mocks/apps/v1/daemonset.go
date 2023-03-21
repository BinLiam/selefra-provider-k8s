package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/apps/v1"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type MockDaemonSetsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockDaemonSetsGetterMockRecorder
}

type MockDaemonSetsGetterMockRecorder struct {
	mock *MockDaemonSetsGetter
}

func NewMockDaemonSetsGetter(ctrl *gomock.Controller) *MockDaemonSetsGetter {
	mock := &MockDaemonSetsGetter{ctrl: ctrl}
	mock.recorder = &MockDaemonSetsGetterMockRecorder{mock}
	return mock
}

func (m *MockDaemonSetsGetter) EXPECT() *MockDaemonSetsGetterMockRecorder {
	return m.recorder
}

func (m *MockDaemonSetsGetter) DaemonSets(arg0 string) v12.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DaemonSets, arg0)
	ret0, _ := ret[0].(v12.DaemonSetInterface)
	return ret0
}

func (mr *MockDaemonSetsGetterMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DaemonSets, reflect.TypeOf((*MockDaemonSetsGetter)(nil).DaemonSets), arg0)
}

type MockDaemonSetInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockDaemonSetInterfaceMockRecorder
}

type MockDaemonSetInterfaceMockRecorder struct {
	mock *MockDaemonSetInterface
}

func NewMockDaemonSetInterface(ctrl *gomock.Controller) *MockDaemonSetInterface {
	mock := &MockDaemonSetInterface{ctrl: ctrl}
	mock.recorder = &MockDaemonSetInterfaceMockRecorder{mock}
	return mock
}

func (m *MockDaemonSetInterface) EXPECT() *MockDaemonSetInterfaceMockRecorder {
	return m.recorder
}

func (m *MockDaemonSetInterface) Apply(arg0 context.Context, arg1 *v11.DaemonSetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockDaemonSetInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) ApplyStatus(arg0 context.Context, arg1 *v11.DaemonSetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockDaemonSetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) Create(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.CreateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockDaemonSetInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDaemonSetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockDaemonSetInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDaemonSetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockDaemonSetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockDaemonSetInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.DaemonSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.DaemonSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockDaemonSetInterface)(nil).List), arg0, arg1)
}

func (m *MockDaemonSetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockDaemonSetInterface)(nil).Patch), varargs...)
}

func (m *MockDaemonSetInterface) Update(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.UpdateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockDaemonSetInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.UpdateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockDaemonSetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockDaemonSetInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDaemonSetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockDaemonSetInterface)(nil).Watch), arg0, arg1)
}
