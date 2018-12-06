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

// (method) value receiver: works on a copy of the structure
func (c car) kmh () float64 {
    return float64(c.gas_pedal) * (c.top_speed_kph/usixteenbitmax)
}
func (c car) mph () float64 {
    return float64(c.gas_pedal) * (c.top_speed_kph/usixteenbitmax/kmh_multiple)
}

// (method) pointer receiver: works on a reference to the structure
func (c *car) new_top_speed (newspeed float64) {
    c.top_speed_kph = newspeed
}

// plain function
func newer_top_speed (c car, speed float64) car {
    c.top_speed_kph = speed
    return c
}

func main () {
    var a_car car = car{gas_pedal:      65000,
                        brake_pedal:        0,
                        steering_wheel: 12561,
                        top_speed_kph:    225.0}
    
    fmt.Println(a_car.gas_pedal)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
    
    a_car.new_top_speed(500)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
    
    a_car = newer_top_speed(a_car, 300)
    fmt.Println(a_car.kmh())
    fmt.Println(a_car.mph())
}

