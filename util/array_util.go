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
