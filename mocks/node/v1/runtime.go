package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/node/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/node/v1"
	v12 "k8s.io/client-go/kubernetes/typed/node/v1"
)

type MockRuntimeClassesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeClassesGetterMockRecorder
}

type MockRuntimeClassesGetterMockRecorder struct {
	mock *MockRuntimeClassesGetter
}

func NewMockRuntimeClassesGetter(ctrl *gomock.Controller) *MockRuntimeClassesGetter {
	mock := &MockRuntimeClassesGetter{ctrl: ctrl}
	mock.recorder = &MockRuntimeClassesGetterMockRecorder{mock}
	return mock
}

func (m *MockRuntimeClassesGetter) EXPECT() *MockRuntimeClassesGetterMockRecorder {
	return m.recorder
}

func (m *MockRuntimeClassesGetter) RuntimeClasses() v12.RuntimeClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RuntimeClasses)
	ret0, _ := ret[0].(v12.RuntimeClassInterface)
	return ret0
}

func (mr *MockRuntimeClassesGetterMockRecorder) RuntimeClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RuntimeClasses, reflect.TypeOf((*MockRuntimeClassesGetter)(nil).RuntimeClasses))
}

type MockRuntimeClassInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeClassInterfaceMockRecorder
}

type MockRuntimeClassInterfaceMockRecorder struct {
	mock *MockRuntimeClassInterface
}

func NewMockRuntimeClassInterface(ctrl *gomock.Controller) *MockRuntimeClassInterface {
	mock := &MockRuntimeClassInterface{ctrl: ctrl}
	mock.recorder = &MockRuntimeClassInterfaceMockRecorder{mock}
	return mock
}

func (m *MockRuntimeClassInterface) EXPECT() *MockRuntimeClassInterfaceMockRecorder {
	return m.recorder
}

func (m *MockRuntimeClassInterface) Apply(arg0 context.Context, arg1 *v11.RuntimeClassApplyConfiguration, arg2 v10.ApplyOptions) (*v1.RuntimeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RuntimeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) Create(arg0 context.Context, arg1 *v1.RuntimeClass, arg2 v10.CreateOptions) (*v1.RuntimeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RuntimeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRuntimeClassInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockRuntimeClassInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.RuntimeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RuntimeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RuntimeClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.RuntimeClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockRuntimeClassInterface)(nil).List), arg0, arg1)
}

func (m *MockRuntimeClassInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.RuntimeClass, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.RuntimeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Patch), varargs...)
}

func (m *MockRuntimeClassInterface) Update(arg0 context.Context, arg1 *v1.RuntimeClass, arg2 v10.UpdateOptions) (*v1.RuntimeClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RuntimeClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockRuntimeClassInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRuntimeClassInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockRuntimeClassInterface)(nil).Watch), arg0, arg1)
}
