package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/networking/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/networking/v1"
	v12 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

type MockNetworkPoliciesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkPoliciesGetterMockRecorder
}

type MockNetworkPoliciesGetterMockRecorder struct {
	mock *MockNetworkPoliciesGetter
}

func NewMockNetworkPoliciesGetter(ctrl *gomock.Controller) *MockNetworkPoliciesGetter {
	mock := &MockNetworkPoliciesGetter{ctrl: ctrl}
	mock.recorder = &MockNetworkPoliciesGetterMockRecorder{mock}
	return mock
}

func (m *MockNetworkPoliciesGetter) EXPECT() *MockNetworkPoliciesGetterMockRecorder {
	return m.recorder
}

func (m *MockNetworkPoliciesGetter) NetworkPolicies(arg0 string) v12.NetworkPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.NetworkPolicies, arg0)
	ret0, _ := ret[0].(v12.NetworkPolicyInterface)
	return ret0
}

func (mr *MockNetworkPoliciesGetterMockRecorder) NetworkPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.NetworkPolicies, reflect.TypeOf((*MockNetworkPoliciesGetter)(nil).NetworkPolicies), arg0)
}

type MockNetworkPolicyInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockNetworkPolicyInterfaceMockRecorder
}

type MockNetworkPolicyInterfaceMockRecorder struct {
	mock *MockNetworkPolicyInterface
}

func NewMockNetworkPolicyInterface(ctrl *gomock.Controller) *MockNetworkPolicyInterface {
	mock := &MockNetworkPolicyInterface{ctrl: ctrl}
	mock.recorder = &MockNetworkPolicyInterfaceMockRecorder{mock}
	return mock
}

func (m *MockNetworkPolicyInterface) EXPECT() *MockNetworkPolicyInterfaceMockRecorder {
	return m.recorder
}

func (m *MockNetworkPolicyInterface) Apply(arg0 context.Context, arg1 *v11.NetworkPolicyApplyConfiguration, arg2 v10.ApplyOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) ApplyStatus(arg0 context.Context, arg1 *v11.NetworkPolicyApplyConfiguration, arg2 v10.ApplyOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) Create(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.CreateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.NetworkPolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.NetworkPolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).List), arg0, arg1)
}

func (m *MockNetworkPolicyInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Patch), varargs...)
}

func (m *MockNetworkPolicyInterface) Update(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.UpdateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) UpdateStatus(arg0 context.Context, arg1 *v1.NetworkPolicy, arg2 v10.UpdateOptions) (*v1.NetworkPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.NetworkPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockNetworkPolicyInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPolicyInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockNetworkPolicyInterface)(nil).Watch), arg0, arg1)
}
