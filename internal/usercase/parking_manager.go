package usecase

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"parking-simulator/internal/domain"
)

// Inicia la simulacion de vehiculos que llegan al estacionamiento
func StartSimulation(parkingLot *domain.ParkingLot, wg *sync.WaitGroup) {
	entryGate := make(chan struct{}, 1) //canalcito en donde se controla la entrada/salida(recurso compartido)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(vehicleID int) {
			defer wg.Done()
			v := domain.NewVehicle(vehicleID)

			//Simula la llegada en distribucion
			time.Sleep(time.Duration(rand.ExpFloat64()/1.5) * time.Second)

			fmt.Printf("Vehiculo %d intenta entrar al estacionamiento.\n", vehicleID)

			//Controla el acceso al recurso compartido(puerta)
			entryGate <- struct{}{}
			if parkingLot.Enter(v) {
				<-entryGate //libera puerta
				v.SimulateStay()

				//acceso al recurso compartido para salir
				entryGate <- struct{}{}
				parkingLot.Exit(v)
				<-entryGate
			} else {
				<-entryGate //libera la puerta sino puede entrar
				fmt.Printf("Vehiculo %d no pudo entrar , estacionamiento lleno.\n", vehicleID)
			}
		}(i)
	}

}
