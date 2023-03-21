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

type MockIngressesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockIngressesGetterMockRecorder
}

type MockIngressesGetterMockRecorder struct {
	mock *MockIngressesGetter
}

func NewMockIngressesGetter(ctrl *gomock.Controller) *MockIngressesGetter {
	mock := &MockIngressesGetter{ctrl: ctrl}
	mock.recorder = &MockIngressesGetterMockRecorder{mock}
	return mock
}

func (m *MockIngressesGetter) EXPECT() *MockIngressesGetterMockRecorder {
	return m.recorder
}

func (m *MockIngressesGetter) Ingresses(arg0 string) v12.IngressInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Ingresses, arg0)
	ret0, _ := ret[0].(v12.IngressInterface)
	return ret0
}

func (mr *MockIngressesGetterMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Ingresses, reflect.TypeOf((*MockIngressesGetter)(nil).Ingresses), arg0)
}

type MockIngressInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockIngressInterfaceMockRecorder
}

type MockIngressInterfaceMockRecorder struct {
	mock *MockIngressInterface
}

func NewMockIngressInterface(ctrl *gomock.Controller) *MockIngressInterface {
	mock := &MockIngressInterface{ctrl: ctrl}
	mock.recorder = &MockIngressInterfaceMockRecorder{mock}
	return mock
}

func (m *MockIngressInterface) EXPECT() *MockIngressInterfaceMockRecorder {
	return m.recorder
}

func (m *MockIngressInterface) Apply(arg0 context.Context, arg1 *v11.IngressApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockIngressInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockIngressInterface) ApplyStatus(arg0 context.Context, arg1 *v11.IngressApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockIngressInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockIngressInterface) Create(arg0 context.Context, arg1 *v1.Ingress, arg2 v10.CreateOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockIngressInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockIngressInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockIngressInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockIngressInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockIngressInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockIngressInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockIngressInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockIngressInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockIngressInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockIngressInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.IngressList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.IngressList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockIngressInterface)(nil).List), arg0, arg1)
}

func (m *MockIngressInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockIngressInterface)(nil).Patch), varargs...)
}

func (m *MockIngressInterface) Update(arg0 context.Context, arg1 *v1.Ingress, arg2 v10.UpdateOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockIngressInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockIngressInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Ingress, arg2 v10.UpdateOptions) (*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockIngressInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockIngressInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIngressInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockIngressInterface)(nil).Watch), arg0, arg1)
}
