package v1

import (
	context "context"
	"github.com/selefra/selefra-provider-k8s/constants"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/storage/v1"
	v12 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

type MockVolumeAttachmentsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeAttachmentsGetterMockRecorder
}

type MockVolumeAttachmentsGetterMockRecorder struct {
	mock *MockVolumeAttachmentsGetter
}

func NewMockVolumeAttachmentsGetter(ctrl *gomock.Controller) *MockVolumeAttachmentsGetter {
	mock := &MockVolumeAttachmentsGetter{ctrl: ctrl}
	mock.recorder = &MockVolumeAttachmentsGetterMockRecorder{mock}
	return mock
}

func (m *MockVolumeAttachmentsGetter) EXPECT() *MockVolumeAttachmentsGetterMockRecorder {
	return m.recorder
}

func (m *MockVolumeAttachmentsGetter) VolumeAttachments() v12.VolumeAttachmentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.VolumeAttachments)
	ret0, _ := ret[0].(v12.VolumeAttachmentInterface)
	return ret0
}

func (mr *MockVolumeAttachmentsGetterMockRecorder) VolumeAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.VolumeAttachments, reflect.TypeOf((*MockVolumeAttachmentsGetter)(nil).VolumeAttachments))
}

type MockVolumeAttachmentInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeAttachmentInterfaceMockRecorder
}

type MockVolumeAttachmentInterfaceMockRecorder struct {
	mock *MockVolumeAttachmentInterface
}

func NewMockVolumeAttachmentInterface(ctrl *gomock.Controller) *MockVolumeAttachmentInterface {
	mock := &MockVolumeAttachmentInterface{ctrl: ctrl}
	mock.recorder = &MockVolumeAttachmentInterfaceMockRecorder{mock}
	return mock
}

func (m *MockVolumeAttachmentInterface) EXPECT() *MockVolumeAttachmentInterfaceMockRecorder {
	return m.recorder
}

func (m *MockVolumeAttachmentInterface) Apply(arg0 context.Context, arg1 *v11.VolumeAttachmentApplyConfiguration, arg2 v10.ApplyOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Apply, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Apply, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Apply), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) ApplyStatus(arg0 context.Context, arg1 *v11.VolumeAttachmentApplyConfiguration, arg2 v10.ApplyOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ApplyStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ApplyStatus, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) Create(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.CreateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Create, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Create, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Create), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Delete, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Delete, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Delete), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.DeleteCollection, arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DeleteCollection, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Get), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.VolumeAttachmentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*v1.VolumeAttachmentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).List), arg0, arg1)
}

func (m *MockVolumeAttachmentInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.Patch, varargs...)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Patch, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Patch), varargs...)
}

func (m *MockVolumeAttachmentInterface) Update(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.UpdateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Update, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Update, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Update), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) UpdateStatus(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.UpdateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.UpdateStatus, arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.UpdateStatus, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

func (m *MockVolumeAttachmentInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Watch, arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockVolumeAttachmentInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Watch, reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Watch), arg0, arg1)
}
