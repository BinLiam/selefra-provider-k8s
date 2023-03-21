package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/api/policy/v1"
	v1beta1 "k8s.io/api/policy/v1beta1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
)

type MockPodsGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockPodsGetterMockRecorder
}

type MockPodsGetterMockRecorder struct {
	mock *MockPodsGetter
}

func NewMockPodsGetter(ctrl *gomock.Controller) *MockPodsGetter {
	mock := &MockPodsGetter{ctrl: ctrl}
	mock.recorder = &MockPodsGetterMockRecorder{mock}
	return mock
}

func (m *MockPodsGetter) EXPECT() *MockPodsGetterMockRecorder {
	return m.recorder
}

func (m *MockPodsGetter) Pods(arg0 string) v13.PodInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Pods, arg0)
	ret0, _ := ret[0].(v13.PodInterface)
	return ret0
}

func (mr *MockPodsGetterMockRecorder) Pods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Pods, reflect.TypeOf((*MockPodsGetter)(nil).Pods), arg0)
}

type MockPodInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockPodInterfaceMockRecorder
}

type MockPodInterfaceMockRecorder struct {
	mock *MockPodInterface
}

func NewMockPodInterface(ctrl *gomock.Controller) *MockPodInterface {
	mock := &MockPodInterface{ctrl: ctrl}
	mock.recorder = &MockPodInterfaceMockRecorder{mock}
	return mock
}

func (m *MockPodInterface) EXPECT() *MockPodInterfaceMockRecorder {
	return m.recorder
}

func (m *MockPodInterface) Apply(arg0 context.Context, arg1 *v12.PodApplyConfiguration, arg2 v11.ApplyOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockPodInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockPodInterface) ApplyStatus(arg0 context.Context, arg1 *v12.PodApplyConfiguration, arg2 v11.ApplyOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockPodInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockPodInterface) Bind(arg0 context.Context, arg1 *v1.Binding, arg2 v11.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Bind, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) Bind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Bind, reflect.TypeOf((*MockPodInterface)(nil).Bind), arg0, arg1, arg2)
}

func (m *MockPodInterface) Create(arg0 context.Context, arg1 *v1.Pod, arg2 v11.CreateOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockPodInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockPodInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockPodInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockPodInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockPodInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockPodInterface) Evict(arg0 context.Context, arg1 *v1beta1.Eviction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Evict, arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) Evict(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Evict, reflect.TypeOf((*MockPodInterface)(nil).Evict), arg0, arg1)
}

func (m *MockPodInterface) EvictV1(arg0 context.Context, arg1 *v10.Eviction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EvictV, arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) EvictV1(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EvictV, reflect.TypeOf((*MockPodInterface)(nil).EvictV1), arg0, arg1)
}

func (m *MockPodInterface) EvictV1beta1(arg0 context.Context, arg1 *v1beta1.Eviction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.EvictVbeta, arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) EvictV1beta1(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.EvictVbeta, reflect.TypeOf((*MockPodInterface)(nil).EvictV1beta1), arg0, arg1)
}

func (m *MockPodInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockPodInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockPodInterface) GetLogs(arg0 string, arg1 *v1.PodLogOptions) *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetLogs, arg0, arg1)
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) GetLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetLogs, reflect.TypeOf((*MockPodInterface)(nil).GetLogs), arg0, arg1)
}

func (m *MockPodInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockPodInterface)(nil).List), arg0, arg1)
}

func (m *MockPodInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockPodInterface)(nil).Patch), varargs...)
}

func (m *MockPodInterface) ProxyGet(arg0, arg1, arg2, arg3 string, arg4 map[string]string) rest.ResponseWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ProxyGet, arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(rest.ResponseWrapper)
	return ret0
}

func (mr *MockPodInterfaceMockRecorder) ProxyGet(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ProxyGet, reflect.TypeOf((*MockPodInterface)(nil).ProxyGet), arg0, arg1, arg2, arg3, arg4)
}

func (m *MockPodInterface) Update(arg0 context.Context, arg1 *v1.Pod, arg2 v11.UpdateOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockPodInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockPodInterface) UpdateEphemeralContainers(arg0 context.Context, arg1 string, arg2 *v1.Pod, arg3 v11.UpdateOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateEphemeralContainers, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) UpdateEphemeralContainers(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateEphemeralContainers, reflect.TypeOf((*MockPodInterface)(nil).UpdateEphemeralContainers), arg0, arg1, arg2, arg3)
}

func (m *MockPodInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Pod, arg2 v11.UpdateOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockPodInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockPodInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockPodInterface)(nil).Watch), arg0, arg1)
}
