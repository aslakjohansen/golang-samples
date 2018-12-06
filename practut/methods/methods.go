package main

import (
    "fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
    gas_pedal       uint16
    brake_pedal     uint16
    steering_wheel   int16
    top_speed_kph  float64
}

// value receiver associated with struct (aka method)
func (c car) kmh () float64 {
    return float64(c.gas_pedal) * (c.top_speed_kph/usixteenbitmax)
}
func (c car) mph () float64 {
    return float64(c.gas_pedal) * (c.top_speed_kph/usixteenbitmax/kmh_multiple)
}

func main () {
    var a_car car = car{gas_pedal:      65000,
                        brake_pedal:        0,
                        steering_wheel: 12561,
                        top_speed_kph:    225.0}
    
    fmt.Println(a_car.gas_pedal)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
}

