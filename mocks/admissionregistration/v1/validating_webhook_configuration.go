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

type MockValidatingWebhookConfigurationsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockValidatingWebhookConfigurationsGetterMockRecorder
}

type MockValidatingWebhookConfigurationsGetterMockRecorder struct {
	mock *MockValidatingWebhookConfigurationsGetter
}

func NewMockValidatingWebhookConfigurationsGetter(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationsGetter {
	mock := &MockValidatingWebhookConfigurationsGetter{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationsGetterMockRecorder{mock}
	return mock
}

func (m *MockValidatingWebhookConfigurationsGetter) EXPECT() *MockValidatingWebhookConfigurationsGetterMockRecorder {
	return m.recorder
}

func (m *MockValidatingWebhookConfigurationsGetter) ValidatingWebhookConfigurations() v12.ValidatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ValidatingWebhookConfigurations)
	ret0, _ := ret[0].(v12.ValidatingWebhookConfigurationInterface)
	return ret0
}

func (mr *MockValidatingWebhookConfigurationsGetterMockRecorder) ValidatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ValidatingWebhookConfigurations, reflect.TypeOf((*MockValidatingWebhookConfigurationsGetter)(nil).ValidatingWebhookConfigurations))
}

type MockValidatingWebhookConfigurationInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockValidatingWebhookConfigurationInterfaceMockRecorder
}

type MockValidatingWebhookConfigurationInterfaceMockRecorder struct {
	mock *MockValidatingWebhookConfigurationInterface
}

func NewMockValidatingWebhookConfigurationInterface(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationInterface {
	mock := &MockValidatingWebhookConfigurationInterface{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationInterfaceMockRecorder{mock}
	return mock
}

func (m *MockValidatingWebhookConfigurationInterface) EXPECT() *MockValidatingWebhookConfigurationInterfaceMockRecorder {
	return m.recorder
}

func (m *MockValidatingWebhookConfigurationInterface) Apply(arg0 context.Context, arg1 *v11.ValidatingWebhookConfigurationApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) Create(arg0 context.Context, arg1 *v1.ValidatingWebhookConfiguration, arg2 v10.CreateOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ValidatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).List), arg0, arg1)
}

func (m *MockValidatingWebhookConfigurationInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Patch), varargs...)
}

func (m *MockValidatingWebhookConfigurationInterface) Update(arg0 context.Context, arg1 *v1.ValidatingWebhookConfiguration, arg2 v10.UpdateOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockValidatingWebhookConfigurationInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockValidatingWebhookConfigurationInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockValidatingWebhookConfigurationInterface)(nil).Watch), arg0, arg1)
}
