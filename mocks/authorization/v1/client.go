package mocks

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	rest "k8s.io/client-go/rest"
)

type MockAuthorizationV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockAuthorizationV1InterfaceMockRecorder
}

type MockAuthorizationV1InterfaceMockRecorder struct {
	mock *MockAuthorizationV1Interface
}

func NewMockAuthorizationV1Interface(ctrl *gomock.Controller) *MockAuthorizationV1Interface {
	mock := &MockAuthorizationV1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthorizationV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockAuthorizationV1Interface) EXPECT() *MockAuthorizationV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockAuthorizationV1Interface) LocalSubjectAccessReviews(arg0 string) v1.LocalSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.LocalSubjectAccessReviews, arg0)
	ret0, _ := ret[0].(v1.LocalSubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1InterfaceMockRecorder) LocalSubjectAccessReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.LocalSubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1Interface)(nil).LocalSubjectAccessReviews), arg0)
}

func (m *MockAuthorizationV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockAuthorizationV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockAuthorizationV1Interface)(nil).RESTClient))
}

func (m *MockAuthorizationV1Interface) SelfSubjectAccessReviews() v1.SelfSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SelfSubjectAccessReviews)
	ret0, _ := ret[0].(v1.SelfSubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1InterfaceMockRecorder) SelfSubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SelfSubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1Interface)(nil).SelfSubjectAccessReviews))
}

func (m *MockAuthorizationV1Interface) SelfSubjectRulesReviews() v1.SelfSubjectRulesReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SelfSubjectRulesReviews)
	ret0, _ := ret[0].(v1.SelfSubjectRulesReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1InterfaceMockRecorder) SelfSubjectRulesReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SelfSubjectRulesReviews, reflect.TypeOf((*MockAuthorizationV1Interface)(nil).SelfSubjectRulesReviews))
}

func (m *MockAuthorizationV1Interface) SubjectAccessReviews() v1.SubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.SubjectAccessReviews)
	ret0, _ := ret[0].(v1.SubjectAccessReviewInterface)
	return ret0
}

func (mr *MockAuthorizationV1InterfaceMockRecorder) SubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SubjectAccessReviews, reflect.TypeOf((*MockAuthorizationV1Interface)(nil).SubjectAccessReviews))
}
