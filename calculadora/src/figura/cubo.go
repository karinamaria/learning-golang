package figura

import (
	"fmt"
	"math"
)

type Cubo struct {
	aresta float64
}

func (cubo Cubo) CalcularArea() float64{
	return 6 * math.Pow(cubo.aresta, 2)
}

func (cubo Cubo) CalcularVolume() float64{
	return math.Pow(cubo.aresta, 3)
}

func (cubo Cubo) Calcular() {
	fmt.Print("Digite o tamanho da aresta do cubo:")
	fmt.Scanf("%f\n", &cubo.aresta)

	fmt.Println("Área do cubo:",cubo.CalcularArea())
	fmt.Println("Perímetro do cubo:",cubo.CalcularVolume())
}
