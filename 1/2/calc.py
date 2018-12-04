#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse


def main():
    args = parse_args()
    inputs = list(map(int, args.input.readlines()))

    freq = 0
    seen = set()
    for _ in range(100000):
        for change in inputs:
            freq += change
            if freq in seen:
                print('FOUND IT', freq)
                return 0
            seen.add(freq)
    print('DIDN\'T FIND IT, {} frequencies seen'.format(len(seen)))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
