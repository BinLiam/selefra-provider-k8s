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

type MockPersistentVolumesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPersistentVolumesGetterMockRecorder
}

type MockPersistentVolumesGetterMockRecorder struct {
	mock *MockPersistentVolumesGetter
}

func NewMockPersistentVolumesGetter(ctrl *gomock.Controller) *MockPersistentVolumesGetter {
	mock := &MockPersistentVolumesGetter{ctrl: ctrl}
	mock.recorder = &MockPersistentVolumesGetterMockRecorder{mock}
	return mock
}

func (m *MockPersistentVolumesGetter) EXPECT() *MockPersistentVolumesGetterMockRecorder {
	return m.recorder
}

func (m *MockPersistentVolumesGetter) PersistentVolumes() v12.PersistentVolumeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PersistentVolumes)
	ret0, _ := ret[0].(v12.PersistentVolumeInterface)
	return ret0
}

func (mr *MockPersistentVolumesGetterMockRecorder) PersistentVolumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PersistentVolumes, reflect.TypeOf((*MockPersistentVolumesGetter)(nil).PersistentVolumes))
}

type MockPersistentVolumeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersistentVolumeInterfaceMockRecorder
}

type MockPersistentVolumeInterfaceMockRecorder struct {
	mock *MockPersistentVolumeInterface
}

func NewMockPersistentVolumeInterface(ctrl *gomock.Controller) *MockPersistentVolumeInterface {
	mock := &MockPersistentVolumeInterface{ctrl: ctrl}
	mock.recorder = &MockPersistentVolumeInterfaceMockRecorder{mock}
	return mock
}

func (m *MockPersistentVolumeInterface) EXPECT() *MockPersistentVolumeInterfaceMockRecorder {
	return m.recorder
}

func (m *MockPersistentVolumeInterface) Apply(arg0 context.Context, arg1 *v11.PersistentVolumeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) ApplyStatus(arg0 context.Context, arg1 *v11.PersistentVolumeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) Create(arg0 context.Context, arg1 *v1.PersistentVolume, arg2 v10.CreateOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PersistentVolumeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.PersistentVolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).List), arg0, arg1)
}

func (m *MockPersistentVolumeInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Patch), varargs...)
}

func (m *MockPersistentVolumeInterface) Update(arg0 context.Context, arg1 *v1.PersistentVolume, arg2 v10.UpdateOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) UpdateStatus(arg0 context.Context, arg1 *v1.PersistentVolume, arg2 v10.UpdateOptions) (*v1.PersistentVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockPersistentVolumeInterface)(nil).Watch), arg0, arg1)
}
