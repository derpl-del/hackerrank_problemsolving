package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func cutTheSticks(arr []int32) []int32 {
	var result []int32
	for len(arr) != 0 {
		x := 0
		c := 0
		fmt.Println(arr)
		min := getMin(arr)
		for x < len(arr) {

			arr[int32(x)] = arr[int32(x)] - min
			if arr[int32(x)] == 0 {
				arr = append(arr[:int32(x)], arr[int32(x)+1:]...)
			} else {
				x++
			}
			c++
		}
		result = append(result, int32(c))
	}
	return result
}

func getMin(arr []int32) int32 {
	var min int32
	min = 9999
	for _, x := range arr {
		min = int32(math.Min(float64(min), float64(x)))
	}
	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
