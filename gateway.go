//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"time"
)

// Шлюз для публикуемых значений.
type OutputGateway struct {
	client       mqtt.Client
	tickWait     Wait
	fullSyncWait Wait
	values       []OutputValue
}

func (gw *OutputGateway) SetClient(client mqtt.Client) {
	gw.client = client
}

func (gw *OutputGateway) SetTick(tick time.Duration) {
	gw.tickWait.Setup(tick, tick*10)
}

func (gw *OutputGateway) SetFullSyncTick(tick time.Duration) {
	gw.fullSyncWait.Setup(tick, tick*10)
}

func (gw *OutputGateway) SetValues(values []OutputValue) {
	gw.values = values
}

func (gw *OutputGateway) Run() error {

	gw.tickWait.Reset()
	gw.fullSyncWait.Reset()

	waitTokens := make([]mqtt.Token, len(gw.values))

	for {
		gw.fullSyncWait.Correct()
		fullSync := gw.fullSyncWait.After()

		if fullSync {
			DEBUG.Printf("Time to full sync.")
			gw.fullSyncWait.Step()
		}

		updates := 0

		for _, v := range gw.values {
			if v.Sync() || fullSync {
				waitTokens[updates] = gw.client.Publish(
					v.Topic(),
					0,
					false,
					v.String(),
				)
				updates++
			}
		}

		if updates > 0 {
			for i := 0; i < updates; i++ {
				token := waitTokens[i]
				token.Wait()
				waitTokens[i] = nil
			}
			DEBUG.Printf("Updated %v values.", updates)
		}

		gw.tickWait.Step()
		gw.tickWait.Correct()
		gw.tickWait.Wait()
	}
}

// Шлюз для отслеживаемых значений.
type InputGateway struct {
	client mqtt.Client
	values map[string]InputValue
}

func (gw *InputGateway) SetClient(client mqtt.Client) {
	gw.client = client
}

func (gw *InputGateway) SetValues(values []InputValue) {
	gw.values = make(map[string]InputValue)
	for _, v := range values {
		topic := v.Topic()
		gw.values[topic] = v
		gw.client.Subscribe(topic, 2, gw.handleMassagefunc)
	}
}

func (gw *InputGateway) handleMassagefunc(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	if v, ok := gw.values[topic]; ok {
		s := string(msg.Payload())
		if err := v.Apply(s); err != nil {
			ERROR.Printf("Bad value [%v] from [%v]: %v.", s, topic, err)
		} else {
			DEBUG.Printf("Incoming value [%v] from [%v]", s, topic)
		}
	} else {
		ERROR.Printf("Unexpected topic [%v].", topic)
	}
}
