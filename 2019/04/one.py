#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint


def check(num):
    digits = list(map(int, list(str(num))))

    if digits != sorted(digits):
        return False

    double = False
    for i in range(len(digits)-1):
        if digits[i] == digits[i+1]:
            double = True
    return double


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
