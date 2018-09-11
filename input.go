//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"strconv"
	"strings"
	"unsafe"
)

// Типы для реализации отслеживаемых значений.

type InputValue interface {
	Topic() string
	Apply(string) error
}

type InputBoolValue struct {
	topic string
	value *int8
}

func (v InputBoolValue) Topic() string {
	return v.topic
}

func (v *InputBoolValue) Apply(s string) error {

	if strings.EqualFold(s, "true") || strings.EqualFold(s, "t") {
		*v.value = 1
		return nil
	}

	if strings.EqualFold(s, "false") || strings.EqualFold(s, "f") {
		*v.value = 0
		return nil
	}

	if strings.HasPrefix(s, "0x") {
		if i, err := strconv.ParseInt(s[2:], 16, 16); err != nil {
			return err
		} else {
			if i != 0 {
				*v.value = 1
			} else {
				*v.value = 0
			}
		}
	} else {
		if i, err := strconv.ParseInt(s, 10, 16); err != nil {
			return err
		} else {
			if i != 0 {
				*v.value = 1
			} else {
				*v.value = 0
			}
		}
	}

	return nil
}

func NewInputBoolValue(topic string, mem []byte, ofs int) InputBoolValue {
	return InputBoolValue{topic: topic, value: (*int8)(unsafe.Pointer(&mem[ofs]))}
}

type InputInt16Value struct {
	topic string
	value *int16
}

func (v InputInt16Value) Topic() string {
	return v.topic
}

func (v *InputInt16Value) Apply(s string) error {
	if strings.HasPrefix(s, "0x") {
		if i, err := strconv.ParseInt(s[2:], 16, 16); err != nil {
			return err
		} else {
			*v.value = int16(i)
		}
	} else {
		if i, err := strconv.ParseInt(s, 10, 16); err != nil {
			return err
		} else {
			*v.value = int16(i)
		}
	}
	return nil
}

func NewInputInt16Value(topic string, mem []byte, ofs int) InputInt16Value {
	return InputInt16Value{topic: topic, value: (*int16)(unsafe.Pointer(&mem[ofs]))}
}

type InputInt32Value struct {
	topic string
	value *int32
}

func (v InputInt32Value) Topic() string {
	return v.topic
}

func (v *InputInt32Value) Apply(s string) error {
	if strings.HasPrefix(s, "0x") {
		if i, err := strconv.ParseInt(s[2:], 16, 32); err != nil {
			return err
		} else {
			*v.value = int32(i)
		}
	} else {
		if i, err := strconv.ParseInt(s, 10, 32); err != nil {
			return err
		} else {
			*v.value = int32(i)
		}
	}
	return nil
}

func NewInputInt32Value(topic string, mem []byte, ofs int) InputInt32Value {
	return InputInt32Value{topic: topic, value: (*int32)(unsafe.Pointer(&mem[ofs]))}
}

type InputFloat32Value struct {
	topic string
	value *float32
}

func (v InputFloat32Value) Topic() string {
	return v.topic
}

func (v *InputFloat32Value) Apply(s string) error {
	if r, err := strconv.ParseFloat(s, 32); err != nil {
		return err
	} else {
		*v.value = float32(r)
		return nil
	}
}

func NewInputFloat32Value(topic string, mem []byte, ofs int) InputFloat32Value {
	return InputFloat32Value{topic: topic, value: (*float32)(unsafe.Pointer(&mem[ofs]))}
}
