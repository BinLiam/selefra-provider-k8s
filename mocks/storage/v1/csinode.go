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

type MockCSINodesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockCSINodesGetterMockRecorder
}

type MockCSINodesGetterMockRecorder struct {
	mock *MockCSINodesGetter
}

func NewMockCSINodesGetter(ctrl *gomock.Controller) *MockCSINodesGetter {
	mock := &MockCSINodesGetter{ctrl: ctrl}
	mock.recorder = &MockCSINodesGetterMockRecorder{mock}
	return mock
}

func (m *MockCSINodesGetter) EXPECT() *MockCSINodesGetterMockRecorder {
	return m.recorder
}

func (m *MockCSINodesGetter) CSINodes() v12.CSINodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSINodes)
	ret0, _ := ret[0].(v12.CSINodeInterface)
	return ret0
}

func (mr *MockCSINodesGetterMockRecorder) CSINodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSINodes, reflect.TypeOf((*MockCSINodesGetter)(nil).CSINodes))
}

type MockCSINodeInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockCSINodeInterfaceMockRecorder
}

type MockCSINodeInterfaceMockRecorder struct {
	mock *MockCSINodeInterface
}

func NewMockCSINodeInterface(ctrl *gomock.Controller) *MockCSINodeInterface {
	mock := &MockCSINodeInterface{ctrl: ctrl}
	mock.recorder = &MockCSINodeInterfaceMockRecorder{mock}
	return mock
}

func (m *MockCSINodeInterface) EXPECT() *MockCSINodeInterfaceMockRecorder {
	return m.recorder
}

func (m *MockCSINodeInterface) Apply(arg0 context.Context, arg1 *v11.CSINodeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockCSINodeInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) Create(arg0 context.Context, arg1 *v1.CSINode, arg2 v10.CreateOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockCSINodeInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSINodeInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockCSINodeInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCSINodeInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockCSINodeInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockCSINodeInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CSINodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.CSINodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCSINodeInterface)(nil).List), arg0, arg1)
}

func (m *MockCSINodeInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockCSINodeInterface)(nil).Patch), varargs...)
}

func (m *MockCSINodeInterface) Update(arg0 context.Context, arg1 *v1.CSINode, arg2 v10.UpdateOptions) (*v1.CSINode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSINode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockCSINodeInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockCSINodeInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCSINodeInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockCSINodeInterface)(nil).Watch), arg0, arg1)
}
