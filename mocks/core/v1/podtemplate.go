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

type MockPodTemplatesGetter struct {
	ctrl		*gomock.Controller
	recorder	*MockPodTemplatesGetterMockRecorder
}

type MockPodTemplatesGetterMockRecorder struct {
	mock *MockPodTemplatesGetter
}

func NewMockPodTemplatesGetter(ctrl *gomock.Controller) *MockPodTemplatesGetter {
	mock := &MockPodTemplatesGetter{ctrl: ctrl}
	mock.recorder = &MockPodTemplatesGetterMockRecorder{mock}
	return mock
}

func (m *MockPodTemplatesGetter) EXPECT() *MockPodTemplatesGetterMockRecorder {
	return m.recorder
}

func (m *MockPodTemplatesGetter) PodTemplates(arg0 string) v12.PodTemplateInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.PodTemplates, arg0)
	ret0, _ := ret[0].(v12.PodTemplateInterface)
	return ret0
}

func (mr *MockPodTemplatesGetterMockRecorder) PodTemplates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.PodTemplates, reflect.TypeOf((*MockPodTemplatesGetter)(nil).PodTemplates), arg0)
}

type MockPodTemplateInterface struct {
	ctrl		*gomock.Controller
	recorder	*MockPodTemplateInterfaceMockRecorder
}

type MockPodTemplateInterfaceMockRecorder struct {
	mock *MockPodTemplateInterface
}

func NewMockPodTemplateInterface(ctrl *gomock.Controller) *MockPodTemplateInterface {
	mock := &MockPodTemplateInterface{ctrl: ctrl}
	mock.recorder = &MockPodTemplateInterfaceMockRecorder{mock}
	return mock
}

func (m *MockPodTemplateInterface) EXPECT() *MockPodTemplateInterfaceMockRecorder {
	return m.recorder
}

func (m *MockPodTemplateInterface) Apply(arg0 context.Context, arg1 *v11.PodTemplateApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockPodTemplateInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) Create(arg0 context.Context, arg1 *v1.PodTemplate, arg2 v10.CreateOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockPodTemplateInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodTemplateInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockPodTemplateInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPodTemplateInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockPodTemplateInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockPodTemplateInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PodTemplateList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.PodTemplateList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockPodTemplateInterface)(nil).List), arg0, arg1)
}

func (m *MockPodTemplateInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockPodTemplateInterface)(nil).Patch), varargs...)
}

func (m *MockPodTemplateInterface) Update(arg0 context.Context, arg1 *v1.PodTemplate, arg2 v10.UpdateOptions) (*v1.PodTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockPodTemplateInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockPodTemplateInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPodTemplateInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockPodTemplateInterface)(nil).Watch), arg0, arg1)
}
