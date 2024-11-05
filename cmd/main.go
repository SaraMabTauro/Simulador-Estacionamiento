// package main

// import (
//     "fmt"
//     "sync"
// 	"os"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"


//     "parking-simulator/internal/domain"
//     "parking-simulator/internal/infrastructure/gui"
//     "parking-simulator/internal/usercase"
// )

// func init() {
//     os.Setenv("FYNE_RENDERER", "software")
// }

// func main() {
//     fmt.Println("Iniciando simulación del estacionamiento...")

//     // Crear el estacionamiento con capacidad de 20 vehículos
//     parkingLot := domain.NewParkingLot(20)
//     wg := &sync.WaitGroup{}

//     // Iniciar la interfaz gráfica en una goroutine
//     go gui.RunGUI(parkingLot, wg)

//     // Iniciar la simulación de llegada de vehículos
//     usecase.StartSimulation(parkingLot, wg)

//     // Esperar a que todas las goroutines finalicen
//     wg.Wait()

//     fmt.Println("Simulación finalizada.")
// }

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}