package main

import (
	"fmt"
	"os"
	"parking-simulator/internal/domain"
	"parking-simulator/internal/infrastructure/gui"
	usecase "parking-simulator/internal/usercase"
	"sync"
)

func init() {
	os.Setenv("FYNE_RENDERER", "software")
}

func main() {
	fmt.Println("Iniciando simulación del estacionamiento...")

	//     // Crear el estacionamiento con capacidad de 20 vehículos
	parkingLot := domain.NewParkingLot(20)
	wg := &sync.WaitGroup{}

	//     // Iniciar la interfaz gráfica en una goroutine
	go gui.StartGUI(parkingLot)

	// Iniciar la simulación de llegada de vehículos
	usecase.StartSimulation(parkingLot, wg)

	// Esperar a que todas las goroutines finalicen
	wg.Wait()

	fmt.Println("Simulación finalizada.")
}
