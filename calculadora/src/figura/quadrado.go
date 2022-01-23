package figura

import (
	"fmt"
)

type Quadrado struct {
	lado float64
}

func (quadrado Quadrado) CalcularArea() float64{
	return quadrado.lado * quadrado.lado
}

func (quadrado Quadrado) CalcularPerimetro() float64{
	return 4*quadrado.lado
}

func (quadrado Quadrado) Calcular() {
	fmt.Print("Digite o tamanho do lado do quadrado:")
	fmt.Scanf("%f\n", &quadrado.lado)

	fmt.Println("Área do quadrado:",quadrado.CalcularArea())
	fmt.Println("Perímetro do quadrado:",quadrado.CalcularPerimetro())
}
