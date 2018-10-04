package chess

func Contains(arr []uint8, v uint8) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}
