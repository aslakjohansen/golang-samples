package main

import (
    "fmt"
)

type car struct {
    gas_pedal       uint16
    brake_pedal     uint16
    steering_wheel   int16
    top_speed_kph  float64
}

func main () {
    var a_car car = car{gas_pedal:      22341,
                        brake_pedal:        0,
                        steering_wheel: 12561,
                        top_speed_kph:    225.0}
    b_car := car{22341, 0, 12561, 225.0} // same as above
    
    fmt.Println(a_car.gas_pedal)
    
    b_car_ref := &b_car
    fmt.Println(b_car_ref.gas_pedal)
}

