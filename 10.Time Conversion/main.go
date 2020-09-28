package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */ //
func timeConversion(s string) string {
	var a string
	t := s[8:10]
	if t == "PM" {
		a = s[0:2]
		i, _ := strconv.Atoi(a)
		r := 12 + i
		if r >= 24 {
			a = "12"
		} else {
			a = fmt.Sprint(r)
		}
	} else {
		a = s[0:2]
		i, _ := strconv.Atoi(a)
		if i >= 12 {
			a = "00"
		} else {
			a = s[0:2]
		}

	}
	t2 := s[2:10]
	t3 := t2[0:6]
	return fmt.Sprintf("%v%v", a, t3)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
