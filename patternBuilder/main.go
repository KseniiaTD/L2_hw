package main

import "fmt"

func main() {
	pickupTruck := createCar("pickup truck")
	roadster := createCar("roadster")
	sedan := createCar("sedan")

	fmt.Printf("pickupTruck: %v\n", pickupTruck)
	fmt.Printf("roadster: %v\n", roadster)
	fmt.Printf("sedan: %v\n", sedan)

}
