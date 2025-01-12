package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type MockConfigMapsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMapsGetterMockRecorder
}

type MockConfigMapsGetterMockRecorder struct {
	mock *MockConfigMapsGetter
}

func NewMockConfigMapsGetter(ctrl *gomock.Controller) *MockConfigMapsGetter {
	mock := &MockConfigMapsGetter{ctrl: ctrl}
	mock.recorder = &MockConfigMapsGetterMockRecorder{mock}
	return mock
}

func (m *MockConfigMapsGetter) EXPECT() *MockConfigMapsGetterMockRecorder {
	return m.recorder
}

func (m *MockConfigMapsGetter) ConfigMaps(arg0 string) v12.ConfigMapInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ConfigMaps, arg0)
	ret0, _ := ret[0].(v12.ConfigMapInterface)
	return ret0
}

func (mr *MockConfigMapsGetterMockRecorder) ConfigMaps(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ConfigMaps, reflect.TypeOf((*MockConfigMapsGetter)(nil).ConfigMaps), arg0)
}

type MockConfigMapInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMapInterfaceMockRecorder
}

type MockConfigMapInterfaceMockRecorder struct {
	mock *MockConfigMapInterface
}

func NewMockConfigMapInterface(ctrl *gomock.Controller) *MockConfigMapInterface {
	mock := &MockConfigMapInterface{ctrl: ctrl}
	mock.recorder = &MockConfigMapInterfaceMockRecorder{mock}
	return mock
}

func (m *MockConfigMapInterface) EXPECT() *MockConfigMapInterfaceMockRecorder {
	return m.recorder
}

func (m *MockConfigMapInterface) Apply(arg0 context.Context, arg1 *v11.ConfigMapApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockConfigMapInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) Create(arg0 context.Context, arg1 *v1.ConfigMap, arg2 v10.CreateOptions) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockConfigMapInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockConfigMapInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockConfigMapInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockConfigMapInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockConfigMapInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockConfigMapInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ConfigMapList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ConfigMapList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockConfigMapInterface)(nil).List), arg0, arg1)
}

func (m *MockConfigMapInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockConfigMapInterface)(nil).Patch), varargs...)
}

func (m *MockConfigMapInterface) Update(arg0 context.Context, arg1 *v1.ConfigMap, arg2 v10.UpdateOptions) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockConfigMapInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockConfigMapInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockConfigMapInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockConfigMapInterface)(nil).Watch), arg0, arg1)
}
