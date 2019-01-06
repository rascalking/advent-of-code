#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
from collections import defaultdict


def get_counts(boxid):
    counts = defaultdict(int)
    for char in boxid:
        counts[char] += 1

    two, three = 0, 0
    for char, instances in counts.items():
        if instances == 2:
            two = 1
        elif instances == 3:
            three = 1
    return two, three


def main():
    args = parse_args()

    # [(0, 0), ...]
    twos, threes = 0, 0
    for two, three in map(get_counts, args.input):
        twos += two
        threes += three

    print('checksum is', twos * threes)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
