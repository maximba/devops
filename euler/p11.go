package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const DIM = 20
const CHUNK = 4

type imatrix interface {
	left(int, int) uint64
	right(int, int) uint64
	up(int, int) uint64
	down(int, int) uint64
	up_left(int, int) uint64
	up_right(int, int) uint64
	down_left(int, int) uint64
	down_right(int, int) uint64
	prod_max(int, int) uint64
	read_matrix()
}

type matrix struct {
	grid [][]uint64
	dim, chunk int
}

func max( nums ... uint64) uint64 {
	var max_value uint64 = 0
	for _, num := range nums {
		if num > max_value {
//			fmt.Printf("%d\n", num)
			max_value = num
		}
	}	
	return max_value
}

func (m matrix) left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.chunk -1
	if x >= boundx {
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx-i][y]
		}
	}
	return prod
}

func (m matrix ) right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.dim - m.chunk
	if x <= boundx {
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx+i][y]
		}
	}
	return prod
}

func (m matrix) up(x int, y int) uint64 {
	var prod uint64 = 0
	boundy := m.chunk -1
	if y >= boundy {
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[x][boundy-i]
		}
	}
	return prod
}

func (m matrix) down(x int, y int) uint64 {
	var prod uint64 = 0
	boundy := m.dim - m.chunk
	if y <= boundy {
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod = prod*m.grid[x][boundy+i]
		}
	}
	return prod
}

func (m matrix) up_left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.chunk - 1
	boundy := m.chunk - 1
	if y >= boundy && x>=boundx{
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx-i][boundy-i]
		}
	}
	return prod
}

func (m matrix) up_right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.dim - m.chunk
	boundy := m.chunk - 1
	if y >= boundy && x<=boundx{
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx+i][boundy-i]
		}
	}
	return prod
}

func (m matrix) down_left(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.chunk -1
	boundy := m.dim - m.chunk
	if y <= boundy && x>=boundx{
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx-i][boundy+i]
		}
	}
	return prod
}

func (m matrix) down_right(x int, y int) uint64 {
	var prod uint64 = 0
	boundx := m.dim - m.chunk
	boundy := m.dim - m.chunk
	if y <= boundy && x<=boundx{
		prod=1
		for i:=0; i<m.chunk; i++ {
			prod *= m.grid[boundx+i][boundy+i]
		}
	}
	return prod
}

func (m matrix) prod_max(x int, y int) uint64 {
	return max( m.left(x,y), 
				m.right(x,y), 
				m.up(x,y), 
				m.down(x,y), 
				m.up_left(x,y), 
				m.up_right(x,y), 
				m.down_left(x,y), 
				m.down_right(x,y))
}

func (m matrix) read_matrix() (matrix) {
	x := 0
	y := 0

	fs := bufio.NewScanner(os.Stdin)
	fs.Split(bufio.ScanWords)

	m.grid[0] = make([]uint64, m.dim)
	for fs.Scan() {
		n, _ := strconv.Atoi(strings.TrimLeft(fs.Text(), "0"))
		m.grid[x][y] = uint64(n)
		y = y+1
		if y == m.dim {
			x = x+1
			y = 0
			if (x < m.dim) {
				m.grid[x] = make([]uint64, m.dim)
			}
		}
	}
	return m
}

func print_grid(m matrix) {
	for i:=0; i<m.dim; i++ {
		for j:=0; j<m.dim; j++ {
			fmt.Printf("%02d ", m.grid[i][j])
		}
		fmt.Println("")
	}

}

func main () {
	var max_value uint64 = 0
	m := matrix{
		grid: make([][]uint64, DIM), 
		dim: DIM, 
		chunk: CHUNK,
	}	
	m.read_matrix()
	print_grid(m)
	for i:=0; i<m.dim; i++ {
		for j:=0; j<m.dim; j++ {
			value := m.prod_max(i, j)
			fmt.Printf("Current value is: %d. Old Max Value is: %d\n", value, max_value)
			max_value = max(value, max_value)			
		}
	}
	fmt.Printf("Max number is %d", max_value)
}