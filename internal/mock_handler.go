// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler.go

// Package internal is a generated GoMock package.
package internal

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCounterHandler is a mock of CounterHandler interface.
type MockCounterHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCounterHandlerMockRecorder
}

// MockCounterHandlerMockRecorder is the mock recorder for MockCounterHandler.
type MockCounterHandlerMockRecorder struct {
	mock *MockCounterHandler
}

// NewMockCounterHandler creates a new mock instance.
func NewMockCounterHandler(ctrl *gomock.Controller) *MockCounterHandler {
	mock := &MockCounterHandler{ctrl: ctrl}
	mock.recorder = &MockCounterHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCounterHandler) EXPECT() *MockCounterHandlerMockRecorder {
	return m.recorder
}

// DecrementHandler mocks base method.
func (m *MockCounterHandler) DecrementHandler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DecrementHandler", w, r)
}

// DecrementHandler indicates an expected call of DecrementHandler.
func (mr *MockCounterHandlerMockRecorder) DecrementHandler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementHandler", reflect.TypeOf((*MockCounterHandler)(nil).DecrementHandler), w, r)
}

// IncrementHandler mocks base method.
func (m *MockCounterHandler) IncrementHandler(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncrementHandler", w, r)
}

// IncrementHandler indicates an expected call of IncrementHandler.
func (mr *MockCounterHandlerMockRecorder) IncrementHandler(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementHandler", reflect.TypeOf((*MockCounterHandler)(nil).IncrementHandler), w, r)
}