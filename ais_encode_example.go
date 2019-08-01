package main

import (
	"fmt"

	"github.com/serg-2/libs-go/marinelib"
)

func main() {
	//TYPE 1
	var mmsi int = 247320162
	var speed float64 = 0.1
	var longtitude float64 = 5.6
	var latitude float64 = 3.5
	var course float64 = 273.1
	//Timestamp (sec) integer
	var ts int = 1
	
	fmt.Println("TYPE 1")
	binstr := marinelib.EncodeType1(mmsi, speed, longtitude, latitude, course, ts)
	fmt.Println(marinelib.GenerateNMEA(binstr))
	
	//TYPE 24A
	fmt.Println("TYPE 24A")
	vesselName := "VSHIP1"
	binstr = marinelib.EncodeType24(mmsi, "A", vesselName , "", "", 0)
	fmt.Println(marinelib.GenerateNMEA(binstr))

	
	callSign := "CASN1"
	vesselDimensions := "20x10"

	//Vessel Type and Cargo
	passengers := 100
	
	//TYPE 24B
	fmt.Println("TYPE 24B")
	binstr = marinelib.EncodeType24(mmsi, "B", "" , callSign, vesselDimensions, passengers)
	fmt.Println(marinelib.GenerateNMEA(binstr))	

}
