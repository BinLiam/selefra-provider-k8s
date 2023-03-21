package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/autoscaling/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type MockReplicationControllersGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockReplicationControllersGetterMockRecorder
}

type MockReplicationControllersGetterMockRecorder struct {
	mock *MockReplicationControllersGetter
}

func NewMockReplicationControllersGetter(ctrl *gomock.Controller) *MockReplicationControllersGetter {
	mock := &MockReplicationControllersGetter{ctrl: ctrl}
	mock.recorder = &MockReplicationControllersGetterMockRecorder{mock}
	return mock
}

func (m *MockReplicationControllersGetter) EXPECT() *MockReplicationControllersGetterMockRecorder {
	return m.recorder
}

func (m *MockReplicationControllersGetter) ReplicationControllers(arg0 string) v13.ReplicationControllerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicationControllers, arg0)
	ret0, _ := ret[0].(v13.ReplicationControllerInterface)
	return ret0
}

func (mr *MockReplicationControllersGetterMockRecorder) ReplicationControllers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicationControllers, reflect.TypeOf((*MockReplicationControllersGetter)(nil).ReplicationControllers), arg0)
}

type MockReplicationControllerInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockReplicationControllerInterfaceMockRecorder
}

type MockReplicationControllerInterfaceMockRecorder struct {
	mock *MockReplicationControllerInterface
}

func NewMockReplicationControllerInterface(ctrl *gomock.Controller) *MockReplicationControllerInterface {
	mock := &MockReplicationControllerInterface{ctrl: ctrl}
	mock.recorder = &MockReplicationControllerInterfaceMockRecorder{mock}
	return mock
}

func (m *MockReplicationControllerInterface) EXPECT() *MockReplicationControllerInterfaceMockRecorder {
	return m.recorder
}

func (m *MockReplicationControllerInterface) Apply(arg0 context.Context, arg1 *v12.ReplicationControllerApplyConfiguration, arg2 v11.ApplyOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) ApplyStatus(arg0 context.Context, arg1 *v12.ReplicationControllerApplyConfiguration, arg2 v11.ApplyOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockReplicationControllerInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) Create(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.CreateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReplicationControllerInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockReplicationControllerInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetScale, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetScale, reflect.TypeOf((*MockReplicationControllerInterface)(nil).GetScale), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v10.ReplicationControllerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v10.ReplicationControllerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockReplicationControllerInterface)(nil).List), arg0, arg1)
}

func (m *MockReplicationControllerInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Patch), varargs...)
}

func (m *MockReplicationControllerInterface) Update(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.UpdateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v1.Scale, arg3 v11.UpdateOptions) (*v1.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateScale, reflect.TypeOf((*MockReplicationControllerInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

func (m *MockReplicationControllerInterface) UpdateStatus(arg0 context.Context, arg1 *v10.ReplicationController, arg2 v11.UpdateOptions) (*v10.ReplicationController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ReplicationController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockReplicationControllerInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockReplicationControllerInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicationControllerInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockReplicationControllerInterface)(nil).Watch), arg0, arg1)
}
