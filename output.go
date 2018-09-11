//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"strconv"
	"unsafe"
)

// Типы для реализации публикуемых значений.

type OutputValue interface {
	Sync() bool
	Topic() string
	String() string
}

type OutputBoolValue struct {
	topic string
	last  int8
	value *int8
}

func (v OutputBoolValue) Topic() string {
	return v.topic
}

func (v *OutputBoolValue) Sync() bool {
	if v.last == *v.value {
		return false
	}
	v.last = *v.value
	return true
}

func (v OutputBoolValue) String() string {
	return strconv.FormatInt(int64(v.last), 10)
}

func NewOutputBoolValue(topic string, mem []byte, ofs int) OutputBoolValue {
	return OutputBoolValue{topic: topic, value: (*int8)(unsafe.Pointer(&mem[ofs]))}
}

type OutputInt16Value struct {
	topic string
	last  int16
	value *int16
}

func (v OutputInt16Value) Topic() string {
	return v.topic
}

func (v *OutputInt16Value) Sync() bool {
	if v.last == *v.value {
		return false
	}
	v.last = *v.value
	return true
}

func (v OutputInt16Value) String() string {
	return strconv.FormatInt(int64(v.last), 10)
}

func NewOutputInt16Value(topic string, mem []byte, ofs int) OutputInt16Value {
	return OutputInt16Value{topic: topic, value: (*int16)(unsafe.Pointer(&mem[ofs]))}
}

type OutputInt32Value struct {
	topic string
	last  int32
	value *int32
}

func (v OutputInt32Value) Topic() string {
	return v.topic
}

func (v *OutputInt32Value) Sync() bool {
	if v.last == *v.value {
		return false
	}
	v.last = *v.value
	return true
}

func (v OutputInt32Value) String() string {
	return strconv.FormatInt(int64(v.last), 10)
}

func NewOutputInt32Value(topic string, mem []byte, ofs int) OutputInt32Value {
	return OutputInt32Value{topic: topic, value: (*int32)(unsafe.Pointer(&mem[ofs]))}
}

type OutputFloat32Value struct {
	topic string
	last  float32
	value *float32
}

func (v OutputFloat32Value) Topic() string {
	return v.topic
}

func (v *OutputFloat32Value) Sync() bool {
	if v.last == *v.value {
		return false
	}
	v.last = *v.value
	return true
}

func (v OutputFloat32Value) String() string {
	return strconv.FormatFloat(float64(v.last), 'f', -1, 32)
}

func NewOutputFloat32Value(topic string, mem []byte, ofs int) OutputFloat32Value {
	return OutputFloat32Value{topic: topic, value: (*float32)(unsafe.Pointer(&mem[ofs]))}
}
