package util

func GetKeysOfMap[K comparable, T any](maps map[K]T) []K {
	arr := make([]K, len(maps))
	i := 0
	for k, _ := range maps {
		arr[i] = k
		i++
	}
	return arr
}

func GetValuesOfMap[K comparable, T any](maps map[K]T) []T {
	arr := make([]T, len(maps))
	i := 0
	for _, v := range maps {
		arr[i] = v
		i++
	}
	return arr
}
