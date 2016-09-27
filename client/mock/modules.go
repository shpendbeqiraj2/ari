// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: Modules)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Modules interface
type MockModules struct {
	ctrl     *gomock.Controller
	recorder *_MockModulesRecorder
}

// Recorder for MockModules (not exported)
type _MockModulesRecorder struct {
	mock *MockModules
}

func NewMockModules(ctrl *gomock.Controller) *MockModules {
	mock := &MockModules{ctrl: ctrl}
	mock.recorder = &_MockModulesRecorder{mock}
	return mock
}

func (_m *MockModules) EXPECT() *_MockModulesRecorder {
	return _m.recorder
}

func (_m *MockModules) Data(_param0 string) (ari.ModuleData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(ari.ModuleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockModulesRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockModules) Get(_param0 string) *ari.ModuleHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*ari.ModuleHandle)
	return ret0
}

func (_mr *_MockModulesRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockModules) List() ([]*ari.ModuleHandle, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]*ari.ModuleHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockModulesRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockModules) Load(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Load", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockModulesRecorder) Load(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Load", arg0)
}

func (_m *MockModules) Reload(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Reload", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockModulesRecorder) Reload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reload", arg0)
}

func (_m *MockModules) Unload(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Unload", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockModulesRecorder) Unload(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unload", arg0)
}
