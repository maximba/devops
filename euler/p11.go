package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Dimension of the Matrix
const Dimension = 20

// Factors to use
const Factors = 4

type imatrix interface {
	left(int, int) uint64
	right(int, int) uint64
	up(int, int) uint64
	down(int, int) uint64
	upLeft(int, int) uint64
	upRight(int, int) uint64
	downLeft(int, int) uint64
	downRight(int, int) uint64
	prodMax(int, int) uint64
	readMatrix()
}

type matrix struct {
	grid               [][]uint64
	dimension, factors int
}

func max(nums ...uint64) uint64 {
	var maxValue uint64
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

func (m matrix) left(y int, x int) uint64 {
	var prod uint64
	boundx := m.factors - 1
	if x >= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y][x-i]
		}
	}
	return prod
}

func (m matrix) right(y int, x int) uint64 {
	var prod uint64
	boundx := m.dimension - m.factors
	if x <= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y][x+i]
		}
	}
	return prod
}

func (m matrix) up(y int, x int) uint64 {
	var prod uint64
	boundy := m.factors - 1
	if y >= boundy {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y-i][x]
		}
	}
	return prod
}

func (m matrix) down(y int, x int) uint64 {
	var prod uint64
	boundy := m.dimension - m.factors
	if y <= boundy {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y+i][x]
		}
	}
	return prod
}

func (m matrix) upLeft(y int, x int) uint64 {
	var prod uint64
	boundx := m.factors - 1
	boundy := m.factors - 1
	if y >= boundy && x >= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y-i][x-i]
		}
	}
	return prod
}

func (m matrix) upRight(y int, x int) uint64 {
	var prod uint64
	boundx := m.dimension - m.factors
	boundy := m.factors - 1
	if y >= boundy && x <= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y-i][x+i]
		}
	}
	return prod
}

func (m matrix) downLeft(y int, x int) uint64 {
	var prod uint64
	boundx := m.factors - 1
	boundy := m.dimension - m.factors
	if y <= boundy && x >= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y+i][x-i]
		}
	}
	return prod
}

func (m matrix) downRight(y int, x int) uint64 {
	var prod uint64
	boundx := m.dimension - m.factors
	boundy := m.dimension - m.factors
	if y <= boundy && x <= boundx {
		prod = 1
		for i := 0; i < m.factors; i++ {
			prod *= m.grid[y+i][x+i]
		}
	}
	return prod
}

func (m matrix) prodMax(y int, x int) uint64 {
	return max(m.left(y, x),
		m.right(y, x),
		m.up(y, x),
		m.down(y, x),
		m.upLeft(y, x),
		m.upRight(y, x),
		m.downLeft(y, x),
		m.downRight(y, x))
}

func (m matrix) readMatrix() matrix {
	i := 0
	j := 0

	fs := bufio.NewScanner(os.Stdin)
	fs.Split(bufio.ScanWords)

	m.grid[0] = make([]uint64, m.dimension)
	for fs.Scan() {
		n, _ := strconv.Atoi(strings.TrimLeft(fs.Text(), "0"))
		m.grid[i][j] = uint64(n)
		j++
		if j == m.dimension {
			i++
			j = 0
			if i < m.dimension {
				m.grid[i] = make([]uint64, m.dimension)
			}
		}
	}
	return m
}

func (m matrix) printGrid() {
	for i := 0; i < m.dimension; i++ {
		for j := 0; j < m.dimension; j++ {
			fmt.Printf("%02d ", m.grid[i][j])
		}
		fmt.Println("")
	}

}

func main() {
	var maxValue uint64
	m := matrix{
		grid:      make([][]uint64, Dimension),
		dimension: Dimension,
		factors:   Factors,
	}
	m.readMatrix()
	m.printGrid()

	for i := 0; i < m.dimension; i++ {
		for j := 0; j < m.dimension; j++ {
			value := m.prodMax(i, j)
			maxValue = max(value, maxValue)
		}
	}
	fmt.Printf("Max number is %d\n", maxValue)
}
