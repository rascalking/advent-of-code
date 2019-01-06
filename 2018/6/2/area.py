#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import sys
from collections import defaultdict, namedtuple
from pprint import pprint


Bounds = namedtuple('Bounds', ['min_x', 'min_y', 'max_x', 'max_y'])
Point = namedtuple('Point', ['x', 'y'])
MAX_SUMMED_DISTANCE = 10000


def calculate_distance(a, b):
    return abs(a.x-b.x) + abs(a.y-b.y)


def find_bounds(points):
    xs = [p.x for p in points]
    ys = [p.y for p in points]
    return Bounds(min(xs), min(ys), max(xs), max(ys))


def main():
    args = parse_args()
    points = [Point(*tuple(map(int, line.strip().split(', '))))
                  for line in args.input]
    bounds = find_bounds(points)

    area = 0
    for x in range(bounds.min_x, bounds.max_x+1):
        for y in range(bounds.min_y, bounds.max_y+1):
            location = Point(x,y)
            distance = sum(calculate_distance(location, p) for p in points)
            if distance < MAX_SUMMED_DISTANCE:
                area += 1
    print(area)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


def within_bounds(bounds, point):
    return (point.x > bounds.min_x and
            point.x < bounds.max_x and
            point.y > bounds.min_y and
            point.y < bounds.max_y)


if __name__ == '__main__':
    raise SystemExit(main())
