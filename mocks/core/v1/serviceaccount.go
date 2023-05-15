package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/authentication/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type MockServiceAccountsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountsGetterMockRecorder
}

type MockServiceAccountsGetterMockRecorder struct {
	mock *MockServiceAccountsGetter
}

func NewMockServiceAccountsGetter(ctrl *gomock.Controller) *MockServiceAccountsGetter {
	mock := &MockServiceAccountsGetter{ctrl: ctrl}
	mock.recorder = &MockServiceAccountsGetterMockRecorder{mock}
	return mock
}

func (m *MockServiceAccountsGetter) EXPECT() *MockServiceAccountsGetterMockRecorder {
	return m.recorder
}

func (m *MockServiceAccountsGetter) ServiceAccounts(arg0 string) v13.ServiceAccountInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ServiceAccounts, arg0)
	ret0, _ := ret[0].(v13.ServiceAccountInterface)
	return ret0
}

func (mr *MockServiceAccountsGetterMockRecorder) ServiceAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ServiceAccounts, reflect.TypeOf((*MockServiceAccountsGetter)(nil).ServiceAccounts), arg0)
}

type MockServiceAccountInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountInterfaceMockRecorder
}

type MockServiceAccountInterfaceMockRecorder struct {
	mock *MockServiceAccountInterface
}

func NewMockServiceAccountInterface(ctrl *gomock.Controller) *MockServiceAccountInterface {
	mock := &MockServiceAccountInterface{ctrl: ctrl}
	mock.recorder = &MockServiceAccountInterfaceMockRecorder{mock}
	return mock
}

func (m *MockServiceAccountInterface) EXPECT() *MockServiceAccountInterfaceMockRecorder {
	return m.recorder
}

func (m *MockServiceAccountInterface) Apply(arg0 context.Context, arg1 *v12.ServiceAccountApplyConfiguration, arg2 v11.ApplyOptions) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockServiceAccountInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) Create(arg0 context.Context, arg1 *v10.ServiceAccount, arg2 v11.CreateOptions) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockServiceAccountInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) CreateToken(arg0 context.Context, arg1 string, arg2 *v1.TokenRequest, arg3 v11.CreateOptions) (*v1.TokenRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CreateToken, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.TokenRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) CreateToken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CreateToken, reflect.TypeOf((*MockServiceAccountInterface)(nil).CreateToken), arg0, arg1, arg2, arg3)
}

func (m *MockServiceAccountInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceAccountInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockServiceAccountInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceAccountInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockServiceAccountInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockServiceAccountInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v10.ServiceAccountList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v10.ServiceAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockServiceAccountInterface)(nil).List), arg0, arg1)
}

func (m *MockServiceAccountInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockServiceAccountInterface)(nil).Patch), varargs...)
}

func (m *MockServiceAccountInterface) Update(arg0 context.Context, arg1 *v10.ServiceAccount, arg2 v11.UpdateOptions) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockServiceAccountInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockServiceAccountInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceAccountInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockServiceAccountInterface)(nil).Watch), arg0, arg1)
}
