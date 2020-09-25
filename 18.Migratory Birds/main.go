package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func exists(a int32, arr []int32) bool {
	for _, x := range arr {
		if x == a {
			return true
		}
	}
	return false
}

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
	var (
		sum     int32
		tmp_sum int32
		max     int32
		tmp_max int32
	)
	tmp_arr := []int32{}
	for _, x := range arr {
		val := exists(x, tmp_arr)
		if val == false {
			tmp_arr = append(tmp_arr, x)
		}
	}

	for c, y := range tmp_arr {
		for _, x := range arr {
			if x == y {
				tmp_sum++
			}
		}
		tmp_max = y

		if c == 0 {
			sum = tmp_sum
			max = tmp_max
		}

		if sum <= tmp_sum {
			if sum == tmp_sum {
				if max < tmp_max {
				} else {
					sum = tmp_sum
					max = tmp_max
				}
			} else {
				sum = tmp_sum
				max = tmp_max
			}

		}

		tmp_sum = 0

	}
	return max

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
