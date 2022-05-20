#!/usr/bin/env python3

import argparse
from pprint import pprint


def main():
    args = parse_args()
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
