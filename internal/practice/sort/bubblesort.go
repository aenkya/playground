package sort

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		swapped := false

		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
