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

type MockCronJobsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCronJobsGetterMockRecorder
}

type MockCronJobsGetterMockRecorder struct {
	mock *MockCronJobsGetter
}

func NewMockCronJobsGetter(ctrl *gomock.Controller) *MockCronJobsGetter {
	mock := &MockCronJobsGetter{ctrl: ctrl}
	mock.recorder = &MockCronJobsGetterMockRecorder{mock}
	return mock
}

func (m *MockCronJobsGetter) EXPECT() *MockCronJobsGetterMockRecorder {
	return m.recorder
}

func (m *MockCronJobsGetter) CronJobs(arg0 string) v12.CronJobInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CronJobs, arg0)
	ret0, _ := ret[0].(v12.CronJobInterface)
	return ret0
}

func (mr *MockCronJobsGetterMockRecorder) CronJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CronJobs, reflect.TypeOf((*MockCronJobsGetter)(nil).CronJobs), arg0)
}

type MockCronJobInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCronJobInterfaceMockRecorder
}

type MockCronJobInterfaceMockRecorder struct {
	mock *MockCronJobInterface
}

func NewMockCronJobInterface(ctrl *gomock.Controller) *MockCronJobInterface {
	mock := &MockCronJobInterface{ctrl: ctrl}
	mock.recorder = &MockCronJobInterfaceMockRecorder{mock}
	return mock
}

func (m *MockCronJobInterface) EXPECT() *MockCronJobInterfaceMockRecorder {
	return m.recorder
}

func (m *MockCronJobInterface) Apply(arg0 context.Context, arg1 *v11.CronJobApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockCronJobInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) ApplyStatus(arg0 context.Context, arg1 *v11.CronJobApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockCronJobInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) Create(arg0 context.Context, arg1 *v1.CronJob, arg2 v10.CreateOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockCronJobInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCronJobInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockCronJobInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCronJobInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockCronJobInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockCronJobInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CronJobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.CronJobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCronJobInterface)(nil).List), arg0, arg1)
}

func (m *MockCronJobInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockCronJobInterface)(nil).Patch), varargs...)
}

func (m *MockCronJobInterface) Update(arg0 context.Context, arg1 *v1.CronJob, arg2 v10.UpdateOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockCronJobInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) UpdateStatus(arg0 context.Context, arg1 *v1.CronJob, arg2 v10.UpdateOptions) (*v1.CronJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CronJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockCronJobInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockCronJobInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCronJobInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockCronJobInterface)(nil).Watch), arg0, arg1)
}
