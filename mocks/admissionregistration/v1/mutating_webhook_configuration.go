package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/admissionregistration/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	v12 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

type MockMutatingWebhookConfigurationsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockMutatingWebhookConfigurationsGetterMockRecorder
}

type MockMutatingWebhookConfigurationsGetterMockRecorder struct {
	mock *MockMutatingWebhookConfigurationsGetter
}

func NewMockMutatingWebhookConfigurationsGetter(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationsGetter {
	mock := &MockMutatingWebhookConfigurationsGetter{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationsGetterMockRecorder{mock}
	return mock
}

func (m *MockMutatingWebhookConfigurationsGetter) EXPECT() *MockMutatingWebhookConfigurationsGetterMockRecorder {
	return m.recorder
}

func (m *MockMutatingWebhookConfigurationsGetter) MutatingWebhookConfigurations() v12.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.MutatingWebhookConfigurations)
	ret0, _ := ret[0].(v12.MutatingWebhookConfigurationInterface)
	return ret0
}

func (mr *MockMutatingWebhookConfigurationsGetterMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.MutatingWebhookConfigurations, reflect.TypeOf((*MockMutatingWebhookConfigurationsGetter)(nil).MutatingWebhookConfigurations))
}

type MockMutatingWebhookConfigurationInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockMutatingWebhookConfigurationInterfaceMockRecorder
}

type MockMutatingWebhookConfigurationInterfaceMockRecorder struct {
	mock *MockMutatingWebhookConfigurationInterface
}

func NewMockMutatingWebhookConfigurationInterface(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationInterface {
	mock := &MockMutatingWebhookConfigurationInterface{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationInterfaceMockRecorder{mock}
	return mock
}

func (m *MockMutatingWebhookConfigurationInterface) EXPECT() *MockMutatingWebhookConfigurationInterfaceMockRecorder {
	return m.recorder
}

func (m *MockMutatingWebhookConfigurationInterface) Apply(arg0 context.Context, arg1 *v11.MutatingWebhookConfigurationApplyConfiguration, arg2 v10.ApplyOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) Create(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.CreateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.MutatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).List), arg0, arg1)
}

func (m *MockMutatingWebhookConfigurationInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Patch), varargs...)
}

func (m *MockMutatingWebhookConfigurationInterface) Update(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.UpdateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockMutatingWebhookConfigurationInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMutatingWebhookConfigurationInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockMutatingWebhookConfigurationInterface)(nil).Watch), arg0, arg1)
}
