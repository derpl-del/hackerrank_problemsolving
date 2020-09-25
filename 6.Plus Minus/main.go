package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var (
		p  int32
		n  int32
		z  int32
		rp float64
		rn float64
		rz float64
	)
	l := len(arr)

	for _, x := range arr {
		if x < 0 {
			n++
		} else if x > 0 {
			p++
		} else {
			z++
		}
	}
	if p == 0 {
		rp = 0
	} else {
		rp = float64(p) / float64(l)
	}
	if n == 0 {
		rn = 0
	} else {
		rn = float64(n) / float64(l)
	}
	if z == 0 {
		rz = 0
	} else {
		rz = float64(z) / float64(l)
	}
	fmt.Println(fmt.Sprintf("%.5f", rp))
	fmt.Println(fmt.Sprintf("%.5f", rn))
	fmt.Println(fmt.Sprintf("%.5f", rz))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

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

	plusMinus(arr)
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
