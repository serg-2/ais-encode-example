package main

import (
	"fmt"

	"github.com/serg-2/libs-go/marinelib"
)

func main() {
	var mmsi int = 247320162
	var speed float64 = 0.1
	var longtitude float64 = 5.6
	var latitude float64 = 3.5
	var course float64 = 273.1
	//Timestamp (sec) integer
	var ts int = 1

	binstr := marinelib.EncodeType1(mmsi, speed, longtitude, latitude, course, ts)
	fmt.Println(marinelib.GenerateNMEA(binstr))
}
