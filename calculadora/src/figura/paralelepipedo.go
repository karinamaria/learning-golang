package figura

import (
	"fmt"
)

type Paralelepipedo struct {
	aresta1, aresta2, aresta3 float64
}

func (paralelepipedo Paralelepipedo) CalcularArea() float64{
	var primeiraFace float64 = 2 * paralelepipedo.aresta1 * paralelepipedo.aresta2
	var segundaFace float64 = 2 * paralelepipedo.aresta1 * paralelepipedo.aresta3
	var terceiraFace float64 = 2 * paralelepipedo.aresta2 * paralelepipedo.aresta3
	return primeiraFace + segundaFace + terceiraFace
}

func (paralelepipedo Paralelepipedo) CalcularVolume() float64{
	return paralelepipedo.aresta1 * paralelepipedo.aresta2 * paralelepipedo.aresta3
}

func (paralelepipedo Paralelepipedo) Calcular() {
	fmt.Print("Digite o tamanho da aresta1 do retângulo:")
	fmt.Scanf("%f\n", &paralelepipedo.aresta1)
	fmt.Print("Digite o tamanho da aresta2 do retângulo:")
	fmt.Scanf("%f\n", &paralelepipedo.aresta2)
	fmt.Print("Digite o tamanho da aresta3 do retângulo:")
	fmt.Scanf("%f\n", &paralelepipedo.aresta3)

	fmt.Println("Área do retângulo:",paralelepipedo.CalcularArea())
	fmt.Println("Perímetro do retângulo:",paralelepipedo.CalcularVolume())
}
