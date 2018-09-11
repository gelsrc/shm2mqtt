//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"fmt"
	"regexp"
)

// Парсинг данных в конфигурационном файлы для секций Coil, Holdreg, Instat, Inreg.
// В реализации есть продублированный код для большей гибкости правки и доработки.

var paramRegexp = regexp.MustCompile(",")

func ExtractCoils(cfg ConfigFile) ([]Coil, error) {

	var res []Coil

	for key, value := range cfg.Values["Coil"] {

		params := paramRegexp.Split(value, 7)

		if len(params) < 7 {
			return nil, fmt.Errorf("unexpected data format in Coil section [%v=%v]", key, value)
		}

		label := params[6]

		newVar := Coil{Label: label}

		if i, err := ParseInt(params[0]); err != nil {
			return nil, err
		} else {
			newVar.Unknown1 = i
		}

		if i, err := ParseInt(params[1]); err != nil {
			return nil, err
		} else {
			newVar.Size = i
		}

		if i, err := ParseInt(params[2]); err != nil {
			return nil, err
		} else {
			newVar.Type = VarType(i)
		}

		if i, err := ParseInt(params[3]); err != nil {
			return nil, err
		} else {
			newVar.Offset = i
		}

		if i, err := ParseInt(params[4]); err != nil {
			return nil, err
		} else {
			newVar.Unknown5 = i
		}

		if i, err := ParseInt(params[5]); err != nil {
			return nil, err
		} else {
			newVar.Unknown6 = i
		}

		res = append(res, newVar)
	}

	return res, nil
}

func ExtractHoldregs(cfg ConfigFile) ([]HoldReg, error) {

	var res []HoldReg

	for key, value := range cfg.Values["Holdreg"] {

		params := paramRegexp.Split(value, 7)

		if len(params) < 7 {
			return nil, fmt.Errorf("unexpected data format in Holdreg section [%v=%v]", key, value)
		}

		label := params[6]

		newVar := HoldReg{Label: label}

		if i, err := ParseInt(params[0]); err != nil {
			return nil, err
		} else {
			newVar.Unknown1 = i
		}

		if i, err := ParseInt(params[1]); err != nil {
			return nil, err
		} else {
			newVar.Size = i
		}

		if i, err := ParseInt(params[2]); err != nil {
			return nil, err
		} else {
			newVar.Type = VarType(i)
		}

		if i, err := ParseInt(params[3]); err != nil {
			return nil, err
		} else {
			newVar.Offset = i
		}

		if i, err := ParseInt(params[4]); err != nil {
			return nil, err
		} else {
			newVar.Unknown5 = i
		}

		if i, err := ParseInt(params[5]); err != nil {
			return nil, err
		} else {
			newVar.Unknown6 = i
		}

		res = append(res, newVar)
	}

	return res, nil
}

func ExtractInStats(cfg ConfigFile) ([]InStat, error) {

	var res []InStat

	for key, value := range cfg.Values["Instat"] {

		params := paramRegexp.Split(value, 7)

		if len(params) < 7 {
			return nil, fmt.Errorf("unexpected data format in Instat section [%v=%v]", key, value)
		}

		label := params[6]

		newVar := InStat{Label: label}

		if i, err := ParseInt(params[0]); err != nil {
			return nil, err
		} else {
			newVar.Unknown1 = i
		}

		if i, err := ParseInt(params[1]); err != nil {
			return nil, err
		} else {
			newVar.Size = i
		}

		if i, err := ParseInt(params[2]); err != nil {
			return nil, err
		} else {
			newVar.Type = VarType(i)
		}

		if i, err := ParseInt(params[3]); err != nil {
			return nil, err
		} else {
			newVar.Offset = i
		}

		if i, err := ParseInt(params[4]); err != nil {
			return nil, err
		} else {
			newVar.Unknown5 = i
		}

		if i, err := ParseInt(params[5]); err != nil {
			return nil, err
		} else {
			newVar.Unknown6 = i
		}

		res = append(res, newVar)
	}

	return res, nil
}

func ExtractInregs(cfg ConfigFile) ([]InReg, error) {

	var res []InReg

	for key, value := range cfg.Values["Inreg"] {

		params := paramRegexp.Split(value, 7)

		if len(params) < 7 {
			return nil, fmt.Errorf("unexpected data format in Inreg section [%v=%v]", key, value)
		}

		label := params[6]

		newVar := InReg{Label: label}

		if i, err := ParseInt(params[0]); err != nil {
			return nil, err
		} else {
			newVar.Unknown1 = i
		}

		if i, err := ParseInt(params[1]); err != nil {
			return nil, err
		} else {
			newVar.Size = i
		}

		if i, err := ParseInt(params[2]); err != nil {
			return nil, err
		} else {
			newVar.Type = VarType(i)
		}

		if i, err := ParseInt(params[3]); err != nil {
			return nil, err
		} else {
			newVar.Offset = i
		}

		if i, err := ParseInt(params[4]); err != nil {
			return nil, err
		} else {
			newVar.Unknown5 = i
		}

		if i, err := ParseInt(params[5]); err != nil {
			return nil, err
		} else {
			newVar.Unknown6 = i
		}

		res = append(res, newVar)
	}

	return res, nil
}
