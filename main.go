package main

import "fmt"

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0, 1}

	gridLoad := 75.

	menuOption := menu()

	switch menuOption {
	case "1":
		generatePlantCapacityReport(plantCapacities...)
	case "2":
		generateGridLoadReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknown option, no option taken")
	}
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func generateGridLoadReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	fmt.Printf("%-20s%.0f\n", "capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)
}

func menu() (option string) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")
	fmt.Scanln(&option)
	return
}
