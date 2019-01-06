#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse


def main():
    args = parse_args()
    print(sum(map(int, args.input.readlines())))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
