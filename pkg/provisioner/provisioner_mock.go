// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/provisioner/types.go

// Package provisioner is a generated GoMock package.
package provisioner

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	node "github.com/mahakamcloud/mahakam/pkg/node"
)

// MockProvisioner is a mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// CreateNode mocks base method
func (m *MockProvisioner) CreateNode(config node.NodeCreateConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNode", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNode indicates an expected call of CreateNode
func (mr *MockProvisionerMockRecorder) CreateNode(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNode", reflect.TypeOf((*MockProvisioner)(nil).CreateNode), config)
}
