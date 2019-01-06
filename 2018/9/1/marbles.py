#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import re
from pprint import pprint


def main():
    args = parse_args()
    match = re.search('(\d+) players; last marble is worth (\d+) points',
                      args.input.read())
    num_players, last_marble = map(int, match.groups())

    scores = [0] * num_players
    circle = [0]
    current_index, current_player = 0, 0
    for marble in range(1, last_marble+1):
        if marble % 23 == 0:
            scores[current_player] += marble
            next_index = (current_index - 7) % len(circle) 
            scores[current_player] += circle.pop(next_index)
        else:
            next_index = (current_index + 2) % len(circle)
            # let's not shift the whole damn array if we don't have to
            if next_index == 0:
                circle.append(marble)
                next_index = len(circle)-1
            else:
                circle.insert(next_index, marble)
        current_index = next_index
        current_player = (current_player + 1) % num_players
    print(max(scores))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
