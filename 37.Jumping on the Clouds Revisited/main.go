package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32, k int32) int32 {
	var sum int32
	var v1 int32
	var i int32
	for i != int32(len(c)) {
		i += k
		if i == int32(len(c)) {
			break
		} else if i >= int32(len(c)) {
			i = i - int32(len(c))
			row := i
			v1 = c[row]
		} else {
			//fmt.Println(c[i])
			v1 = c[i]
		}
		if v1 == 1 {
			sum = sum + 3
		} else {
			sum = sum + 1
		}
	}
	val := c[0]
	if val == 1 {
		sum = sum + 3
	} else {
		sum = sum + 1
	}
	return 100 - sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c, k)

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
