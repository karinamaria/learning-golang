package figura

import (
	"fmt"
)

type Retangulo struct {
	base, altura float64
}

func (retangulo Retangulo) CalcularArea() float64{
	return retangulo.base * retangulo.altura
}

func (retangulo Retangulo) CalcularPerimetro() float64{
	return 2*(retangulo.base + retangulo.altura)
}

func (retangulo Retangulo) Calcular() {
	fmt.Print("Digite o tamanho da base do retângulo:")
	fmt.Scanf("%f\n", &retangulo.base)
	fmt.Print("Digite o tamanho da altura do retângulo:")
	fmt.Scanf("%f\n", &retangulo.altura)

	fmt.Println("Área do retângulo:",retangulo.CalcularArea())
	fmt.Println("Perímetro do retângulo:",retangulo.CalcularPerimetro())
}
