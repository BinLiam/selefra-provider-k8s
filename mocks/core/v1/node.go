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

type MockNodesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockNodesGetterMockRecorder
}

type MockNodesGetterMockRecorder struct {
	mock *MockNodesGetter
}

func NewMockNodesGetter(ctrl *gomock.Controller) *MockNodesGetter {
	mock := &MockNodesGetter{ctrl: ctrl}
	mock.recorder = &MockNodesGetterMockRecorder{mock}
	return mock
}

func (m *MockNodesGetter) EXPECT() *MockNodesGetterMockRecorder {
	return m.recorder
}

func (m *MockNodesGetter) Nodes() v12.NodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Nodes)
	ret0, _ := ret[0].(v12.NodeInterface)
	return ret0
}

func (mr *MockNodesGetterMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Nodes, reflect.TypeOf((*MockNodesGetter)(nil).Nodes))
}

type MockNodeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNodeInterfaceMockRecorder
}

type MockNodeInterfaceMockRecorder struct {
	mock *MockNodeInterface
}

func NewMockNodeInterface(ctrl *gomock.Controller) *MockNodeInterface {
	mock := &MockNodeInterface{ctrl: ctrl}
	mock.recorder = &MockNodeInterfaceMockRecorder{mock}
	return mock
}

func (m *MockNodeInterface) EXPECT() *MockNodeInterfaceMockRecorder {
	return m.recorder
}

func (m *MockNodeInterface) Apply(arg0 context.Context, arg1 *v11.NodeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockNodeInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockNodeInterface) ApplyStatus(arg0 context.Context, arg1 *v11.NodeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockNodeInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockNodeInterface) Create(arg0 context.Context, arg1 *v1.Node, arg2 v10.CreateOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockNodeInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockNodeInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNodeInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockNodeInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockNodeInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNodeInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockNodeInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockNodeInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockNodeInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockNodeInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockNodeInterface)(nil).List), arg0, arg1)
}

func (m *MockNodeInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Node, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockNodeInterface)(nil).Patch), varargs...)
}

func (m *MockNodeInterface) PatchStatus(arg0 context.Context, arg1 string, arg2 []byte) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PatchStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) PatchStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PatchStatus, reflect.TypeOf((*MockNodeInterface)(nil).PatchStatus), arg0, arg1, arg2)
}

func (m *MockNodeInterface) Update(arg0 context.Context, arg1 *v1.Node, arg2 v10.UpdateOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockNodeInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockNodeInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Node, arg2 v10.UpdateOptions) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockNodeInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockNodeInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNodeInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockNodeInterface)(nil).Watch), arg0, arg1)
}
