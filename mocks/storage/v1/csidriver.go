package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/storage/v1"
	v12 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

type MockCSIDriversGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockCSIDriversGetterMockRecorder
}

type MockCSIDriversGetterMockRecorder struct {
	mock *MockCSIDriversGetter
}

func NewMockCSIDriversGetter(ctrl *gomock.Controller) *MockCSIDriversGetter {
	mock := &MockCSIDriversGetter{ctrl: ctrl}
	mock.recorder = &MockCSIDriversGetterMockRecorder{mock}
	return mock
}

func (m *MockCSIDriversGetter) EXPECT() *MockCSIDriversGetterMockRecorder {
	return m.recorder
}

func (m *MockCSIDriversGetter) CSIDrivers() v12.CSIDriverInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSIDrivers)
	ret0, _ := ret[0].(v12.CSIDriverInterface)
	return ret0
}

func (mr *MockCSIDriversGetterMockRecorder) CSIDrivers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSIDrivers, reflect.TypeOf((*MockCSIDriversGetter)(nil).CSIDrivers))
}

type MockCSIDriverInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockCSIDriverInterfaceMockRecorder
}

type MockCSIDriverInterfaceMockRecorder struct {
	mock *MockCSIDriverInterface
}

func NewMockCSIDriverInterface(ctrl *gomock.Controller) *MockCSIDriverInterface {
	mock := &MockCSIDriverInterface{ctrl: ctrl}
	mock.recorder = &MockCSIDriverInterfaceMockRecorder{mock}
	return mock
}

func (m *MockCSIDriverInterface) EXPECT() *MockCSIDriverInterfaceMockRecorder {
	return m.recorder
}

func (m *MockCSIDriverInterface) Apply(arg0 context.Context, arg1 *v11.CSIDriverApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CSIDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockCSIDriverInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) Create(arg0 context.Context, arg1 *v1.CSIDriver, arg2 v10.CreateOptions) (*v1.CSIDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockCSIDriverInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSIDriverInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockCSIDriverInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSIDriverInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockCSIDriverInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CSIDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockCSIDriverInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CSIDriverList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.CSIDriverList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCSIDriverInterface)(nil).List), arg0, arg1)
}

func (m *MockCSIDriverInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CSIDriver, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.CSIDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockCSIDriverInterface)(nil).Patch), varargs...)
}

func (m *MockCSIDriverInterface) Update(arg0 context.Context, arg1 *v1.CSIDriver, arg2 v10.UpdateOptions) (*v1.CSIDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockCSIDriverInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockCSIDriverInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSIDriverInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockCSIDriverInterface)(nil).Watch), arg0, arg1)
}
