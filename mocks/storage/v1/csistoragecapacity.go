package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/storage/v1"
	v12 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

type MockCSIStorageCapacitiesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCSIStorageCapacitiesGetterMockRecorder
}

type MockCSIStorageCapacitiesGetterMockRecorder struct {
	mock *MockCSIStorageCapacitiesGetter
}

func NewMockCSIStorageCapacitiesGetter(ctrl *gomock.Controller) *MockCSIStorageCapacitiesGetter {
	mock := &MockCSIStorageCapacitiesGetter{ctrl: ctrl}
	mock.recorder = &MockCSIStorageCapacitiesGetterMockRecorder{mock}
	return mock
}

func (m *MockCSIStorageCapacitiesGetter) EXPECT() *MockCSIStorageCapacitiesGetterMockRecorder {
	return m.recorder
}

func (m *MockCSIStorageCapacitiesGetter) CSIStorageCapacities(arg0 string) v12.CSIStorageCapacityInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSIStorageCapacities, arg0)
	ret0, _ := ret[0].(v12.CSIStorageCapacityInterface)
	return ret0
}

func (mr *MockCSIStorageCapacitiesGetterMockRecorder) CSIStorageCapacities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSIStorageCapacities, reflect.TypeOf((*MockCSIStorageCapacitiesGetter)(nil).CSIStorageCapacities), arg0)
}

type MockCSIStorageCapacityInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCSIStorageCapacityInterfaceMockRecorder
}

type MockCSIStorageCapacityInterfaceMockRecorder struct {
	mock *MockCSIStorageCapacityInterface
}

func NewMockCSIStorageCapacityInterface(ctrl *gomock.Controller) *MockCSIStorageCapacityInterface {
	mock := &MockCSIStorageCapacityInterface{ctrl: ctrl}
	mock.recorder = &MockCSIStorageCapacityInterfaceMockRecorder{mock}
	return mock
}

func (m *MockCSIStorageCapacityInterface) EXPECT() *MockCSIStorageCapacityInterfaceMockRecorder {
	return m.recorder
}

func (m *MockCSIStorageCapacityInterface) Apply(arg0 context.Context, arg1 *v11.CSIStorageCapacityApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) Create(arg0 context.Context, arg1 *v1.CSIStorageCapacity, arg2 v10.CreateOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CSIStorageCapacityList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.CSIStorageCapacityList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).List), arg0, arg1)
}

func (m *MockCSIStorageCapacityInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Patch), varargs...)
}

func (m *MockCSIStorageCapacityInterface) Update(arg0 context.Context, arg1 *v1.CSIStorageCapacity, arg2 v10.UpdateOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockCSIStorageCapacityInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Watch), arg0, arg1)
}
