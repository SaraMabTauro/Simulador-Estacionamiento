package gui

import (
	"fmt"
	"os"
	"parking-simulator/internal/domain"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartGUI(parkingLot *domain.ParkingLot) {
	// Asegurarse de usar el renderer de software
	os.Setenv("FYNE_RENDERER", "software")

	application := app.New()
	window := application.NewWindow("Parking Simulator")

	// Crear widgets
	status := widget.NewLabel("Initializing...")

	// Actualizar estado
	updateStatus := func() {
		status.SetText(fmt.Sprintf("Vehículos estacionados: %d/%d",
			len(parkingLot.Vehicles), parkingLot.Capacity))
	}

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Estado del Estacionamiento"),
		status,
	)

	// Actualización periódica
	go func() {
		for {
			time.Sleep(time.Second)
			updateStatus()
		}
	}()

	// Configurar ventana
	window.SetContent(content)
	window.Resize(fyne.NewSize(300, 200))

	// Actualización inicial
	updateStatus()

	// Mostrar y ejecutar
	window.ShowAndRun()
}
