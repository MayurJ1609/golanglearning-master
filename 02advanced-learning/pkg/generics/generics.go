package generics

type Number interface {
	int | int32 | int64 | float32 | float64
}

func SumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}
