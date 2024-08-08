package readData

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-1/functionFiles"
)

func Reader() {
	var numslice []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		numInt := scanner.Text()
		numFloat, err := strconv.ParseFloat(numInt, 64)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		numslice = append(numslice, numFloat)
		if len(numslice) > 1 {
			lowerLimit, upperLimit := functionFiles.Range(numslice)
			fmt.Printf("%d %d\n", lowerLimit, upperLimit)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
