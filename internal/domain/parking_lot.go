package domain

import (
	"fmt"
	"sync"
)

type ParkingLot struct {
	Capacity int
	Vehicles chan *Vehicle
	Mutex    sync.Mutex
}

// Se crea un nuevo estacionamiento con la capacidad que se especifica
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity: capacity,
		Vehicles: make(chan *Vehicle, capacity),
	}
}

// Intenta agregar un carrito al estacionamiento
func (p *ParkingLot) Enter(v *Vehicle) bool {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if len(p.Vehicles) < p.Capacity {
		p.Vehicles <- v
		fmt.Printf("Vehiculo %d ha entrado al estacionamiento.\n", v.ID)
		return true
	} else {
		fmt.Printf("Vehiculo %d no pudo entrar, estacionamiento lleno.\n", v.ID)
		return false
	}
}

// Permite que un carro salga
func (p *ParkingLot) Exit(v *Vehicle) *Vehicle {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if len(p.Vehicles) > 0 {
		vehicle := <-p.Vehicles
		fmt.Printf("Vehiculo %d ha salido del estacionaiento.\n", vehicle.ID)
		return vehicle
	}
	return nil
}
