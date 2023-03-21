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

type MockPersistentVolumeClaimsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockPersistentVolumeClaimsGetterMockRecorder
}

type MockPersistentVolumeClaimsGetterMockRecorder struct {
	mock *MockPersistentVolumeClaimsGetter
}

func NewMockPersistentVolumeClaimsGetter(ctrl *gomock.Controller) *MockPersistentVolumeClaimsGetter {
	mock := &MockPersistentVolumeClaimsGetter{ctrl: ctrl}
	mock.recorder = &MockPersistentVolumeClaimsGetterMockRecorder{mock}
	return mock
}

func (m *MockPersistentVolumeClaimsGetter) EXPECT() *MockPersistentVolumeClaimsGetterMockRecorder {
	return m.recorder
}

func (m *MockPersistentVolumeClaimsGetter) PersistentVolumeClaims(arg0 string) v12.PersistentVolumeClaimInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PersistentVolumeClaims, arg0)
	ret0, _ := ret[0].(v12.PersistentVolumeClaimInterface)
	return ret0
}

func (mr *MockPersistentVolumeClaimsGetterMockRecorder) PersistentVolumeClaims(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PersistentVolumeClaims, reflect.TypeOf((*MockPersistentVolumeClaimsGetter)(nil).PersistentVolumeClaims), arg0)
}

type MockPersistentVolumeClaimInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockPersistentVolumeClaimInterfaceMockRecorder
}

type MockPersistentVolumeClaimInterfaceMockRecorder struct {
	mock *MockPersistentVolumeClaimInterface
}

func NewMockPersistentVolumeClaimInterface(ctrl *gomock.Controller) *MockPersistentVolumeClaimInterface {
	mock := &MockPersistentVolumeClaimInterface{ctrl: ctrl}
	mock.recorder = &MockPersistentVolumeClaimInterfaceMockRecorder{mock}
	return mock
}

func (m *MockPersistentVolumeClaimInterface) EXPECT() *MockPersistentVolumeClaimInterfaceMockRecorder {
	return m.recorder
}

func (m *MockPersistentVolumeClaimInterface) Apply(arg0 context.Context, arg1 *v11.PersistentVolumeClaimApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) ApplyStatus(arg0 context.Context, arg1 *v11.PersistentVolumeClaimApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) Create(arg0 context.Context, arg1 *v1.PersistentVolumeClaim, arg2 v10.CreateOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PersistentVolumeClaimList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaimList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).List), arg0, arg1)
}

func (m *MockPersistentVolumeClaimInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Patch), varargs...)
}

func (m *MockPersistentVolumeClaimInterface) Update(arg0 context.Context, arg1 *v1.PersistentVolumeClaim, arg2 v10.UpdateOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) UpdateStatus(arg0 context.Context, arg1 *v1.PersistentVolumeClaim, arg2 v10.UpdateOptions) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockPersistentVolumeClaimInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPersistentVolumeClaimInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockPersistentVolumeClaimInterface)(nil).Watch), arg0, arg1)
}
