#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import sys
from collections import defaultdict, namedtuple
from pprint import pprint


Bounds = namedtuple('Bounds', ['min_x', 'min_y', 'max_x', 'max_y'])
Point = namedtuple('Point', ['x', 'y'])
MARGIN = 10


def calculate_areas(grid):
    areas = defaultdict(int)
    for point, closest in grid.items():
        areas[closest] += 1

    # assume anything with an area that reaches the edge of our margin
    # is an infinite area, so exclude it from the set of areas we return
    bounds = find_bounds(grid.keys())
    for point in grid:
        if not within_bounds(bounds, point):
            try:
                del areas[grid[point]]
            except KeyError:
                pass

    return areas


def calculate_closest(bounds, points):
    grid = {}
    for x in range(bounds.min_x-MARGIN, bounds.max_x+MARGIN+1):
        for y in range(bounds.min_y-MARGIN, bounds.max_y+MARGIN+1):
            current = Point(x,y)
            distances = defaultdict(set)
            for point in points:
                distances[calculate_distance(current, point)].add(point)
            min_distance = min(distances.keys())
            if len(distances[min_distance]) == 1:
                grid[current] = distances[min_distance].pop()
            else:
                grid[current] = None
    return grid


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
    grid = calculate_closest(bounds, points)
    areas = calculate_areas(grid)
    largest_area = max(areas.values())
    print(largest_area)
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
