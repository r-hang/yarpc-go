// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/peer (interfaces: Identifier,Peer)

// Copyright (c) 2023 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package peertest is a generated GoMock package.
package peertest

import (
	gomock "github.com/golang/mock/gomock"
	peer "go.uber.org/yarpc/api/peer"
	reflect "reflect"
)

// MockIdentifier is a mock of Identifier interface
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// Identifier mocks base method
func (m *MockIdentifier) Identifier() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockIdentifierMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockIdentifier)(nil).Identifier))
}

// MockPeer is a mock of Peer interface
type MockPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMockRecorder
}

// MockPeerMockRecorder is the mock recorder for MockPeer
type MockPeerMockRecorder struct {
	mock *MockPeer
}

// NewMockPeer creates a new mock instance
func NewMockPeer(ctrl *gomock.Controller) *MockPeer {
	mock := &MockPeer{ctrl: ctrl}
	mock.recorder = &MockPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeer) EXPECT() *MockPeerMockRecorder {
	return m.recorder
}

// EndRequest mocks base method
func (m *MockPeer) EndRequest() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndRequest")
}

// EndRequest indicates an expected call of EndRequest
func (mr *MockPeerMockRecorder) EndRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndRequest", reflect.TypeOf((*MockPeer)(nil).EndRequest))
}

// Identifier mocks base method
func (m *MockPeer) Identifier() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockPeerMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockPeer)(nil).Identifier))
}

// StartRequest mocks base method
func (m *MockPeer) StartRequest() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartRequest")
}

// StartRequest indicates an expected call of StartRequest
func (mr *MockPeerMockRecorder) StartRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRequest", reflect.TypeOf((*MockPeer)(nil).StartRequest))
}

// Status mocks base method
func (m *MockPeer) Status() peer.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(peer.Status)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockPeerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockPeer)(nil).Status))
}
