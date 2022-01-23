package figura

import (
	"fmt"
	"math"
)

type TrianguloEquilatero struct {
	lado float64
}

func (triangulo TrianguloEquilatero) CalcularArea() float64{
	return (math.Pow(triangulo.lado, 2)*math.Sqrt(3))/4
}

func (triangulo TrianguloEquilatero) CalcularPerimetro() float64{
	return 3*triangulo.lado
}

func (triangulo TrianguloEquilatero) Calcular() {
	fmt.Print("Digite o tamanho do lado do triângulo equilátero:")
	fmt.Scanf("%f\n", &triangulo.lado)

	fmt.Println("Área do triângulo equilátero:",triangulo.CalcularArea())
	fmt.Println("Perímetro do triângulo equilátero:",triangulo.CalcularPerimetro())
}
