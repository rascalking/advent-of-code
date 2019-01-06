#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import re

CLAIM_RE = re.compile(
    '#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)')


def main():
    args = parse_args()

    # parse it first so we know how big the unconflicted set needs to be
    claims = []
    for line in args.input:
        match = CLAIM_RE.match(line)
        assert match
        claim = {k: int(v) for k,v in match.groupdict().items()}
        claims.append(claim)

    # 1k x 1k isn't big enough to not do this the easy way
    unconflicted = set(range(1,claims[-1]['id']+1))
    cloth = [[set() for _ in range(1000)] for _ in range(1000)]
    for claim in claims:
        for x in range(claim['left'], claim['left']+claim['width']):
            for y in range(claim['top'], claim['top']+claim['height']):
                cloth[x][y].add(claim['id'])
                if len(cloth[x][y]) > 1:
                    unconflicted.difference_update(cloth[x][y])
    print(unconflicted)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
