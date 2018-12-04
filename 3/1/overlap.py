#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import re

CLAIM_RE = re.compile(
    '#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)')


def main():
    args = parse_args()

    # 1k x 1k isn't big enough to not do this the easy way
    cloth = [[set() for _ in range(1000)] for _ in range(1000)]
    for line in args.input:
        match = CLAIM_RE.match(line)
        assert match
        claim = {k: int(v) for k,v in match.groupdict().items()}
        for x in range(claim['left'], claim['left']+claim['width']):
            for y in range(claim['top'], claim['top']+claim['height']):
                cloth[x][y].add(claim['id'])

    total = sum(sum(1 if len(cloth[x][y]) > 1 else 0 for y in range(1000)) for x in range(1000))
    print(total)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
