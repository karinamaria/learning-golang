package figura

import (
	"fmt"
	"math"
)

type Esfera struct {
	raio float64
}

func (esfera Esfera) CalcularArea() float64{
	return 4 * PI * math.Pow(esfera.raio, 2)
}

func (esfera Esfera) CalcularVolume() float64{
	const VALUE_ESFERA  = float64(4)/float64(3)
	return VALUE_ESFERA * PI * math.Pow(esfera.raio, 3)
}

func (esfera Esfera) Calcular() {
	fmt.Print("Digite o tamanho do raio do esfera:")
	fmt.Scanf("%f\n", &esfera.raio)

	fmt.Println("Área do esfera:",esfera.CalcularArea())
	fmt.Println("Perímetro do esfera:",esfera.CalcularVolume())
}
