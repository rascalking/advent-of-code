#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import math
from collections import namedtuple
from pprint import pprint


def calculate_levels(serial):
    levels = [[0] * 301 for _ in range(301)]
    for x in range(1, 301):
        for y in range(1, 301):
            levels[x][y] = power_level(x, y, serial)
    return levels


def calculate_squares(levels):
    squares = []
    for x in range(1, 299):
        for y in range(1, 299):
            power = sum(levels[x1][y1]
                            for x1 in range(x, x+3)
                            for y1 in range(y, y+3))
            squares.append((power, (x,y)))
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


if __name__ == '__main__':
    raise SystemExit(main())
