package main

import (
	"fmt"
	"calculadora/src/operacao"
	"calculadora/src/figura"
)

func main() {
	var geometriaPlana operacao.GeometriaPlana
	var geometriaEspacial operacao.GeometriaEspacial

	var option int = -1
	for option != 0 {
		fmt.Println("Calculadora de Geometria Plana e Espacial")
		fmt.Println("(1) Triângulo equilátero")
		fmt.Println("(2) Retângulo")
		fmt.Println("(3) Quadrado")
		fmt.Println("(4) Circulo")
		fmt.Println("(5) Pirâmide com base quadrangular")
		fmt.Println("(6) Cubo")
		fmt.Println("(7) Paralelepipedo")
		fmt.Println("(8) Esfera")
		fmt.Println("(0) Sair")
		fmt.Print("Digite a sua opção:")
		fmt.Scanf("%d\n", &option)

		switch option {
			case 0:
				fmt.Printf("Saindo...")
			case 1:
				geometriaPlana = figura. TrianguloEquilatero{}
				operacao.DescribeFields(geometriaPlana)
			case 2:
				geometriaPlana = figura. Retangulo{}
				operacao.DescribeFields(geometriaPlana)
			case 3:
				geometriaPlana = figura. Quadrado{}
				operacao.DescribeFields(geometriaPlana)
			case 4:
				geometriaPlana = figura. Circulo{}
				operacao.DescribeFields(geometriaPlana)
			case 5:
				geometriaEspacial = figura. PiramideQuadrangular{}
				operacao.DescribeFieldsEspacial(geometriaEspacial)
			case 6:
				geometriaEspacial = figura. Cubo{}
				operacao.DescribeFieldsEspacial(geometriaEspacial)
			case 7:
				geometriaEspacial = figura. Paralelepipedo{}
				operacao.DescribeFieldsEspacial(geometriaEspacial)
			case 8:
				geometriaEspacial = figura. Esfera{}
				operacao.DescribeFieldsEspacial(geometriaEspacial)
			default:
				fmt.Println("Escolha uma opção válida")
		}
	}
}
