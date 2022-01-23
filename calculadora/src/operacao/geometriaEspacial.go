package operacao

type GeometriaEspacial interface {
	CalcularArea() float64
	CalcularVolume() float64
	Calcular()
}

func DescribeFieldsEspacial(i GeometriaEspacial) {
	i.Calcular();
}