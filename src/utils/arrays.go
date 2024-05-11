package utils

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func ConvertArrayItems[T, K Number](arr []T) []K {
	var newArr []K
	for _, item := range arr {
		newArr = append(newArr, K(item))
	}

	return newArr
}
