package main

import "fmt"

// AirConditioner is the interface that defines the methods for controlling the air conditioner
type AirConditioner interface {
	TurnOn()
	TurnOff()
	SetTemperature(temp int)
}

func (airC *RealAirConditioner) TurnOff()                { airC.work = false }
func (airC *RealAirConditioner) TurnOn()                 { airC.work = true }
func (airC *RealAirConditioner) SetTemperature(temp int) { airC.temperature = temp }

// RealAirConditioner is the implementation of the air conditioner
type RealAirConditioner struct {
	work        bool
	temperature int
}

type AirConditionerProxy struct {
	AirConditioner AirConditioner
	Auth           bool
	Logger         []string
}

func NewAirConditionerProxy(auth bool) *AirConditionerProxy {
	return &AirConditionerProxy{AirConditioner: &RealAirConditioner{}, Auth: auth, Logger: make([]string, 0)}
}

func (proxy *AirConditionerProxy) TurnOff() {
	if !proxy.Auth {
		proxy.Logger = append(proxy.Logger, "Access denied: authentication required to turn off the air condition")
		return
	}
	proxy.AirConditioner.TurnOff()
	proxy.Logger = append(proxy.Logger, "Turn off the air condition")
}
func (proxy *AirConditionerProxy) TurnOn() {
	if !proxy.Auth {
		proxy.Logger = append(proxy.Logger, "Access denied: authentication required to turn on the air condition")
		return
	}
	proxy.AirConditioner.TurnOn()
	proxy.Logger = append(proxy.Logger, "Turn on the air condition")
}
func (proxy *AirConditionerProxy) SetTemperature(temp int) {
	if !proxy.Auth {
		proxy.Logger = append(proxy.Logger, "Access denied: authentication required to set temperature of the air condition")
		return
	}
	proxy.AirConditioner.SetTemperature(temp)
	proxy.Logger = append(proxy.Logger, "Setting air condition temperature to "+fmt.Sprintf("%d", temp))
}

func (proxy *AirConditionerProxy) ShowLog() {
	for _, v := range proxy.Logger {
		fmt.Println(v)
	}
}

func main() {
	airConditioner := NewAirConditionerProxy(false) // without auth
	airConditioner.TurnOn()
	airConditioner.TurnOff()
	airConditioner.SetTemperature(25)

	airConditioner.ShowLog()

	airConditioner = NewAirConditionerProxy(true) // with auth
	airConditioner.TurnOn()
	airConditioner.TurnOff()
	airConditioner.SetTemperature(25)

	airConditioner.ShowLog()
}
