package figura

import (
	"fmt"
	"math"
)
const PI float64 = math.Pi

type Circulo struct {
	raio float64
}

func (circulo Circulo) CalcularArea() float64{
	return PI * math.Pow(circulo.raio, 2)
}

func (circulo Circulo) CalcularPerimetro() float64{
	return 2*PI*circulo.raio
}

func (circulo Circulo) Calcular() {
	fmt.Print("Digite o tamanho do raio do circulo:")
	fmt.Scanf("%f\n", &circulo.raio)

	fmt.Println("Área do circulo:",circulo.CalcularArea())
	fmt.Println("Perímetro do circulo:",circulo.CalcularPerimetro())
}
