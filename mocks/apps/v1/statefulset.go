package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
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

type MockStatefulSetsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulSetsGetterMockRecorder
}

type MockStatefulSetsGetterMockRecorder struct {
	mock *MockStatefulSetsGetter
}

func NewMockStatefulSetsGetter(ctrl *gomock.Controller) *MockStatefulSetsGetter {
	mock := &MockStatefulSetsGetter{ctrl: ctrl}
	mock.recorder = &MockStatefulSetsGetterMockRecorder{mock}
	return mock
}

func (m *MockStatefulSetsGetter) EXPECT() *MockStatefulSetsGetterMockRecorder {
	return m.recorder
}

func (m *MockStatefulSetsGetter) StatefulSets(arg0 string) v14.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StatefulSets, arg0)
	ret0, _ := ret[0].(v14.StatefulSetInterface)
	return ret0
}

func (mr *MockStatefulSetsGetterMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StatefulSets, reflect.TypeOf((*MockStatefulSetsGetter)(nil).StatefulSets), arg0)
}

type MockStatefulSetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStatefulSetInterfaceMockRecorder
}

type MockStatefulSetInterfaceMockRecorder struct {
	mock *MockStatefulSetInterface
}

func NewMockStatefulSetInterface(ctrl *gomock.Controller) *MockStatefulSetInterface {
	mock := &MockStatefulSetInterface{ctrl: ctrl}
	mock.recorder = &MockStatefulSetInterfaceMockRecorder{mock}
	return mock
}

func (m *MockStatefulSetInterface) EXPECT() *MockStatefulSetInterfaceMockRecorder {
	return m.recorder
}

func (m *MockStatefulSetInterface) Apply(arg0 context.Context, arg1 *v12.StatefulSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockStatefulSetInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) ApplyScale(arg0 context.Context, arg1 string, arg2 *v13.ScaleApplyConfiguration, arg3 v11.ApplyOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) ApplyScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyScale, reflect.TypeOf((*MockStatefulSetInterface)(nil).ApplyScale), arg0, arg1, arg2, arg3)
}

func (m *MockStatefulSetInterface) ApplyStatus(arg0 context.Context, arg1 *v12.StatefulSetApplyConfiguration, arg2 v11.ApplyOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockStatefulSetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) Create(arg0 context.Context, arg1 *v1.StatefulSet, arg2 v11.CreateOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockStatefulSetInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockStatefulSetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockStatefulSetInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockStatefulSetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockStatefulSetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockStatefulSetInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetScale, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetScale, reflect.TypeOf((*MockStatefulSetInterface)(nil).GetScale), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v1.StatefulSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.StatefulSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockStatefulSetInterface)(nil).List), arg0, arg1)
}

func (m *MockStatefulSetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockStatefulSetInterface)(nil).Patch), varargs...)
}

func (m *MockStatefulSetInterface) Update(arg0 context.Context, arg1 *v1.StatefulSet, arg2 v11.UpdateOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockStatefulSetInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v10.Scale, arg3 v11.UpdateOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateScale, reflect.TypeOf((*MockStatefulSetInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

func (m *MockStatefulSetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.StatefulSet, arg2 v11.UpdateOptions) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockStatefulSetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockStatefulSetInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockStatefulSetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockStatefulSetInterface)(nil).Watch), arg0, arg1)
}
