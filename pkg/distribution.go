package pkg

import (
    "math"
    "math/rand"
    "time"
)

// PoissonInterval genera un intervalo de tiempo con una distribución de Poisson
func PoissonInterval(lambda float64) time.Duration {
    rand.Seed(time.Now().UnixNano())
    interval := -math.Log(1.0-rand.Float64()) / lambda
    return time.Duration(interval * float64(time.Second))
}
