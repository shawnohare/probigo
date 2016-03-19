package probigo

func bytesEqual(xs, ys []byte) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i, xi := range xs {
		if xi != ys[i] {
			return false
		}
	}
	return true
}
