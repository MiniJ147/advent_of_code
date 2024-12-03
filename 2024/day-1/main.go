package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	buff1 := []int{}
	buff2 := []int{}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := strings.Split(line, " ")

		num1, err := strconv.Atoi(parsed[0])
		num2, err2 := strconv.Atoi(parsed[3])

		if err != nil || err2 != nil {
			panic(errors.Join(err, err2))
		}

		buff1 = append(buff1, num1)
		buff2 = append(buff2, num2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(buff1)
	sort.Ints(buff2)

	ans := 0
	for i := 0; i < len(buff1); i++ {
		dist := math.Abs(float64(buff1[i] - buff2[i]))
		ans += int(dist)
	}
	fmt.Println(ans)
}
