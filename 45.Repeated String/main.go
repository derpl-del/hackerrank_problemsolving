package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	if len(s) == 1 {
		if s == "a" {
			return n
		} else {
			return 0
		}
	}
	s1 := int64(strings.Count(s, "a"))
	l1 := int64(len(s))
	l2 := n % l1
	sum := int64(0)
	c := int64(n / l1)
	r := []byte(s)
	for x := 0; x < int(l2); x++ {
		val := r[x]
		if val == 97 {
			sum++
		}
	}
	return (s1 * c) + sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
