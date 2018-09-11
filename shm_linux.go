//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

// Отображение разделяемой памяти в Linux.

func Shm(fname string, size int) ([]byte, error) {

	stat := unix.Stat_t{}

	var key int

	if err := unix.Stat(fname, &stat); err != nil {
		return nil, err
	} else {
		key = int(stat.Ino&0xffff) + int(stat.Dev&0xff<<16) + 0x62<<24
	}

	r1, _, err := unix.RawSyscall(unix.SYS_SHMGET, uintptr(key), uintptr(size), 0)

	if err != 0 {
		return nil, fmt.Errorf("shmget: %v", err)
	}

	r1, _, err = unix.RawSyscall(unix.SYS_SHMAT, r1, 0, 0)

	if err != 0 {
		return nil, fmt.Errorf("shmat: %v", err)
	}

	return (*[1<<31 - 1]byte) (unsafe.Pointer(r1)) [:size:size], nil
}
