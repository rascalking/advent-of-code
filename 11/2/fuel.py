#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import math
import os
from collections import namedtuple
from multiprocessing import Pool
from pprint import pprint


def calculate_levels(serial):
    levels = [[0] * 301 for _ in range(301)]
    for x in range(1, 301):
        for y in range(1, 301):
            levels[x][y] = power_level(x, y, serial)
    return levels


def best_level_for_coordinate(args):
    levels, x, y = args
    sizes = []
    previous = 0
    for size in range(1, 301-(max(x,y)-1)):
        power = (previous +
                 sum(levels[x][y1] for y1 in range(y,y+size)) +
                 sum(levels[x1][y] for x1 in range(x,x+size)))
        sizes.append((power, size))
        previous = power
    sizes.sort()
    power, size = sizes[-1]
    return (power, (x, y, size))


def calculate_squares(levels):
    coords = [(x,y) for x in range(1,301) for y in range(1,301)]
    with Pool() as pool:
        squares = pool.map(best_level_for_coordinate,
                           [(levels, x, y) for x in range(1,301)
                                           for y in range(1,301)])
    squares.sort()
    return squares


def main():
    args = parse_args()
    serial = int(args.input.read().strip())
    levels = calculate_levels(serial)
    squares = calculate_squares(levels)
    print(squares[-1])
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


def power_level(x, y, serial):
    rack_id = x + 10
    level = rack_id * y
    level += serial
    level *= rack_id
    level = math.floor((level % 1000) / 100)
    level -= 5
    return level


if __name__ == '__main__':
    raise SystemExit(main())
