#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from collections import defaultdict
from operator import add, sub
from pprint import pprint


def main():
    args = parse_args()
    operations = {
        'U': lambda x: list(map(add, x, [0,1])),
        'D': lambda x: list(map(sub, x, [0,1])),
        'L': lambda x: list(map(sub, x, [1,0])),
        'R': lambda x: list(map(add, x, [1,0])),
    }

    steps = defaultdict(lambda: list((list(), list())))
    for wire, path in enumerate(args.input):
        moves = path.split(',')
        position = [0,0]
        step = 0
        for move in moves:
            operation = operations[move[0]]
            length = int(move[1:])
            for i in range(length):
                step += 1
                position = operation(position)
                if position != [0,0]:
                    steps[tuple(position)][wire].append(step)
    intersections = [x for x in steps if all(steps[x])]
    intersections.sort(key=lambda x: steps[x][0][0] + steps[x][1][0])
    print(intersections[0], steps[intersections[0]][0][0] + steps[intersections[0]][1][0])
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
