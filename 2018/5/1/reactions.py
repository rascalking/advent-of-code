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
    print(len(react(polymer)))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
