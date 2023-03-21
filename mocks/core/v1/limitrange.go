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
)

type MockLimitRangesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockLimitRangesGetterMockRecorder
}

type MockLimitRangesGetterMockRecorder struct {
	mock *MockLimitRangesGetter
}

func NewMockLimitRangesGetter(ctrl *gomock.Controller) *MockLimitRangesGetter {
	mock := &MockLimitRangesGetter{ctrl: ctrl}
	mock.recorder = &MockLimitRangesGetterMockRecorder{mock}
	return mock
}

func (m *MockLimitRangesGetter) EXPECT() *MockLimitRangesGetterMockRecorder {
	return m.recorder
}

func (m *MockLimitRangesGetter) LimitRanges(arg0 string) v12.LimitRangeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.LimitRanges, arg0)
	ret0, _ := ret[0].(v12.LimitRangeInterface)
	return ret0
}

func (mr *MockLimitRangesGetterMockRecorder) LimitRanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.LimitRanges, reflect.TypeOf((*MockLimitRangesGetter)(nil).LimitRanges), arg0)
}

type MockLimitRangeInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockLimitRangeInterfaceMockRecorder
}

type MockLimitRangeInterfaceMockRecorder struct {
	mock *MockLimitRangeInterface
}

func NewMockLimitRangeInterface(ctrl *gomock.Controller) *MockLimitRangeInterface {
	mock := &MockLimitRangeInterface{ctrl: ctrl}
	mock.recorder = &MockLimitRangeInterfaceMockRecorder{mock}
	return mock
}

func (m *MockLimitRangeInterface) EXPECT() *MockLimitRangeInterfaceMockRecorder {
	return m.recorder
}

func (m *MockLimitRangeInterface) Apply(arg0 context.Context, arg1 *v11.LimitRangeApplyConfiguration, arg2 v10.ApplyOptions) (*v1.LimitRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.LimitRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockLimitRangeInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) Create(arg0 context.Context, arg1 *v1.LimitRange, arg2 v10.CreateOptions) (*v1.LimitRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.LimitRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockLimitRangeInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLimitRangeInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockLimitRangeInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLimitRangeInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockLimitRangeInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.LimitRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.LimitRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockLimitRangeInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.LimitRangeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.LimitRangeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockLimitRangeInterface)(nil).List), arg0, arg1)
}

func (m *MockLimitRangeInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.LimitRange, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.LimitRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockLimitRangeInterface)(nil).Patch), varargs...)
}

func (m *MockLimitRangeInterface) Update(arg0 context.Context, arg1 *v1.LimitRange, arg2 v10.UpdateOptions) (*v1.LimitRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.LimitRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockLimitRangeInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockLimitRangeInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLimitRangeInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockLimitRangeInterface)(nil).Watch), arg0, arg1)
}
