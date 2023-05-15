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

type MockResourceQuotasGetter struct {
	ctrl     *gomock.Controller
	recorder *MockResourceQuotasGetterMockRecorder
}

type MockResourceQuotasGetterMockRecorder struct {
	mock *MockResourceQuotasGetter
}

func NewMockResourceQuotasGetter(ctrl *gomock.Controller) *MockResourceQuotasGetter {
	mock := &MockResourceQuotasGetter{ctrl: ctrl}
	mock.recorder = &MockResourceQuotasGetterMockRecorder{mock}
	return mock
}

func (m *MockResourceQuotasGetter) EXPECT() *MockResourceQuotasGetterMockRecorder {
	return m.recorder
}

func (m *MockResourceQuotasGetter) ResourceQuotas(arg0 string) v12.ResourceQuotaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ResourceQuotas, arg0)
	ret0, _ := ret[0].(v12.ResourceQuotaInterface)
	return ret0
}

func (mr *MockResourceQuotasGetterMockRecorder) ResourceQuotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ResourceQuotas, reflect.TypeOf((*MockResourceQuotasGetter)(nil).ResourceQuotas), arg0)
}

type MockResourceQuotaInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceQuotaInterfaceMockRecorder
}

type MockResourceQuotaInterfaceMockRecorder struct {
	mock *MockResourceQuotaInterface
}

func NewMockResourceQuotaInterface(ctrl *gomock.Controller) *MockResourceQuotaInterface {
	mock := &MockResourceQuotaInterface{ctrl: ctrl}
	mock.recorder = &MockResourceQuotaInterfaceMockRecorder{mock}
	return mock
}

func (m *MockResourceQuotaInterface) EXPECT() *MockResourceQuotaInterfaceMockRecorder {
	return m.recorder
}

func (m *MockResourceQuotaInterface) Apply(arg0 context.Context, arg1 *v11.ResourceQuotaApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) ApplyStatus(arg0 context.Context, arg1 *v11.ResourceQuotaApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockResourceQuotaInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) Create(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.CreateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockResourceQuotaInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockResourceQuotaInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ResourceQuotaList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ResourceQuotaList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockResourceQuotaInterface)(nil).List), arg0, arg1)
}

func (m *MockResourceQuotaInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Patch), varargs...)
}

func (m *MockResourceQuotaInterface) Update(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.UpdateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) UpdateStatus(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.UpdateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockResourceQuotaInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockResourceQuotaInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceQuotaInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockResourceQuotaInterface)(nil).Watch), arg0, arg1)
}
