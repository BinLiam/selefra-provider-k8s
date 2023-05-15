package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/certificates/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/certificates/v1"
	v12 "k8s.io/client-go/kubernetes/typed/certificates/v1"
)

type MockCertificateSigningRequestsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestsGetterMockRecorder
}

type MockCertificateSigningRequestsGetterMockRecorder struct {
	mock *MockCertificateSigningRequestsGetter
}

func NewMockCertificateSigningRequestsGetter(ctrl *gomock.Controller) *MockCertificateSigningRequestsGetter {
	mock := &MockCertificateSigningRequestsGetter{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestsGetterMockRecorder{mock}
	return mock
}

func (m *MockCertificateSigningRequestsGetter) EXPECT() *MockCertificateSigningRequestsGetterMockRecorder {
	return m.recorder
}

func (m *MockCertificateSigningRequestsGetter) CertificateSigningRequests() v12.CertificateSigningRequestInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CertificateSigningRequests)
	ret0, _ := ret[0].(v12.CertificateSigningRequestInterface)
	return ret0
}

func (mr *MockCertificateSigningRequestsGetterMockRecorder) CertificateSigningRequests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CertificateSigningRequests, reflect.TypeOf((*MockCertificateSigningRequestsGetter)(nil).CertificateSigningRequests))
}

type MockCertificateSigningRequestInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestInterfaceMockRecorder
}

type MockCertificateSigningRequestInterfaceMockRecorder struct {
	mock *MockCertificateSigningRequestInterface
}

func NewMockCertificateSigningRequestInterface(ctrl *gomock.Controller) *MockCertificateSigningRequestInterface {
	mock := &MockCertificateSigningRequestInterface{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestInterfaceMockRecorder{mock}
	return mock
}

func (m *MockCertificateSigningRequestInterface) EXPECT() *MockCertificateSigningRequestInterfaceMockRecorder {
	return m.recorder
}

func (m *MockCertificateSigningRequestInterface) Apply(arg0 context.Context, arg1 *v11.CertificateSigningRequestApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) ApplyStatus(arg0 context.Context, arg1 *v11.CertificateSigningRequestApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) Create(arg0 context.Context, arg1 *v1.CertificateSigningRequest, arg2 v10.CreateOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CertificateSigningRequestList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.CertificateSigningRequestList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).List), arg0, arg1)
}

func (m *MockCertificateSigningRequestInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Patch), varargs...)
}

func (m *MockCertificateSigningRequestInterface) Update(arg0 context.Context, arg1 *v1.CertificateSigningRequest, arg2 v10.UpdateOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) UpdateApproval(arg0 context.Context, arg1 string, arg2 *v1.CertificateSigningRequest, arg3 v10.UpdateOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateApproval, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) UpdateApproval(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateApproval, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).UpdateApproval), arg0, arg1, arg2, arg3)
}

func (m *MockCertificateSigningRequestInterface) UpdateStatus(arg0 context.Context, arg1 *v1.CertificateSigningRequest, arg2 v10.UpdateOptions) (*v1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockCertificateSigningRequestInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCertificateSigningRequestInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockCertificateSigningRequestInterface)(nil).Watch), arg0, arg1)
}
