//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"strconv"
	"strings"
)

// Разные утилитарные функции общего назначения.

func ParseInt(s string) (int, error) {
	if strings.HasPrefix(s, "0x") {
		i, err := strconv.ParseInt(s[2:], 16, 32)
		return int(i), err
	} else {
		i, err := strconv.ParseInt(s, 10, 32)
		return int(i), err
	}
}
