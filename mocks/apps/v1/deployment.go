package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/autoscaling/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/apps/v1"
	v13 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	v14 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type MockDeploymentsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentsGetterMockRecorder
}

type MockDeploymentsGetterMockRecorder struct {
	mock *MockDeploymentsGetter
}

func NewMockDeploymentsGetter(ctrl *gomock.Controller) *MockDeploymentsGetter {
	mock := &MockDeploymentsGetter{ctrl: ctrl}
	mock.recorder = &MockDeploymentsGetterMockRecorder{mock}
	return mock
}

func (m *MockDeploymentsGetter) EXPECT() *MockDeploymentsGetterMockRecorder {
	return m.recorder
}

func (m *MockDeploymentsGetter) Deployments(arg0 string) v14.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Deployments, arg0)
	ret0, _ := ret[0].(v14.DeploymentInterface)
	return ret0
}

func (mr *MockDeploymentsGetterMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Deployments, reflect.TypeOf((*MockDeploymentsGetter)(nil).Deployments), arg0)
}

type MockDeploymentInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentInterfaceMockRecorder
}

type MockDeploymentInterfaceMockRecorder struct {
	mock *MockDeploymentInterface
}

func NewMockDeploymentInterface(ctrl *gomock.Controller) *MockDeploymentInterface {
	mock := &MockDeploymentInterface{ctrl: ctrl}
	mock.recorder = &MockDeploymentInterfaceMockRecorder{mock}
	return mock
}

func (m *MockDeploymentInterface) EXPECT() *MockDeploymentInterfaceMockRecorder {
	return m.recorder
}

func (m *MockDeploymentInterface) Apply(arg0 context.Context, arg1 *v12.DeploymentApplyConfiguration, arg2 v11.ApplyOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockDeploymentInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) ApplyScale(arg0 context.Context, arg1 string, arg2 *v13.ScaleApplyConfiguration, arg3 v11.ApplyOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) ApplyScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyScale, reflect.TypeOf((*MockDeploymentInterface)(nil).ApplyScale), arg0, arg1, arg2, arg3)
}

func (m *MockDeploymentInterface) ApplyStatus(arg0 context.Context, arg1 *v12.DeploymentApplyConfiguration, arg2 v11.ApplyOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockDeploymentInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) Create(arg0 context.Context, arg1 *v1.Deployment, arg2 v11.CreateOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockDeploymentInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) Delete(arg0 context.Context, arg1 string, arg2 v11.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDeploymentInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockDeploymentInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) DeleteCollection(arg0 context.Context, arg1 v11.DeleteOptions, arg2 v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDeploymentInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockDeploymentInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) Get(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockDeploymentInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) GetScale(arg0 context.Context, arg1 string, arg2 v11.GetOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetScale, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) GetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetScale, reflect.TypeOf((*MockDeploymentInterface)(nil).GetScale), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) List(arg0 context.Context, arg1 v11.ListOptions) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockDeploymentInterface)(nil).List), arg0, arg1)
}

func (m *MockDeploymentInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v11.PatchOptions, arg5 ...string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockDeploymentInterface)(nil).Patch), varargs...)
}

func (m *MockDeploymentInterface) Update(arg0 context.Context, arg1 *v1.Deployment, arg2 v11.UpdateOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockDeploymentInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) UpdateScale(arg0 context.Context, arg1 string, arg2 *v10.Scale, arg3 v11.UpdateOptions) (*v10.Scale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateScale, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Scale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) UpdateScale(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateScale, reflect.TypeOf((*MockDeploymentInterface)(nil).UpdateScale), arg0, arg1, arg2, arg3)
}

func (m *MockDeploymentInterface) UpdateStatus(arg0 context.Context, arg1 *v1.Deployment, arg2 v11.UpdateOptions) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockDeploymentInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockDeploymentInterface) Watch(arg0 context.Context, arg1 v11.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDeploymentInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockDeploymentInterface)(nil).Watch), arg0, arg1)
}
