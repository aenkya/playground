package utils

func CompareSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}

// TODO: Update codebase to use sort.Slice instead of this
func Contains[T comparable](m map[T]bool, val T) bool {
	if _, ok := m[val]; ok {
		return true
	}

	return false
}
