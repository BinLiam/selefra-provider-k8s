package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	rest "k8s.io/client-go/rest"
)

type MockAuthorizationV1beta1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAuthorizationV1beta1InterfaceMockRecorder
}

type MockAuthorizationV1beta1InterfaceMockRecorder struct {
	mock *MockAuthorizationV1beta1Interface
}

func NewMockAuthorizationV1beta1Interface(ctrl *gomock.Controller) *MockAuthorizationV1beta1Interface {
	mock := &MockAuthorizationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthorizationV1beta1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAuthorizationV1beta1Interface) EXPECT() *MockAuthorizationV1beta1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAuthorizationV1beta1Interface) LocalSubjectAccessReviews(arg0 string) v1beta1.LocalSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.LocalSubjectAccessReviews, arg0)
	ret0, _ := ret[0].(v1beta1.LocalSubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) LocalSubjectAccessReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.LocalSubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).LocalSubjectAccessReviews), arg0)
}

func (m *MockAuthorizationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).RESTClient))
}

func (m *MockAuthorizationV1beta1Interface) SelfSubjectAccessReviews() v1beta1.SelfSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SelfSubjectAccessReviews)
	ret0, _ := ret[0].(v1beta1.SelfSubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SelfSubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SelfSubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SelfSubjectAccessReviews))
}

func (m *MockAuthorizationV1beta1Interface) SelfSubjectRulesReviews() v1beta1.SelfSubjectRulesReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SelfSubjectRulesReviews)
	ret0, _ := ret[0].(v1beta1.SelfSubjectRulesReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SelfSubjectRulesReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SelfSubjectRulesReviews, reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SelfSubjectRulesReviews))
}

func (m *MockAuthorizationV1beta1Interface) SubjectAccessReviews() v1beta1.SubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SubjectAccessReviews)
	ret0, _ := ret[0].(v1beta1.SubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SubjectAccessReviews))
}
