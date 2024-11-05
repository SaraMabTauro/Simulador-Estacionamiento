package shared

import (
    "math"
    "math/rand"
)

// PoissonInterval genera un intervalo de tiempo con distribución de Poisson
func PoissonInterval(lambda float64) float64 {
    return -math.Log(1.0-rand.Float64()) / lambda
}
