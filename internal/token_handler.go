// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: TokenEndpointHandler)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	fosite "github.com/ory/fosite"
)

// MockTokenEndpointHandler is a mock of TokenEndpointHandler interface.
type MockTokenEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTokenEndpointHandlerMockRecorder
}

// MockTokenEndpointHandlerMockRecorder is the mock recorder for MockTokenEndpointHandler.
type MockTokenEndpointHandlerMockRecorder struct {
	mock *MockTokenEndpointHandler
}

// NewMockTokenEndpointHandler creates a new mock instance.
func NewMockTokenEndpointHandler(ctrl *gomock.Controller) *MockTokenEndpointHandler {
	mock := &MockTokenEndpointHandler{ctrl: ctrl}
	mock.recorder = &MockTokenEndpointHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenEndpointHandler) EXPECT() *MockTokenEndpointHandlerMockRecorder {
	return m.recorder
}

// CanHandleTokenEndpointRequest mocks base method.
func (m *MockTokenEndpointHandler) CanHandleTokenEndpointRequest(arg0 context.Context, arg1 fosite.AccessRequester) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanHandleTokenEndpointRequest", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanHandleTokenEndpointRequest indicates an expected call of CanHandleTokenEndpointRequest.
func (mr *MockTokenEndpointHandlerMockRecorder) CanHandleTokenEndpointRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanHandleTokenEndpointRequest", reflect.TypeOf((*MockTokenEndpointHandler)(nil).CanHandleTokenEndpointRequest), arg0, arg1)
}

// CanSkipClientAuth mocks base method.
func (m *MockTokenEndpointHandler) CanSkipClientAuth(arg0 context.Context, arg1 fosite.AccessRequester) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSkipClientAuth", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSkipClientAuth indicates an expected call of CanSkipClientAuth.
func (mr *MockTokenEndpointHandlerMockRecorder) CanSkipClientAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSkipClientAuth", reflect.TypeOf((*MockTokenEndpointHandler)(nil).CanSkipClientAuth), arg0, arg1)
}

// HandleTokenEndpointRequest mocks base method.
func (m *MockTokenEndpointHandler) HandleTokenEndpointRequest(arg0 context.Context, arg1 fosite.AccessRequester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTokenEndpointRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTokenEndpointRequest indicates an expected call of HandleTokenEndpointRequest.
func (mr *MockTokenEndpointHandlerMockRecorder) HandleTokenEndpointRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTokenEndpointRequest", reflect.TypeOf((*MockTokenEndpointHandler)(nil).HandleTokenEndpointRequest), arg0, arg1)
}

// PopulateTokenEndpointResponse mocks base method.
func (m *MockTokenEndpointHandler) PopulateTokenEndpointResponse(arg0 context.Context, arg1 fosite.AccessRequester, arg2 fosite.AccessResponder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateTokenEndpointResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateTokenEndpointResponse indicates an expected call of PopulateTokenEndpointResponse.
func (mr *MockTokenEndpointHandlerMockRecorder) PopulateTokenEndpointResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateTokenEndpointResponse", reflect.TypeOf((*MockTokenEndpointHandler)(nil).PopulateTokenEndpointResponse), arg0, arg1, arg2)
}
