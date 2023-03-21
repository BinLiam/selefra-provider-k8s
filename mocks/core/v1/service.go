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
	rest "k8s.io/client-go/rest"
)

type MockServicesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockServicesGetterMockRecorder
}

type MockServicesGetterMockRecorder struct {
	mock *MockServicesGetter
}

func NewMockServicesGetter(ctrl *gomock.Controller) *MockServicesGetter {
	mock := &MockServicesGetter{ctrl: ctrl}
	mock.recorder = &MockServicesGetterMockRecorder{mock}
	return mock
}

func (m *MockServicesGetter) EXPECT() *MockServicesGetterMockRecorder {
	return m.recorder
}

func (m *MockServicesGetter) Services(arg0 string) v12.ServiceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Services, arg0)
	ret0, _ := ret[0].(v12.ServiceInterface)
	return ret0
}

func (mr *MockServicesGetterMockRecorder) Services(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Services, reflect.TypeOf((*MockServicesGetter)(nil).Services), arg0)
}

type MockServiceInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockServiceInterfaceMockRecorder
}

type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

func (m *MockServiceInterface) Apply(arg0 context.Context, arg1 *v11.ServiceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockServiceInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockServiceInterface) ApplyStatus(arg0 context.Context, arg1 *v11.ServiceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockServiceInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockServiceInterface) Create(arg0 context.Context, arg1 *v1.Service, arg2 v10.CreateOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockServiceInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockServiceInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockServiceInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockServiceInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockServiceInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockServiceInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ServiceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockServiceInterface)(nil).List), arg0, arg1)
}

func (m *MockServiceInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Service, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockServiceInterface)(nil).Patch), varargs...)
}

func (m *MockServiceInterface) ProxyGet(arg0, arg1, arg2, arg3 string, arg4 map[string]string) rest.ResponseWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ProxyGet, arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(rest.ResponseWrapper)
	return ret0
}

func (mr *MockServiceInterfaceMockRecorder) ProxyGet(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ProxyGet, reflect.TypeOf((*MockServiceInterface)(nil).ProxyGet), arg0, arg1, arg2, arg3, arg4)
}

func (m *MockServiceInterface) Update(arg0 context.Context, arg1 *v1.Service, arg2 v10.UpdateOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockServiceInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockServiceInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Service, arg2 v10.UpdateOptions) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockServiceInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockServiceInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockServiceInterface)(nil).Watch), arg0, arg1)
}
