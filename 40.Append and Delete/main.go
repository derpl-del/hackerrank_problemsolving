package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
if s == t {
        if k >= int32(len(s)*2) || k%2 == 0 {
            return "Yes"
        } else {
            return "No"
        }
    } else {
        min := 0
        v1 := []byte(s)
        v2 := []byte(t)
        for i := 0; i < int(math.Min(float64(len(s)), float64(len(t)))); i++ {
            if v1[i] != v2[i] {
                break
            }
            min++
        }

        d1 := len(s) - min
        d2 := len(t) - min
        if  d1+d2 < int(k) && (d1+d2-int(k))%2 == 0 ||  len(s)+len(t) <= int(k)  || d1+d2 == int(k)        {
            return "Yes"
        }
        return "No"
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

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
