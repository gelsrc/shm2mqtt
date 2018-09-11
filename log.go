//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"log"
	"os"
)

// Точки доступа для вывода логов в утилите.

var DEBUG = log.New(os.Stdout, "", 0)
var ERROR = log.New(os.Stderr, "", 0)
