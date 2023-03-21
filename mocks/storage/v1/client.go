package v1

import (
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	rest "k8s.io/client-go/rest"
)

type MockStorageV1Interface struct {
	ctrl		*gomock.Controller
	recorder	*MockStorageV1InterfaceMockRecorder
}

type MockStorageV1InterfaceMockRecorder struct {
	mock *MockStorageV1Interface
}

func NewMockStorageV1Interface(ctrl *gomock.Controller) *MockStorageV1Interface {
	mock := &MockStorageV1Interface{ctrl: ctrl}
	mock.recorder = &MockStorageV1InterfaceMockRecorder{mock}
	return mock
}

func (m *MockStorageV1Interface) EXPECT() *MockStorageV1InterfaceMockRecorder {
	return m.recorder
}

func (m *MockStorageV1Interface) CSIDrivers() v1.CSIDriverInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSIDrivers)
	ret0, _ := ret[0].(v1.CSIDriverInterface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) CSIDrivers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSIDrivers, reflect.TypeOf((*MockStorageV1Interface)(nil).CSIDrivers))
}

func (m *MockStorageV1Interface) CSINodes() v1.CSINodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSINodes)
	ret0, _ := ret[0].(v1.CSINodeInterface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) CSINodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSINodes, reflect.TypeOf((*MockStorageV1Interface)(nil).CSINodes))
}

func (m *MockStorageV1Interface) CSIStorageCapacities(arg0 string) v1.CSIStorageCapacityInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.CSIStorageCapacities, arg0)
	ret0, _ := ret[0].(v1.CSIStorageCapacityInterface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) CSIStorageCapacities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.CSIStorageCapacities, reflect.TypeOf((*MockStorageV1Interface)(nil).CSIStorageCapacities), arg0)
}

func (m *MockStorageV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.RESTClient)
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.RESTClient, reflect.TypeOf((*MockStorageV1Interface)(nil).RESTClient))
}

func (m *MockStorageV1Interface) StorageClasses() v1.StorageClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.StorageClasses)
	ret0, _ := ret[0].(v1.StorageClassInterface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.StorageClasses, reflect.TypeOf((*MockStorageV1Interface)(nil).StorageClasses))
}

func (m *MockStorageV1Interface) VolumeAttachments() v1.VolumeAttachmentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.VolumeAttachments)
	ret0, _ := ret[0].(v1.VolumeAttachmentInterface)
	return ret0
}

func (mr *MockStorageV1InterfaceMockRecorder) VolumeAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.VolumeAttachments, reflect.TypeOf((*MockStorageV1Interface)(nil).VolumeAttachments))
}
