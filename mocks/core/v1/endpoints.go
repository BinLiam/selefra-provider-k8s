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

type MockEndpointsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointsGetterMockRecorder
}

type MockEndpointsGetterMockRecorder struct {
	mock *MockEndpointsGetter
}

func NewMockEndpointsGetter(ctrl *gomock.Controller) *MockEndpointsGetter {
	mock := &MockEndpointsGetter{ctrl: ctrl}
	mock.recorder = &MockEndpointsGetterMockRecorder{mock}
	return mock
}

func (m *MockEndpointsGetter) EXPECT() *MockEndpointsGetterMockRecorder {
	return m.recorder
}

func (m *MockEndpointsGetter) Endpoints(arg0 string) v12.EndpointsInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Endpoints, arg0)
	ret0, _ := ret[0].(v12.EndpointsInterface)
	return ret0
}

func (mr *MockEndpointsGetterMockRecorder) Endpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Endpoints, reflect.TypeOf((*MockEndpointsGetter)(nil).Endpoints), arg0)
}

type MockEndpointsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointsInterfaceMockRecorder
}

type MockEndpointsInterfaceMockRecorder struct {
	mock *MockEndpointsInterface
}

func NewMockEndpointsInterface(ctrl *gomock.Controller) *MockEndpointsInterface {
	mock := &MockEndpointsInterface{ctrl: ctrl}
	mock.recorder = &MockEndpointsInterfaceMockRecorder{mock}
	return mock
}

func (m *MockEndpointsInterface) EXPECT() *MockEndpointsInterfaceMockRecorder {
	return m.recorder
}

func (m *MockEndpointsInterface) Apply(arg0 context.Context, arg1 *v11.EndpointsApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockEndpointsInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) Create(arg0 context.Context, arg1 *v1.Endpoints, arg2 v10.CreateOptions) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockEndpointsInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEndpointsInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockEndpointsInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEndpointsInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockEndpointsInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockEndpointsInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.EndpointsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.EndpointsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockEndpointsInterface)(nil).List), arg0, arg1)
}

func (m *MockEndpointsInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockEndpointsInterface)(nil).Patch), varargs...)
}

func (m *MockEndpointsInterface) Update(arg0 context.Context, arg1 *v1.Endpoints, arg2 v10.UpdateOptions) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockEndpointsInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockEndpointsInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointsInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockEndpointsInterface)(nil).Watch), arg0, arg1)
}
