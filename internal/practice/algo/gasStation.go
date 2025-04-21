package algo

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank, currentTank, start := 0, 0, 0

	for i := range len(gas) {
		diff := gas[i] - cost[i]
		totalTank += diff
		currentTank += diff

		if currentTank < 0 {
			start = i + 1
			currentTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}

	return start
}
