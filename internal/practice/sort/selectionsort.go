package sort

func selectionsort(a []int) {

	for i := range a {
		min_index := i
		swapped := false

		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min_index] {
				min_index = j
				swapped = true
			}
		}

		if !swapped {
			break
		}

		a[min_index], a[i] = a[i], a[min_index]
	}
}
