package algo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NoOP = 1
	Add  = 2
)

func adventOfCode10(filename string) int {
	file, err := os.Open("./input/" + filename)
	if err != nil {
		fmt.Println(err)
		panic("Error opening file")
	}

	defer file.Close()

	clockCycle := 1
	signalStrength := 0
	registerValue := 1
	renderContent := make([][]string, 6)

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			renderContent[i] = append(renderContent[i], "░░")
		}
	}

	lines := 0
	checkCycles := []int{20, 60, 100, 140, 180, 220}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(checkCycles) == 0 {
			break
		}

		line := scanner.Text()
		lines++

		lineSegments := strings.Split(line, " ")

		if len(lineSegments) == 2 {
			renderContent = fillRenderer(renderContent, clockCycle, registerValue)
			renderContent = fillRenderer(renderContent, clockCycle+1, registerValue)

			clockCycle += 2
			num, err := strconv.Atoi(lineSegments[1])
			if err != nil {
				panic("Error converting string to int")
			}

			if clockCycle >= checkCycles[0] {
				signalStrength += (registerValue * checkCycles[0])
				checkCycles = checkCycles[1:]
			}

			registerValue += num
		} else {
			renderContent = fillRenderer(renderContent, clockCycle, registerValue)
			clockCycle++

			if clockCycle >= checkCycles[0] {
				signalStrength += (registerValue * checkCycles[0])
				checkCycles = checkCycles[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error reading file")
	}

	fmt.Println("Total clock cycles: ", clockCycle)
	renderToScreen(renderContent)

	return signalStrength
}

func fillRenderer(renderContent [][]string, clockCycle, registerValue int) [][]string {
	renderVal := "░░"

	tracker := clockCycle % 40
	if tracker == registerValue || tracker == registerValue+1 || tracker == registerValue+2 {
		renderVal = "██"
	}

	if (clockCycle % 40) == 0 {
		renderContent[(clockCycle/40)-1][39] = renderVal
	} else {
		renderContent[clockCycle/40][clockCycle%40] = renderVal
	}

	return renderContent
}

func renderToScreen(content [][]string) {
	for _, row := range content {
		for _, col := range row {
			fmt.Print(col)
		}

		fmt.Println()
	}
}
