// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Math is an autogenerated mock type for the Math type
type Math struct {
	mock.Mock
}

// GetRectCenterColProperties provides a mock function with given fields: imageWidth, imageHeight, qtdCols, colHeight, indexCol
func (_m *Math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64) (float64, float64, float64, float64) {
	ret := _m.Called(int(imageWidth), int(imageHeight), int(qtdCols), int(colHeight), int(indexCol))

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64, float64, float64, float64, float64) float64); ok {
		r0 = rf(imageWidth, imageHeight, qtdCols, colHeight, indexCol)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(float64, float64, float64, float64, float64) float64); ok {
		r1 = rf(imageWidth, imageHeight, qtdCols, colHeight, indexCol)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 float64
	if rf, ok := ret.Get(2).(func(float64, float64, float64, float64, float64) float64); ok {
		r2 = rf(imageWidth, imageHeight, qtdCols, colHeight, indexCol)
	} else {
		r2 = ret.Get(2).(float64)
	}

	var r3 float64
	if rf, ok := ret.Get(3).(func(float64, float64, float64, float64, float64) float64); ok {
		r3 = rf(imageWidth, imageHeight, qtdCols, colHeight, indexCol)
	} else {
		r3 = ret.Get(3).(float64)
	}

	return r0, r1, r2, r3
}

// GetWidthPerCol provides a mock function with given fields: qtdCols
func (_m *Math) GetWidthPerCol(qtdCols float64) float64 {
	ret := _m.Called(qtdCols)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(qtdCols)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
