package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const dim = 20
const chunk = 4

var matrix [][]uint64

func max( nums ... uint64) uint64 {
	var max_value uint64 = 0
	for _, num := range nums {
		if num > max_value {
			fmt.Printf("%d\n", num)
			max_value = num
		}
	}	
	return max_value
}

func left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := chunk -1
	if x >= boundx {
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx-i][y]
		}
	}
	return prod
}

func right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := dim - chunk
	if x <= boundx {
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx+i][y]
		}
	}
	return prod
}

func up(x int, y int) uint64 {
	var prod uint64 = 0
	boundy := chunk -1
	if y >= boundy {
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[x][boundy-i]
		}
	}
	return prod
}

func down(x int, y int) uint64 {
	var prod uint64 = 0
	boundy := dim - chunk
	if y <= boundy {
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[x][boundy+i]
		}
	}
	return prod
}

func up_left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := chunk - 1
	boundy := chunk - 1
	if y >= boundy && x>=boundx{
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx-i][boundy-i]
		}
	}
	return prod
}

func up_right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := dim - chunk
	boundy := chunk - 1
	if y >= boundy && x<=boundx{
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx+i][boundy-i]
		}
	}
	return prod
}

func down_left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := chunk -1
	boundy := dim - chunk
	if y <= boundy && x>=boundx{
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx-i][boundy+i]
		}
	}
	return prod
}

func down_right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := dim - chunk
	boundy := dim - chunk
	if y <= boundy && x<=boundx{
		prod=1
		for i:=0; i<chunk; i++ {
			prod = prod*matrix[boundx+i][boundy+i]
		}
	}
	return prod}

func prod_max(x int, y int) uint64 {
	return max(left(x,y), right(x,y), up(x,y), down(x,y), up_left(x,y), up_right(x,y), down_left(x,y), down_right(x,y))
}

func read_matrix(size int) {
	x := 0
	y := 0
	fs := bufio.NewScanner(os.Stdin)
	fs.Split(bufio.ScanWords)
	matrix = make([][]uint64, size)
	matrix[0] = make([]uint64, size)
	for fs.Scan() {
		snumber := strings.TrimLeft(fs.Text(), "0")
		fmt.Print(snumber+" ")
		n, _ := strconv.Atoi(fs.Text())
		matrix[x][y] = uint64(n)
		y = y+1
		if y == size {
			x = x+1
			y = 0
			if (x < size) {
				matrix[x] = make([]uint64, size)
			}
			fmt.Println("")
		}
	}
}

func main () {
	read_matrix(dim)

    var max_value uint64 = 0
	for i:=0; i<dim; i++ {
		for j:=0; j<dim; j++ {
			max_value = max(prod_max(i, j), max_value)
		}
	}
	fmt.Printf("Max number is %d", max_value)
}