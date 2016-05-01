package main

import (
	"fmt"
	"github.com/lemonlatte/aqi-go/aqi"
)

func main() {
	fmt.Println("IAQI: ", aqi.GetAQI(aqi.SO2, 305))
}
