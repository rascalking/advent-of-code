#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import string


def react(polymer):
    length = 0
    while len(polymer) != length:
        length = len(polymer)
        for i in range(len(polymer)-1):
            if polymer[i].swapcase() == polymer[i+1]:
                polymer = polymer.replace(polymer[i:i+2], '')
                break
    return polymer


def main():
    args = parse_args()
    polymer = args.input.read().strip()

    shortest_length = len(polymer)
    for char in string.ascii_lowercase:
        candidate = polymer.replace(char, '').replace(char.upper(), '')
        length = len(react(candidate))
        shortest_length = min(length, shortest_length)

    print(shortest_length)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
