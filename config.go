//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"bufio"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"regexp"
)

// Загрузка INI-файла и распределение строк по секциям.

type SectionData map[string]string

type ConfigSections map[string]SectionData

// Представление конфигурационного файла.
type ConfigFile struct {
	Values ConfigSections
}

// Загрузка конфигурационного файла
func (cfg *ConfigFile) Load(fname string) error {

	cfg.Values = make(ConfigSections)

	input, err := os.Open(fname)

	if err != nil {
		return err
	}

	rows := bufio.NewReader(input)

	section := ""

	for {
		line, err := rows.ReadBytes('\n')

		if err == nil {
			lf := len(line) - 1
			cr := lf - 1

			if cr >= 0 && line[cr] == '\r' {
				line = line[:cr]
			} else {
				line = line[:lf]
			}
		}

		if err == nil || len(line) > 0 {

			dec := charmap.Windows1251.NewDecoder()

			var s string

			if utf8bytes, err := dec.Bytes(line); err != nil {
				return err
			} else {
				s = string(utf8bytes)
			}

			if s, err := cfg.parse(section, s); err != nil {
				return err
			} else {
				section = s
			}
		}

		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
	}

	return nil
}

var sectionRegexp = regexp.MustCompile("^\\[(.*)\\]$")

var valueRegexp = regexp.MustCompile("^(.+?)=(.*)$")

func (cfg *ConfigFile) parse(section string, line string) (string, error) {

	if match := sectionRegexp.FindStringSubmatch(line); match != nil {
		return match[1], nil
	}

	if match := valueRegexp.FindStringSubmatch(line); match != nil {
		key, val := match[1], match[2]
		data := cfg.Values[section]
		if data == nil {
			data = make(SectionData)
			cfg.Values[section] = data
		}
		data[key] = val
		return section, nil
	}

	return section, nil
}
