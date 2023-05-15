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

type MockStorageClassesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClassesGetterMockRecorder
}

type MockStorageClassesGetterMockRecorder struct {
	mock *MockStorageClassesGetter
}

func NewMockStorageClassesGetter(ctrl *gomock.Controller) *MockStorageClassesGetter {
	mock := &MockStorageClassesGetter{ctrl: ctrl}
	mock.recorder = &MockStorageClassesGetterMockRecorder{mock}
	return mock
}

func (m *MockStorageClassesGetter) EXPECT() *MockStorageClassesGetterMockRecorder {
	return m.recorder
}

func (m *MockStorageClassesGetter) StorageClasses() v12.StorageClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageClasses)
	ret0, _ := ret[0].(v12.StorageClassInterface)
	return ret0
}

func (mr *MockStorageClassesGetterMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageClasses, reflect.TypeOf((*MockStorageClassesGetter)(nil).StorageClasses))
}

type MockStorageClassInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClassInterfaceMockRecorder
}

type MockStorageClassInterfaceMockRecorder struct {
	mock *MockStorageClassInterface
}

func NewMockStorageClassInterface(ctrl *gomock.Controller) *MockStorageClassInterface {
	mock := &MockStorageClassInterface{ctrl: ctrl}
	mock.recorder = &MockStorageClassInterfaceMockRecorder{mock}
	return mock
}

func (m *MockStorageClassInterface) EXPECT() *MockStorageClassInterfaceMockRecorder {
	return m.recorder
}

func (m *MockStorageClassInterface) Apply(arg0 context.Context, arg1 *v11.StorageClassApplyConfiguration, arg2 v10.ApplyOptions) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockStorageClassInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) Create(arg0 context.Context, arg1 *v1.StorageClass, arg2 v10.CreateOptions) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockStorageClassInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockStorageClassInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockStorageClassInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockStorageClassInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockStorageClassInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockStorageClassInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.StorageClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.StorageClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockStorageClassInterface)(nil).List), arg0, arg1)
}

func (m *MockStorageClassInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockStorageClassInterface)(nil).Patch), varargs...)
}

func (m *MockStorageClassInterface) Update(arg0 context.Context, arg1 *v1.StorageClass, arg2 v10.UpdateOptions) (*v1.StorageClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StorageClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockStorageClassInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockStorageClassInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStorageClassInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockStorageClassInterface)(nil).Watch), arg0, arg1)
}
