package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fields "k8s.io/apimachinery/pkg/fields"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type MockEventsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEventsGetterMockRecorder
}

type MockEventsGetterMockRecorder struct {
	mock *MockEventsGetter
}

func NewMockEventsGetter(ctrl *gomock.Controller) *MockEventsGetter {
	mock := &MockEventsGetter{ctrl: ctrl}
	mock.recorder = &MockEventsGetterMockRecorder{mock}
	return mock
}

func (m *MockEventsGetter) EXPECT() *MockEventsGetterMockRecorder {
	return m.recorder
}

func (m *MockEventsGetter) Events(arg0 string) v12.EventInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Events, arg0)
	ret0, _ := ret[0].(v12.EventInterface)
	return ret0
}

func (mr *MockEventsGetterMockRecorder) Events(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Events, reflect.TypeOf((*MockEventsGetter)(nil).Events), arg0)
}

type MockEventInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEventInterfaceMockRecorder
}

type MockEventInterfaceMockRecorder struct {
	mock *MockEventInterface
}

func NewMockEventInterface(ctrl *gomock.Controller) *MockEventInterface {
	mock := &MockEventInterface{ctrl: ctrl}
	mock.recorder = &MockEventInterfaceMockRecorder{mock}
	return mock
}

func (m *MockEventInterface) EXPECT() *MockEventInterfaceMockRecorder {
	return m.recorder
}

func (m *MockEventInterface) Apply(arg0 context.Context, arg1 *v11.EventApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockEventInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockEventInterface) Create(arg0 context.Context, arg1 *v1.Event, arg2 v10.CreateOptions) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockEventInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockEventInterface) CreateWithEventNamespace(arg0 *v1.Event) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CreateWithEventNamespace, arg0)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) CreateWithEventNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CreateWithEventNamespace, reflect.TypeOf((*MockEventInterface)(nil).CreateWithEventNamespace), arg0)
}

func (m *MockEventInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockEventInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockEventInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockEventInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockEventInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockEventInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockEventInterface) GetFieldSelector(arg0, arg1, arg2, arg3 *string) fields.Selector {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetFieldSelector, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(fields.Selector)
	return ret0
}

func (mr *MockEventInterfaceMockRecorder) GetFieldSelector(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFieldSelector, reflect.TypeOf((*MockEventInterface)(nil).GetFieldSelector), arg0, arg1, arg2, arg3)
}

func (m *MockEventInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.EventList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.EventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockEventInterface)(nil).List), arg0, arg1)
}

func (m *MockEventInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockEventInterface)(nil).Patch), varargs...)
}

func (m *MockEventInterface) PatchWithEventNamespace(arg0 *v1.Event, arg1 []byte) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PatchWithEventNamespace, arg0, arg1)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) PatchWithEventNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PatchWithEventNamespace, reflect.TypeOf((*MockEventInterface)(nil).PatchWithEventNamespace), arg0, arg1)
}

func (m *MockEventInterface) Search(arg0 *runtime.Scheme, arg1 runtime.Object) (*v1.EventList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Search, arg0, arg1)
	ret0, _ := ret[0].(*v1.EventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Search, reflect.TypeOf((*MockEventInterface)(nil).Search), arg0, arg1)
}

func (m *MockEventInterface) Update(arg0 context.Context, arg1 *v1.Event, arg2 v10.UpdateOptions) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockEventInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockEventInterface) UpdateWithEventNamespace(arg0 *v1.Event) (*v1.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateWithEventNamespace, arg0)
	ret0, _ := ret[0].(*v1.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) UpdateWithEventNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateWithEventNamespace, reflect.TypeOf((*MockEventInterface)(nil).UpdateWithEventNamespace), arg0)
}

func (m *MockEventInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockEventInterface)(nil).Watch), arg0, arg1)
}
