// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bytes "bytes"
import maroto "github.com/johnfercher/maroto"
import mock "github.com/stretchr/testify/mock"

// Maroto is an autogenerated mock type for the Maroto type
type Maroto struct {
	mock.Mock
}

// Barcode provides a mock function with given fields: code, width, height, marginTop
func (_m *Maroto) Barcode(code string, width float64, height float64, marginTop float64) error {
	ret := _m.Called(code, width, height, marginTop)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, float64, float64) error); ok {
		r0 = rf(code, width, height, marginTop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Col provides a mock function with given fields: label, closure
func (_m *Maroto) Col(label string, closure func()) {
	_m.Called(label, closure)
}

// ColSpace provides a mock function with given fields:
func (_m *Maroto) ColSpace() {
	_m.Called()
}

// ColSpaces provides a mock function with given fields: qtd
func (_m *Maroto) ColSpaces(qtd int) {
	_m.Called(qtd)
}

// FileImage provides a mock function with given fields: filePathName, marginTop
func (_m *Maroto) FileImage(filePathName string, marginTop float64) {
	_m.Called(filePathName, marginTop)
}

// Line provides a mock function with given fields:
func (_m *Maroto) Line() {
	_m.Called()
}

// Output provides a mock function with given fields:
func (_m *Maroto) Output() (bytes.Buffer, error) {
	ret := _m.Called()

	var r0 bytes.Buffer
	if rf, ok := ret.Get(0).(func() bytes.Buffer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OutputFileAndClose provides a mock function with given fields: filePathName
func (_m *Maroto) OutputFileAndClose(filePathName string) error {
	ret := _m.Called(filePathName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filePathName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QrCode provides a mock function with given fields: code
func (_m *Maroto) QrCode(code string) {
	_m.Called(code)
}

// Row provides a mock function with given fields: label, height, closure
func (_m *Maroto) Row(label string, height float64, closure func()) {
	_m.Called(label, height, closure)
}

// RowTableList provides a mock function with given fields: label, header, contents
func (_m *Maroto) RowTableList(label string, header []string, contents [][]string) {
	_m.Called(label, header, contents)
}

// SetDebugMode provides a mock function with given fields: on
func (_m *Maroto) SetDebugMode(on bool) {
	_m.Called(on)
}

// sign provides a mock function with given fields: label, fontFamily, fontStyle, fontSize
func (_m *Maroto) Sign(label string, fontFamily maroto.Family, fontStyle maroto.Style, fontSize float64) {
	_m.Called(label, fontFamily, fontStyle, fontSize)
}

// text provides a mock function with given fields: text, fontFamily, fontStyle, fontSize, marginTop, align
func (_m *Maroto) Text(text string, fontFamily maroto.Family, fontStyle maroto.Style, fontSize float64, marginTop float64, align maroto.Align) {
	_m.Called(text, fontFamily, fontStyle, fontSize, marginTop, align)
}
