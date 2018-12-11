#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import copy
import re
from collections import namedtuple
from operator import add
from pprint import pprint


Vector = namedtuple('Vector', ['x', 'y'])
Bounds = namedtuple('Bounds', ['x_min', 'x_max', 'y_min', 'y_max'])


class Star(object):
    def __init__(self, position, velocity):
        self.position = position
        self.velocity = velocity

    def __repr__(self):
        return 'position=<{}, {}> velocity=<{}, {}>'.format(
                   self.position.x, self.position.y,
                   self.velocity.x, self.velocity.y)

    def tick(self, num_secs):
        self.position = Vector(
            self.position.x + (num_secs * self.velocity.x),
            self.position.y + (num_secs * self.velocity.y)
        )


class Sky(object):
    def __init__(self, stars):
        self.stars = stars
        self.time = 0

    def bounds(self):
        xs = [s.position.x for s in self.stars]
        x_min, x_max = min(xs), max(xs)
        ys = [s.position.y for s in self.stars]
        y_min, y_max = min(ys), max(ys)
        return Bounds(x_min, x_max, y_min, y_max)

    def print(self):
        print('time', self.time)
        bounds = self.bounds()
        frame = [['.'] * (bounds.y_max - bounds.y_min + 1) for _ in range(bounds.x_max - bounds.x_min + 1)]
        for star in self.stars:
            frame[star.position.x-bounds.x_min][star.position.y-bounds.y_min] = '#'
        for y in range(0, bounds.y_max - bounds.y_min + 1):
            for x in range(0, bounds.x_max - bounds.x_min + 1):
                print(frame[x][y], end='')
            print('')
        print('\n')

    def tick(self, num_secs=1):
        for star in self.stars:
            star.tick(num_secs)
        self.time += num_secs


def main():
    args = parse_args()
    stars = []
    for line in args.input:
        match = re.search(
                    'position=<\s*(-?\d+), \s*(-?\d+)> velocity=<\s*(-?\d+), \s*(-?\d+)>',
                    line)
        position = Vector(int(match.group(1)), int(match.group(2)))
        velocity = Vector(int(match.group(3)), int(match.group(4)))
        stars.append(Star(position, velocity))
    sky = Sky(copy.deepcopy(stars))

    # find the smallest bounds that will cover the stars
    boundses = []
    for tick in range(100000):
        sky.tick()
        bounds = sky.bounds()
        size = (bounds.x_max-bounds.x_min) * (bounds.y_max-bounds.y_min)
        boundses.append((size, tick, bounds))
    boundses.sort()
    print(boundses[0])
    smallest = boundses[0][1]
    del(boundses)

    # reset to the time with the smalled bounds
    sky = Sky(copy.deepcopy(stars))
    sky.tick(smallest-5)
    for _ in range(10):
        sky.print()
        sky.tick()
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
