package sort

func selectionsort(a []int) {
	for i := range a {
		minIndex := i
		swapped := false

		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
				swapped = true
			}
		}

		if !swapped {
			break
		}

		a[minIndex], a[i] = a[i], a[minIndex]
	}
}
