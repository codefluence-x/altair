// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gin-gonic/gin (interfaces: ResponseWriter)

// Package mock is a generated GoMock package.
package mock

import (
	bufio "bufio"
	net "net"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResponseWriter is a mock of ResponseWriter interface
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockResponseWriterMockRecorder
}

// MockResponseWriterMockRecorder is the mock recorder for MockResponseWriter
type MockResponseWriterMockRecorder struct {
	mock *MockResponseWriter
}

// NewMockResponseWriter creates a new mock instance
func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &MockResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseWriter) EXPECT() *MockResponseWriterMockRecorder {
	return m.recorder
}

// CloseNotify mocks base method
func (m *MockResponseWriter) CloseNotify() <-chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseNotify")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

// CloseNotify indicates an expected call of CloseNotify
func (mr *MockResponseWriterMockRecorder) CloseNotify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseNotify", reflect.TypeOf((*MockResponseWriter)(nil).CloseNotify))
}

// Flush mocks base method
func (m *MockResponseWriter) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush
func (mr *MockResponseWriterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockResponseWriter)(nil).Flush))
}

// Header mocks base method
func (m *MockResponseWriter) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockResponseWriterMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockResponseWriter)(nil).Header))
}

// Hijack mocks base method
func (m *MockResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hijack")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(*bufio.ReadWriter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Hijack indicates an expected call of Hijack
func (mr *MockResponseWriterMockRecorder) Hijack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hijack", reflect.TypeOf((*MockResponseWriter)(nil).Hijack))
}

// Pusher mocks base method
func (m *MockResponseWriter) Pusher() http.Pusher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pusher")
	ret0, _ := ret[0].(http.Pusher)
	return ret0
}

// Pusher indicates an expected call of Pusher
func (mr *MockResponseWriterMockRecorder) Pusher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pusher", reflect.TypeOf((*MockResponseWriter)(nil).Pusher))
}

// Size mocks base method
func (m *MockResponseWriter) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockResponseWriterMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockResponseWriter)(nil).Size))
}

// Status mocks base method
func (m *MockResponseWriter) Status() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(int)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockResponseWriterMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockResponseWriter)(nil).Status))
}

// Write mocks base method
func (m *MockResponseWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockResponseWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockResponseWriter)(nil).Write), arg0)
}

// WriteHeader mocks base method
func (m *MockResponseWriter) WriteHeader(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", arg0)
}

// WriteHeader indicates an expected call of WriteHeader
func (mr *MockResponseWriterMockRecorder) WriteHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockResponseWriter)(nil).WriteHeader), arg0)
}

// WriteHeaderNow mocks base method
func (m *MockResponseWriter) WriteHeaderNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeaderNow")
}

// WriteHeaderNow indicates an expected call of WriteHeaderNow
func (mr *MockResponseWriterMockRecorder) WriteHeaderNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeaderNow", reflect.TypeOf((*MockResponseWriter)(nil).WriteHeaderNow))
}

// WriteString mocks base method
func (m *MockResponseWriter) WriteString(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString
func (mr *MockResponseWriterMockRecorder) WriteString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockResponseWriter)(nil).WriteString), arg0)
}

// Written mocks base method
func (m *MockResponseWriter) Written() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Written")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Written indicates an expected call of Written
func (mr *MockResponseWriterMockRecorder) Written() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Written", reflect.TypeOf((*MockResponseWriter)(nil).Written))
}