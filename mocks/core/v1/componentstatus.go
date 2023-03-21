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

type MockComponentStatusesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockComponentStatusesGetterMockRecorder
}

type MockComponentStatusesGetterMockRecorder struct {
	mock *MockComponentStatusesGetter
}

func NewMockComponentStatusesGetter(ctrl *gomock.Controller) *MockComponentStatusesGetter {
	mock := &MockComponentStatusesGetter{ctrl: ctrl}
	mock.recorder = &MockComponentStatusesGetterMockRecorder{mock}
	return mock
}

func (m *MockComponentStatusesGetter) EXPECT() *MockComponentStatusesGetterMockRecorder {
	return m.recorder
}

func (m *MockComponentStatusesGetter) ComponentStatuses() v12.ComponentStatusInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ComponentStatuses)
	ret0, _ := ret[0].(v12.ComponentStatusInterface)
	return ret0
}

func (mr *MockComponentStatusesGetterMockRecorder) ComponentStatuses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ComponentStatuses, reflect.TypeOf((*MockComponentStatusesGetter)(nil).ComponentStatuses))
}

type MockComponentStatusInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockComponentStatusInterfaceMockRecorder
}

type MockComponentStatusInterfaceMockRecorder struct {
	mock *MockComponentStatusInterface
}

func NewMockComponentStatusInterface(ctrl *gomock.Controller) *MockComponentStatusInterface {
	mock := &MockComponentStatusInterface{ctrl: ctrl}
	mock.recorder = &MockComponentStatusInterfaceMockRecorder{mock}
	return mock
}

func (m *MockComponentStatusInterface) EXPECT() *MockComponentStatusInterfaceMockRecorder {
	return m.recorder
}

func (m *MockComponentStatusInterface) Apply(arg0 context.Context, arg1 *v11.ComponentStatusApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockComponentStatusInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) Create(arg0 context.Context, arg1 *v1.ComponentStatus, arg2 v10.CreateOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockComponentStatusInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockComponentStatusInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockComponentStatusInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockComponentStatusInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockComponentStatusInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockComponentStatusInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ComponentStatusList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ComponentStatusList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockComponentStatusInterface)(nil).List), arg0, arg1)
}

func (m *MockComponentStatusInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockComponentStatusInterface)(nil).Patch), varargs...)
}

func (m *MockComponentStatusInterface) Update(arg0 context.Context, arg1 *v1.ComponentStatus, arg2 v10.UpdateOptions) (*v1.ComponentStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ComponentStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockComponentStatusInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockComponentStatusInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComponentStatusInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockComponentStatusInterface)(nil).Watch), arg0, arg1)
}
