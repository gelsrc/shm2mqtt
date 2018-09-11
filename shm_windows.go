//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

// Заглушка отображения разделяемой памяти в Windows.

func Shm(fname string, size int) ([]byte, error) {
	return make([]byte, size), nil
}
