#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import math
from collections import namedtuple
from multiprocessing import Pool
from pprint import pprint


def best_size(args):
    levels, x, y = args
    sums = [(levels[x][y], (x,y,1))]
    previous = levels[x][y]
    for size in range(2,302-max(x,y)):
        power = (previous +
                 sum(levels[x+size-1][y1] for y1 in range(y,y+size)) +
                 sum(levels[x1][y+size-1] for x1 in range(x,x+size-1)))
        sums.append((power, (x,y,size)))
        previous = power
    sums.sort()
    return sums[-1]


def calculate_levels(serial):
    levels = [[0] * 301 for _ in range(301)]
    for x in range(1, 301):
        for y in range(1, 301):
            levels[x][y] = power_level(x, y, serial)
    return levels


def calculate_squares(levels):
    with Pool() as pool:
        squares = pool.map(
            best_size,
            ((levels, x, y) for x in range (1,301)
                            for y in range(1,301)))
    squares.sort()
    return squares

def main():
    args = parse_args()
    serial = args.input
    levels = calculate_levels(serial)
    squares = calculate_squares(levels)
    print(squares[-1])
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=int)
    return parser.parse_args()


def power_level(x, y, serial):
    rack_id = x + 10
    level = rack_id * y
    level += serial
    level *= rack_id
    level = math.floor((level % 1000) / 100)
    level -= 5
    return level


def print_square(levels, x, y, size):
    for y1 in range(y, y+size):
        print(' '.join(map(str, (levels[x1][y1] for x1 in range(x, x+size)))))


def square_total(levels, x, y, size):
    return sum(levels[x1][y1] for x1 in range(x,x+size) for y1 in range(y,y+size))


if __name__ == '__main__':
    raise SystemExit(main())
