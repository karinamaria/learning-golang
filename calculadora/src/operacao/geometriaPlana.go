package operacao

type GeometriaPlana interface {
	CalcularArea() float64
	CalcularPerimetro() float64
	Calcular()
}

func DescribeFields(i GeometriaPlana) {
	i.Calcular();
}