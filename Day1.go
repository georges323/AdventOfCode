package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	sum float64 = 0
	max float64 = 0
	elfNumber int = 0
)

func main() {
    f, err := os.Open("data/Day1.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		calorie := scanner.Text()

		if calorie == "" {
			max = math.Max(sum, max)
			sum = 0

			elfNumber += 1
		} else {
			parsedCalorie, err := strconv.Atoi(calorie)

			if err != nil {
				fmt.Printf("Error: Cannot parse %v", calorie)
				return
			}

			sum += float64(parsedCalorie)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal((err))
	}

	fmt.Printf("Elf number %v carries the most calories of %v", elfNumber, max)
}
