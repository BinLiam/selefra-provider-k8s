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

type MockSecretsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockSecretsGetterMockRecorder
}

type MockSecretsGetterMockRecorder struct {
	mock *MockSecretsGetter
}

func NewMockSecretsGetter(ctrl *gomock.Controller) *MockSecretsGetter {
	mock := &MockSecretsGetter{ctrl: ctrl}
	mock.recorder = &MockSecretsGetterMockRecorder{mock}
	return mock
}

func (m *MockSecretsGetter) EXPECT() *MockSecretsGetterMockRecorder {
	return m.recorder
}

func (m *MockSecretsGetter) Secrets(arg0 string) v12.SecretInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Secrets, arg0)
	ret0, _ := ret[0].(v12.SecretInterface)
	return ret0
}

func (mr *MockSecretsGetterMockRecorder) Secrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Secrets, reflect.TypeOf((*MockSecretsGetter)(nil).Secrets), arg0)
}

type MockSecretInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockSecretInterfaceMockRecorder
}

type MockSecretInterfaceMockRecorder struct {
	mock *MockSecretInterface
}

func NewMockSecretInterface(ctrl *gomock.Controller) *MockSecretInterface {
	mock := &MockSecretInterface{ctrl: ctrl}
	mock.recorder = &MockSecretInterfaceMockRecorder{mock}
	return mock
}

func (m *MockSecretInterface) EXPECT() *MockSecretInterfaceMockRecorder {
	return m.recorder
}

func (m *MockSecretInterface) Apply(arg0 context.Context, arg1 *v11.SecretApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockSecretInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockSecretInterface) Create(arg0 context.Context, arg1 *v1.Secret, arg2 v10.CreateOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockSecretInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockSecretInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSecretInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockSecretInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockSecretInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSecretInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockSecretInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockSecretInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockSecretInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockSecretInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockSecretInterface)(nil).List), arg0, arg1)
}

func (m *MockSecretInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockSecretInterface)(nil).Patch), varargs...)
}

func (m *MockSecretInterface) Update(arg0 context.Context, arg1 *v1.Secret, arg2 v10.UpdateOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockSecretInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockSecretInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecretInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockSecretInterface)(nil).Watch), arg0, arg1)
}
