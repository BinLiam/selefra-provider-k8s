package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/batch/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/batch/v1"
	v12 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type MockJobsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockJobsGetterMockRecorder
}

type MockJobsGetterMockRecorder struct {
	mock *MockJobsGetter
}

func NewMockJobsGetter(ctrl *gomock.Controller) *MockJobsGetter {
	mock := &MockJobsGetter{ctrl: ctrl}
	mock.recorder = &MockJobsGetterMockRecorder{mock}
	return mock
}

func (m *MockJobsGetter) EXPECT() *MockJobsGetterMockRecorder {
	return m.recorder
}

func (m *MockJobsGetter) Jobs(arg0 string) v12.JobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Jobs, arg0)
	ret0, _ := ret[0].(v12.JobInterface)
	return ret0
}

func (mr *MockJobsGetterMockRecorder) Jobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Jobs, reflect.TypeOf((*MockJobsGetter)(nil).Jobs), arg0)
}

type MockJobInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobInterfaceMockRecorder
}

type MockJobInterfaceMockRecorder struct {
	mock *MockJobInterface
}

func NewMockJobInterface(ctrl *gomock.Controller) *MockJobInterface {
	mock := &MockJobInterface{ctrl: ctrl}
	mock.recorder = &MockJobInterfaceMockRecorder{mock}
	return mock
}

func (m *MockJobInterface) EXPECT() *MockJobInterfaceMockRecorder {
	return m.recorder
}

func (m *MockJobInterface) Apply(arg0 context.Context, arg1 *v11.JobApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockJobInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockJobInterface) ApplyStatus(arg0 context.Context, arg1 *v11.JobApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockJobInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockJobInterface) Create(arg0 context.Context, arg1 *v1.Job, arg2 v10.CreateOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockJobInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockJobInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockJobInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockJobInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockJobInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockJobInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockJobInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockJobInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockJobInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockJobInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.JobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.JobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockJobInterface)(nil).List), arg0, arg1)
}

func (m *MockJobInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Job, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockJobInterface)(nil).Patch), varargs...)
}

func (m *MockJobInterface) Update(arg0 context.Context, arg1 *v1.Job, arg2 v10.UpdateOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockJobInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockJobInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Job, arg2 v10.UpdateOptions) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockJobInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockJobInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockJobInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockJobInterface)(nil).Watch), arg0, arg1)
}
