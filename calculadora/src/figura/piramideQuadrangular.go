package figura

import (
	"fmt"
	"math"
)

type PiramideQuadrangular struct {
	base, altura float64
}

func (piramide PiramideQuadrangular) CalcularArea() float64{
	var areaBase  float64 =  piramide.base * piramide.base

	//cálculo da área lateral
	var apotemaBase float64 = (piramide.base / 2)
	var apotemaPiramide float64 = math.Pow(apotemaBase, 2)+math.Pow(piramide.altura, 2)
	apotemaPiramide = (math.Sqrt(apotemaPiramide) * piramide.base) / 2
	var areaLateral float64 = 4 * apotemaPiramide

	return areaBase + areaLateral
}

func (piramide PiramideQuadrangular) CalcularVolume() float64{
	const VALUE_PIRAMIDE  = float64(1)/float64(3)
	var areaBase  float64 =  piramide.base * piramide.base
	return VALUE_PIRAMIDE *areaBase*piramide.altura
}

func (piramide PiramideQuadrangular) Calcular() {
	fmt.Print("Digite o tamanho da aresta base da piramide:")
	fmt.Scanf("%f\n", &piramide.base)
	fmt.Print("Digite o tamanho da altura da piramide:")
	fmt.Scanf("%f\n", &piramide.altura)

	fmt.Println("Área do retângulo:",piramide.CalcularArea())
	fmt.Println("Perímetro do retângulo:",piramide.CalcularVolume())
}
