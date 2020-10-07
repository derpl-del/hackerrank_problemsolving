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

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var result int32
	up := n - r_q
	down := r_q - 1
	right := n - c_q
	left := c_q - 1
	d_ul := int32(math.Min(float64(up), float64(left)))
	d_ur := int32(math.Min(float64(up), float64(right)))
	d_dl := int32(math.Min(float64(down), float64(left)))
	d_dr := int32(math.Min(float64(down), float64(right)))
	//fmt.Printf("u:%v d:%v r:%v l:%v", up, down, right, left)
	//fmt.Println()
	//fmt.Println("+++")
	r2_r := int32(999999)
	r2_l := int32(999999)
	r2_u := int32(999999)
	r2_d := int32(999999)
	r2_ur := int32(999999)
	r2_ul := int32(999999)
	r2_dl := int32(999999)
	r2_dr := int32(999999)

	for _, x := range obstacles {
		//fmt.Println(x)
		o_u := n - x[0]
		o_d := x[0] - 1
		o_r := n - x[1]
		o_l := x[1] - 1
		//fmt.Printf("u:%v d:%v r:%v l:%v", o_u, o_d, o_r, o_l)
		//fmt.Println()
		//right
		if o_u == up && o_r < right {
			//fmt.Println("right")
			r_right := right - o_r - 1
			r2_r = int32(math.Min(float64(r2_r), float64(r_right)))
		}

		if o_u == up && o_l < left {
			//fmt.Println("left")
			r_left := left - o_l - 1
			r2_l = int32(math.Min(float64(r2_l), float64(r_left)))
		}

		if o_u > up && o_l == left && o_r == right {
			//fmt.Println("down")
			r_down := down - o_d - 1
			r2_d = int32(math.Min(float64(r2_d), float64(r_down)))
		}

		if o_u < up && o_l == left && o_r == right {
			//fmt.Println("up")
			r_up := up - o_u - 1
			r2_u = int32(math.Min(float64(r2_u), float64(r_up)))
		}
		//fmt.Println("===")

		if math.Abs(float64(o_u-up)) == math.Abs(float64(o_l-left)) && o_u > up && o_l < left {
			//fmt.Println("d_dl")
			d2_dl := int32(math.Min(float64(o_d), float64(o_l)))
			d2_dl = int32(math.Abs(float64((d2_dl + 1) - d_dl)))
			r2_dl = int32(math.Min(float64(r2_dl), float64(d2_dl)))
		}
		if math.Abs(float64(o_u-up)) == math.Abs(float64(o_r-right)) && o_u > up && o_r < right {
			//fmt.Println("d_dr")
			d2_dr := int32(math.Min(float64(o_d), float64(o_r)))
			d2_dr = int32(math.Abs(float64((d2_dr + 1) - d_dr)))
			r2_dr = int32(math.Min(float64(r2_dr), float64(d2_dr)))
			//fmt.Println(r2_dr)
		}
		if math.Abs(float64(o_d-down)) == math.Abs(float64(o_l-left)) && o_u < up && o_l < left {
			//fmt.Println("d_ul")
			d2_ul := int32(math.Min(float64(o_u), float64(o_l)))
			d2_ul = int32(math.Abs(float64((d2_ul + 1) - d_ul)))
			r2_ul = int32(math.Min(float64(r2_ul), float64(d2_ul)))
			//fmt.Println(r2_ul)

		}
		if math.Abs(float64(o_d-down)) == math.Abs(float64(o_l-left)) && o_u < up && o_l > left {
			//fmt.Println("d_ur")
			d2_ur := int32(math.Min(float64(o_u), float64(o_r)))
			d2_ur = int32(math.Abs(float64((d2_ur + 1) - d_ur)))
			r2_ur = int32(math.Min(float64(r2_ur), float64(d2_ur)))
			//fmt.Println(r2_ur)
		}
		//fmt.Println("===")
	}
	//fmt.Println("right")
	if r2_r == 999999 {
		result += right
		//fmt.Println(right)
	} else {
		result += r2_r
		//fmt.Println(r2_r)
	}
	//fmt.Println("left")
	if r2_l == 999999 {
		result += left
		//fmt.Println(left)
	} else {
		result += r2_l
		//fmt.Println(r2_l)
	}
	//fmt.Println("up")
	if r2_u == 999999 {
		result += up
		//fmt.Println(up)
	} else {
		result += r2_u
		//fmt.Println(r2_u)
	}
	//fmt.Println("down")
	if r2_d == 999999 {
		result += down
		//fmt.Println(down)
	} else {
		result += r2_d
		//fmt.Println(r2_d)
	}
	//fmt.Println("ur")
	if r2_ur == 999999 {
		result += d_ur
		//fmt.Println(d_ur)
	} else {
		result += r2_ur
		//fmt.Println(r2_ur)
	}
	//fmt.Println("ul")
	if r2_ul == 999999 {
		result += d_ul
		//fmt.Println(d_ul)
	} else {
		result += r2_ul
		//fmt.Println(r2_ul)
	}
	//fmt.Println("dl")
	if r2_dl == 999999 {
		result += d_dl
		//fmt.Println(d_dl)
	} else {
		result += r2_dl
		//fmt.Println(r2_dl)
	}
	//fmt.Println("dr")
	if r2_dr == 999999 {
		result += d_dr
		//fmt.Println(d_dr)
	} else {
		result += r2_dr
		//fmt.Println(r2_dr)
	}
	return result
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

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
