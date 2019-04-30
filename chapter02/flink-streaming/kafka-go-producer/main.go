package main

import (
	"fmt"
	"time"
)
import "math/rand"

func main(){
	for{
		stamp := time.Now().Unix()
		Producer(fmt.Sprintf("%d, %.2f, %d", stamp,generateTemperature(),generateSensorId()))
		time.Sleep(time.Second*1)
	}
}

func generateTemperature() float64 {
	rand.Seed(time.Now().Unix())
	return 30 * rand.Float64()
}


func generateSensorId() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(100)
}