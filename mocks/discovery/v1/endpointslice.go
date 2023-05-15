package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/discovery/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/discovery/v1"
	v12 "k8s.io/client-go/kubernetes/typed/discovery/v1"
)

type MockEndpointSlicesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointSlicesGetterMockRecorder
}

type MockEndpointSlicesGetterMockRecorder struct {
	mock *MockEndpointSlicesGetter
}

func NewMockEndpointSlicesGetter(ctrl *gomock.Controller) *MockEndpointSlicesGetter {
	mock := &MockEndpointSlicesGetter{ctrl: ctrl}
	mock.recorder = &MockEndpointSlicesGetterMockRecorder{mock}
	return mock
}

func (m *MockEndpointSlicesGetter) EXPECT() *MockEndpointSlicesGetterMockRecorder {
	return m.recorder
}

func (m *MockEndpointSlicesGetter) EndpointSlices(arg0 string) v12.EndpointSliceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EndpointSlices, arg0)
	ret0, _ := ret[0].(v12.EndpointSliceInterface)
	return ret0
}

func (mr *MockEndpointSlicesGetterMockRecorder) EndpointSlices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EndpointSlices, reflect.TypeOf((*MockEndpointSlicesGetter)(nil).EndpointSlices), arg0)
}

type MockEndpointSliceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointSliceInterfaceMockRecorder
}

type MockEndpointSliceInterfaceMockRecorder struct {
	mock *MockEndpointSliceInterface
}

func NewMockEndpointSliceInterface(ctrl *gomock.Controller) *MockEndpointSliceInterface {
	mock := &MockEndpointSliceInterface{ctrl: ctrl}
	mock.recorder = &MockEndpointSliceInterfaceMockRecorder{mock}
	return mock
}

func (m *MockEndpointSliceInterface) EXPECT() *MockEndpointSliceInterfaceMockRecorder {
	return m.recorder
}

func (m *MockEndpointSliceInterface) Apply(arg0 context.Context, arg1 *v11.EndpointSliceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) Create(arg0 context.Context, arg1 *v1.EndpointSlice, arg2 v10.CreateOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEndpointSliceInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockEndpointSliceInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.EndpointSliceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.EndpointSliceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockEndpointSliceInterface)(nil).List), arg0, arg1)
}

func (m *MockEndpointSliceInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Patch), varargs...)
}

func (m *MockEndpointSliceInterface) Update(arg0 context.Context, arg1 *v1.EndpointSlice, arg2 v10.UpdateOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockEndpointSliceInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEndpointSliceInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockEndpointSliceInterface)(nil).Watch), arg0, arg1)
}
