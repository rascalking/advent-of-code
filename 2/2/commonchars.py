#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import itertools


def common(a, b):
    return ''.join(a[i] for i in range(len(a))
                       if a[i] == b[i])


def distance(elems):
    a,b = elems
    dist = sum(0 if a[i] == b[i] else 1
                   for i in range(len(a)))
    return (dist, a, b)


def main():
    args = parse_args()
    boxids = map(str.strip, args.input.readlines())
    distances = sorted(map(distance, itertools.combinations(boxids, 2)))
    dist, a, b = distances[0]
    assert dist == 1
    print(common(a,b))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
