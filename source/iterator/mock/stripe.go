// Code generated by MockGen. DO NOT EDIT.
// Source: source/iterator/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	sdk "github.com/conduitio/conduit-connector-sdk"
	models "github.com/conduitio/conduit-connector-stripe/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockIterator) Next() (sdk.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(sdk.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// MockStripe is a mock of Stripe interface.
type MockStripe struct {
	ctrl     *gomock.Controller
	recorder *MockStripeMockRecorder
}

// MockStripeMockRecorder is the mock recorder for MockStripe.
type MockStripeMockRecorder struct {
	mock *MockStripe
}

// NewMockStripe creates a new mock instance.
func NewMockStripe(ctrl *gomock.Controller) *MockStripe {
	mock := &MockStripe{ctrl: ctrl}
	mock.recorder = &MockStripeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStripe) EXPECT() *MockStripeMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockStripe) GetResource(startingAfter string) (models.StripeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", startingAfter)
	ret0, _ := ret[0].(models.StripeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockStripeMockRecorder) GetResource(startingAfter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockStripe)(nil).GetResource), startingAfter)
}
