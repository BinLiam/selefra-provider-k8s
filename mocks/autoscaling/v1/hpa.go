package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/autoscaling/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	v12 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
)

type MockHorizontalPodAutoscalersGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockHorizontalPodAutoscalersGetterMockRecorder
}

type MockHorizontalPodAutoscalersGetterMockRecorder struct {
	mock *MockHorizontalPodAutoscalersGetter
}

func NewMockHorizontalPodAutoscalersGetter(ctrl *gomock.Controller) *MockHorizontalPodAutoscalersGetter {
	mock := &MockHorizontalPodAutoscalersGetter{ctrl: ctrl}
	mock.recorder = &MockHorizontalPodAutoscalersGetterMockRecorder{mock}
	return mock
}

func (m *MockHorizontalPodAutoscalersGetter) EXPECT() *MockHorizontalPodAutoscalersGetterMockRecorder {
	return m.recorder
}

func (m *MockHorizontalPodAutoscalersGetter) HorizontalPodAutoscalers(arg0 string) v12.HorizontalPodAutoscalerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.HorizontalPodAutoscalers, arg0)
	ret0, _ := ret[0].(v12.HorizontalPodAutoscalerInterface)
	return ret0
}

func (mr *MockHorizontalPodAutoscalersGetterMockRecorder) HorizontalPodAutoscalers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.HorizontalPodAutoscalers, reflect.TypeOf((*MockHorizontalPodAutoscalersGetter)(nil).HorizontalPodAutoscalers), arg0)
}

type MockHorizontalPodAutoscalerInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockHorizontalPodAutoscalerInterfaceMockRecorder
}

type MockHorizontalPodAutoscalerInterfaceMockRecorder struct {
	mock *MockHorizontalPodAutoscalerInterface
}

func NewMockHorizontalPodAutoscalerInterface(ctrl *gomock.Controller) *MockHorizontalPodAutoscalerInterface {
	mock := &MockHorizontalPodAutoscalerInterface{ctrl: ctrl}
	mock.recorder = &MockHorizontalPodAutoscalerInterfaceMockRecorder{mock}
	return mock
}

func (m *MockHorizontalPodAutoscalerInterface) EXPECT() *MockHorizontalPodAutoscalerInterfaceMockRecorder {
	return m.recorder
}

func (m *MockHorizontalPodAutoscalerInterface) Apply(arg0 context.Context, arg1 *v11.HorizontalPodAutoscalerApplyConfiguration, arg2 v10.ApplyOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) ApplyStatus(arg0 context.Context, arg1 *v11.HorizontalPodAutoscalerApplyConfiguration, arg2 v10.ApplyOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) Create(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.CreateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.HorizontalPodAutoscalerList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscalerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).List), arg0, arg1)
}

func (m *MockHorizontalPodAutoscalerInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Patch), varargs...)
}

func (m *MockHorizontalPodAutoscalerInterface) Update(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.UpdateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) UpdateStatus(arg0 context.Context, arg1 *v1.HorizontalPodAutoscaler, arg2 v10.UpdateOptions) (*v1.HorizontalPodAutoscaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.HorizontalPodAutoscaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockHorizontalPodAutoscalerInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockHorizontalPodAutoscalerInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockHorizontalPodAutoscalerInterface)(nil).Watch), arg0, arg1)
}
