#!/usr/local/bin/python3
import sys

FACTORS = 4
DIMENSION = 20
matrix = [[]]

def readMatrix(dimension):
    return [[int(n) for n in line.split()] for line in sys.stdin]

def left(y, x):
    boundx = FACTORS-1
    prod = 1
    if x>=boundx:
        for i in range(FACTORS):
            prod *= matrix[y][x-i]
    return prod

def right(y, x):
    boundx = DIMENSION-FACTORS
    prod = 1
    if x<=boundx:
        for i in range(FACTORS):
            prod *= matrix[y][x+i]
    return prod

def up(y, x):
    boundy = FACTORS-1
    prod = 1
    if y>=boundy:
        for i in range(FACTORS):
            prod *= matrix[y-i][x]
    return prod
    

def down(y, x):
    boundy = DIMENSION-FACTORS
    prod = 1
    if y<=boundy:
        for i in range(FACTORS):
            prod *= matrix[y+i][x]
    return prod

def upLeft(y, x):
    boundx = FACTORS-1
    boundy = FACTORS-1
    prod = 1
    if x>=boundx and y>=boundy:
        for i in range(FACTORS):
            prod *= matrix[y-i][x-i]
    return prod

def upRight(y, x):
    boundx = DIMENSION-FACTORS
    boundy = FACTORS-1
    prod = 1
    if x<=boundx and y>=boundy:
        for i in range(FACTORS):
            prod *= matrix[y-i][x+i]
    return prod

def downLeft(y, x):
    boundx = FACTORS-1
    boundy = DIMENSION-FACTORS
    prod = 1
    if x>=boundx and y<=boundy:
        for i in range(FACTORS):
            prod *= matrix[y+i][x-i]
    return prod

def downRight(y, x):
    boundx = DIMENSION-FACTORS
    boundy = DIMENSION-FACTORS
    prod = 1
    if x<=boundx and y<=boundy:
        for i in range(FACTORS):
            prod *= matrix[y+i][x+i]
    return prod


def prodMax(y, x):
    return max( left(y,x), \
                right(y,x), \
                up(y,x), \
                down(y,x), \
                upLeft(y,x), \
                upRight(y,x), \
                downLeft(y,x), \
                downRight(y,x))


if __name__ == "__main__":
    matrix = readMatrix(DIMENSION)
    maxValue = 0
    for i in range(DIMENSION):
        for j in range(DIMENSION):
            maxValue = max(prodMax(i,j), maxValue )
    print ("Max number is " + str(maxValue))
