package main

import "fmt"

type falcon9Rocket struct {
	payload []interface{}
}

func (f9r *falcon9Rocket) insertSatelliteIntoStarlinkPort(s satellite) {
	fmt.Println("Attaching satellite to Falcon 9 Rocket.")
	s.insertSatelliteIntoStarlinkPort(f9r)
}

func (f9r *falcon9Rocket) getPayload() {
	fmt.Printf("Falcon 9 payload: %s\n", f9r.payload)
}

type satellite interface {
	insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket)
}

type starlinkSatellite struct {
	name string
}
type oco2Satellite struct {
	name string
}

func newStarlinkSatellite() *starlinkSatellite {
	return &starlinkSatellite{
		name: "Starlink Satellite",
	}
}
func newOco2Satellite() *oco2Satellite {
	return &oco2Satellite{
		name: "OCO2 Satellite",
	}
}

func (sls *starlinkSatellite) insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket) {
	f9r.payload = append(f9r.payload, sls)
	fmt.Println("Starlink satellite is attached to Falcon 9 Rocket.")
}
func (oco2s *oco2Satellite) insertSatelliteIntoOco2Port(f9r *falcon9Rocket) {
	f9r.payload = append(f9r.payload, oco2s)
	fmt.Println("OCO2 satellite is attached to Falcon 9 Rocket.")
}

type oco2SatelliteAdapter struct {
	oco2Satellite *oco2Satellite
}

func (oco2sa *oco2SatelliteAdapter) insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket) {
	fmt.Println("Satellite adapter converts Starlink port to OCO2 port.")
	oco2sa.oco2Satellite.insertSatelliteIntoOco2Port(f9r)
}

func main() {

	//Допишите main используя пример к уроку
	falcon9Rocket := &falcon9Rocket{}
	falcon9Rocket.getPayload()

	starlinkSatellite := newStarlinkSatellite()
	oco2Satellite := newOco2Satellite()

	co2Satellite_adapter := &oco2SatelliteAdapter{oco2Satellite: oco2Satellite}

	falcon9Rocket.insertSatelliteIntoStarlinkPort(starlinkSatellite)
	falcon9Rocket.insertSatelliteIntoStarlinkPort(co2Satellite_adapter)

	falcon9Rocket.getPayload()
}

/**
Последовательность:

Falcon 9 payload: []
Attaching satellite to Falcon 9 Rocket.
Starlink satellite is attached to Falcon 9 Rocket.
Attaching satellite to Falcon 9 Rocket.
Satellite adapter converts Starlink port to OCO2 port.
OCO2 satellite is attached to Falcon 9 Rocket.
Falcon 9 payload: [%!s(*main.starlinkSatellite=&{Starlink Satellite}) %!s(*main.oco2Satellite=&{OCO2 Satellite})]
**/
