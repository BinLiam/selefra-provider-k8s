package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/autoscaling/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/apps/v1"
	v13 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type MockReplicaSetsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockReplicaSetsGetterMockRecorder
}

type MockReplicaSetsGetterMockRecorder struct {
	mock *MockReplicaSetsGetter
}

func NewMockReplicaSetsGetter(ctrl *gomock.Controller) *MockReplicaSetsGetter {
	mock := &MockReplicaSetsGetter{ctrl: ctrl}
	mock.recorder = &MockReplicaSetsGetterMockRecorder{mock}
	return mock
}

func (m *MockReplicaSetsGetter) EXPECT() *MockReplicaSetsGetterMockRecorder {
	return m.recorder
}

func (m *MockReplicaSetsGetter) ReplicaSets(arg0 string) v14.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ReplicaSets, arg0)
	ret0, _ := ret[0].(v14.ReplicaSetInterface)
	return ret0
}

func (mr *MockReplicaSetsGetterMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ReplicaSets, reflect.TypeOf((*MockReplicaSetsGetter)(nil).ReplicaSets), arg0)
}

type MockReplicaSetInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockReplicaSetInterfaceMockRecorder
}

type MockReplicaSetInterfaceMockRecorder struct {
	mock *MockReplicaSetInterface
}

func NewMockReplicaSetInterface(ctrl *gomock.Controller) *MockReplicaSetInterface {
	mock := &MockReplicaSetInterface{ctrl: ctrl}
	mock.recorder = &MockReplicaSetInterfaceMockRecorder{mock}
	return mock
}

func (m *MockReplicaSetInterface) EXPECT() *MockReplicaSetInterfaceMockRecorder {
	return m.recorder
}

func (m *MockReplicaSetInterface) Apply(arg0 context.Context, arg1 *v12.ReplicaSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockReplicaSetInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) ApplyScale(arg0 context.Context, arg1 string, arg2 *v13.ScaleApplyConfiguration, arg3 v11.ApplyOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) ApplyScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyScale, reflect.TypeOf((*MockReplicaSetInterface)(nil).ApplyScale), arg0, arg1, arg2, arg3)
}

func (m *MockReplicaSetInterface) ApplyStatus(arg0 context.Context, arg1 *v12.ReplicaSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockReplicaSetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) Create(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.CreateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockReplicaSetInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReplicaSetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockReplicaSetInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockReplicaSetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockReplicaSetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockReplicaSetInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetScale, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetScale, reflect.TypeOf((*MockReplicaSetInterface)(nil).GetScale), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v1.ReplicaSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ReplicaSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockReplicaSetInterface)(nil).List), arg0, arg1)
}

func (m *MockReplicaSetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockReplicaSetInterface)(nil).Patch), varargs...)
}

func (m *MockReplicaSetInterface) Update(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.UpdateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockReplicaSetInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v10.Scale, arg3 v11.UpdateOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateScale, reflect.TypeOf((*MockReplicaSetInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

func (m *MockReplicaSetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.ReplicaSet, arg2 v11.UpdateOptions) (*v1.ReplicaSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ReplicaSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockReplicaSetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockReplicaSetInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockReplicaSetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockReplicaSetInterface)(nil).Watch), arg0, arg1)
}
