//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

// Значения и типы, связанные с конфигурационным файлом "load_files.srv".

type VarType int

const (
	VarTypeBool VarType = 0
	VarTypeInt  VarType = 1
	VarTypeLong VarType = 2
	VarTypeReal VarType = 3
)

type InStat struct {
	Unknown1 int
	Size     int
	Type     VarType
	Offset   int
	Unknown5 int
	Unknown6 int
	Label    string
}

type InReg struct {
	Unknown1 int
	Size     int
	Type     VarType
	Offset   int
	Unknown5 int
	Unknown6 int
	Label    string
}

type HoldReg struct {
	Unknown1 int
	Size     int
	Type     VarType
	Offset   int
	Unknown5 int
	Unknown6 int
	Label    string
}

type Coil struct {
	Unknown1 int
	Size     int
	Type     VarType
	Offset   int
	Unknown5 int
	Unknown6 int
	Label    string
}
