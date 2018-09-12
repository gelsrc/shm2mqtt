//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"flag"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"strings"
	"time"
)

// Обработка ппараметров утилиты командной строки, чтени конфигурационного файла Logix,
// подключение к MQTT-брокеру и запуск задач слежения.
//
// Usage of ./shm2mqtt:
//   -broker string
//         The broker URI (default "tcp://127.0.0.1:1883")
//   -cfg string
//         Path to logix load_files.srv (default "/projects/load_files.srv")
//   -fullsync int
//         Full publish tick interval, s (default 60)
//   -prefix string
//         Variable label prefix (default "mqtt:")
//   -reconnect int
//         Reconnect pause, s (default 10)
//   -shm string
//         System V shared memory file (default "/dev/shm/wsi")
//   -tick int
//         Publish tick interval, ms (default 100)
//
func tool() int {
	shmFile := flag.String("shm", "/dev/shm/wsi", "System V shared memory file")
	cfgFile := flag.String("cfg", "/projects/load_files.srv", "Path to logix load_files.srv")
	tickMs := flag.Int("tick", 100, "Publish tick interval, ms")
	fullSyncTickS := flag.Int("fullsync", 60, "Full publish tick interval, s")
	reconnect := flag.Int("reconnect", 10, "Reconnect pause, s")
	broker := flag.String("broker", "tcp://127.0.0.1:1883", "The broker URI")
	prefix := flag.String("prefix", "mqtt:", "Variable label prefix")

	flag.Parse()

	cfg := ConfigFile{}

	if err := cfg.Load(*cfgFile); err != nil {
		ERROR.Printf("Configuration load error: %v", err)
		return 1
	}

	var shmSize int

	if r, err := ParseInt(cfg.Values["Slave"]["ShmSize"]); err != nil {
		ERROR.Printf("Can't read shared memory size: %v", err)
		return 1
	} else {
		shmSize = r
	}

	DEBUG.Printf("Use shared memory size %v", shmSize)

	var shm []byte

	if r, err := Shm(*shmFile, shmSize); err != nil {
		ERROR.Printf("Can't map shared memory: %v", err)
	} else {
		shm = r
	}

	shm = shm[24:] // 24 = sizeof(pthread_mutex_t)

	DEBUG.Printf("Use broker: %v", *broker)
	DEBUG.Printf("Use tick: %v", *tickMs)

	var input []InputValue

	if r, err := getInputs(*prefix, cfg, shm); err != nil {
		ERROR.Printf("Values read error: %v", err)
		return 1
	} else {
		input = r
	}

	var output []OutputValue

	if r, err := getOutputs(*prefix, cfg, shm); err != nil {
		ERROR.Printf("Values read error: %v", err)
		return 1
	} else {
		output = r
	}

	for _, v := range output {
		DEBUG.Printf("Export topic: [%v]", v.Topic())
	}

	for _, v := range input {
		DEBUG.Printf("Import topic: [%v]", v.Topic())
	}

	opts := mqtt.NewClientOptions()

	opts.AddBroker(*broker)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)

	for {
		if token := client.Connect(); token.Wait() && token.Error() == nil {
			break
		} else {
			ERROR.Printf("Connect error: %v", token.Error())
		}
		time.Sleep(time.Second * time.Duration(*reconnect))
	}

	inpgw := InputGateway{}

	inpgw.SetClient(client)
	inpgw.SetValues(input)

	outgw := OutputGateway{}

	outgw.SetClient(client)
	outgw.SetTick(time.Millisecond * time.Duration(*tickMs))
	outgw.SetFullSyncTick(time.Second * time.Duration(*fullSyncTickS))
	outgw.SetValues(output)

	if err := outgw.Run(); err != nil {
		ERROR.Printf("OutputGateway error: %v", err)
		return 1
	}

	return 0
}

func getInputs(prefix string, cfg ConfigFile, shm []byte) ([]InputValue, error) {

	var r []InputValue

	var coils []Coil

	if r, err := ExtractCoils(cfg); err != nil {
		return nil, err
	} else {
		coils = r
	}

	for _, v := range coils {
		if !strings.HasPrefix(v.Label, prefix) {
			continue
		}

		label := v.Label[len(prefix):]

		if v.Type == VarTypeBool && v.Size == 1 {
			tmp := NewInputBoolValue(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}
	}

	var holdregs []HoldReg

	if r, err := ExtractHoldregs(cfg); err != nil {
		return nil, err
	} else {
		holdregs = r
	}

	for _, v := range holdregs {
		if !strings.HasPrefix(v.Label, prefix) {
			continue
		}

		label := v.Label[len(prefix):]

		if v.Type == VarTypeInt && v.Size == 2 {
			tmp := NewInputInt16Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}

		if v.Type == VarTypeLong && v.Size == 4 {
			tmp := NewInputInt32Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}

		if v.Type == VarTypeReal && v.Size == 4 {
			tmp := NewInputFloat32Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}
	}

	return r, nil
}

func getOutputs(prefix string, cfg ConfigFile, shm []byte) ([]OutputValue, error) {

	var r []OutputValue

	var instat []InStat

	if r, err := ExtractInStats(cfg); err != nil {
		return nil, err
	} else {
		instat = r
	}

	for _, v := range instat {
		if !strings.HasPrefix(v.Label, prefix) {
			continue
		}

		label := v.Label[len(prefix):]

		if v.Type == VarTypeBool && v.Size == 1 {
			tmp := NewOutputBoolValue(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}
	}

	var inregs []InReg

	if r, err := ExtractInregs(cfg); err != nil {
		return nil, err
	} else {
		inregs = r
	}

	for _, v := range inregs {
		if !strings.HasPrefix(v.Label, prefix) {
			continue
		}

		label := v.Label[len(prefix):]

		if v.Type == VarTypeInt && v.Size == 2 {
			tmp := NewOutputInt16Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}

		if v.Type == VarTypeLong && v.Size == 4 {
			tmp := NewOutputInt32Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}

		if v.Type == VarTypeReal && v.Size == 4 {
			tmp := NewOutputFloat32Value(label, shm, v.Offset)
			r = append(r, &tmp)
			continue
		}
	}

	return r, nil
}

func main() {
	os.Exit(tool())
}
