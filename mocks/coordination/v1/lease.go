package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/coordination/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/coordination/v1"
	v12 "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

type MockLeasesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockLeasesGetterMockRecorder
}

type MockLeasesGetterMockRecorder struct {
	mock *MockLeasesGetter
}

func NewMockLeasesGetter(ctrl *gomock.Controller) *MockLeasesGetter {
	mock := &MockLeasesGetter{ctrl: ctrl}
	mock.recorder = &MockLeasesGetterMockRecorder{mock}
	return mock
}

func (m *MockLeasesGetter) EXPECT() *MockLeasesGetterMockRecorder {
	return m.recorder
}

func (m *MockLeasesGetter) Leases(arg0 string) v12.LeaseInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Leases, arg0)
	ret0, _ := ret[0].(v12.LeaseInterface)
	return ret0
}

func (mr *MockLeasesGetterMockRecorder) Leases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Leases, reflect.TypeOf((*MockLeasesGetter)(nil).Leases), arg0)
}

type MockLeaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseInterfaceMockRecorder
}

type MockLeaseInterfaceMockRecorder struct {
	mock *MockLeaseInterface
}

func NewMockLeaseInterface(ctrl *gomock.Controller) *MockLeaseInterface {
	mock := &MockLeaseInterface{ctrl: ctrl}
	mock.recorder = &MockLeaseInterfaceMockRecorder{mock}
	return mock
}

func (m *MockLeaseInterface) EXPECT() *MockLeaseInterfaceMockRecorder {
	return m.recorder
}

func (m *MockLeaseInterface) Apply(arg0 context.Context, arg1 *v11.LeaseApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockLeaseInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) Create(arg0 context.Context, arg1 *v1.Lease, arg2 v10.CreateOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockLeaseInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLeaseInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockLeaseInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockLeaseInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockLeaseInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockLeaseInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.LeaseList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.LeaseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockLeaseInterface)(nil).List), arg0, arg1)
}

func (m *MockLeaseInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockLeaseInterface)(nil).Patch), varargs...)
}

func (m *MockLeaseInterface) Update(arg0 context.Context, arg1 *v1.Lease, arg2 v10.UpdateOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockLeaseInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockLeaseInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLeaseInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockLeaseInterface)(nil).Watch), arg0, arg1)
}
