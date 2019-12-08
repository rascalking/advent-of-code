#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from collections import Counter
from pprint import pprint


def check(num):
    digits = list(map(int, list(str(num))))

    if digits != sorted(digits):
        return False

    counter = Counter(digits)
    for digit, instances in counter.items():
        if instances == 2:
            return True
    return False


def main():
    args = parse_args()
    low, high = map(int, args.input.read().strip().split('-'))

    matches = list(filter(check, range(low, high+1)))
    print(len(matches))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
