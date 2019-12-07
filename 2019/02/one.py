#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint


def main():
    args = parse_args()
    ints = list(map(int, args.input.read().split(',')))
    ints[1] = 12
    ints[2] = 2
    for ptr in range(0, len(ints), 4):
        opcode = ints[ptr]
        if opcode == 1:
            src1, src2, dest = ints[ptr+1:ptr+4]
            ints[dest] = ints[src1] + ints[src2]
        elif opcode == 2:
            src1, src2, dest = ints[ptr+1:ptr+4]
            ints[dest] = ints[src1] * ints[src2]
        elif opcode == 99:
            print(ints[0])
            break
        else:
            raise SystemExit('Got opcode {} at index {}'.format(opcode, ptr))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
