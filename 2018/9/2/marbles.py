#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import re
from pprint import pprint


class Node(object):
    def __init__(self, marble, previous, next):
        self.marble = marble
        self.previous = previous
        self.next = next


class Circle(object):
    def __init__(self):
        node = Node(0, None, None)
        node.previous = node.next = node
        self.current = node

    def insert_marble(self, marble):
        new = Node(marble, self.current.next, self.current.next.next)
        new.previous.next = new
        new.next.previous = new
        self.current = new

    def remove_marble(self):
        to_remove = self.current
        for _ in range(7):
            to_remove = to_remove.previous
        to_remove.previous.next = to_remove.next
        to_remove.next.previous = to_remove.previous
        self.current = to_remove.next
        return to_remove.marble


def main():
    args = parse_args()
    match = re.search('(\d+) players; last marble is worth (\d+) points',
                      args.input.read())
    num_players, last_marble = map(int, match.groups())
    last_marble *= 100

    scores = [0] * num_players
    circle = Circle()
    current_player = 0
    for marble in range(1, last_marble+1):
        if marble % 23 == 0:
            scores[current_player] += marble
            scores[current_player] += circle.remove_marble()
        else:
            circle.insert_marble(marble)
        current_player = (current_player + 1) % num_players
    print(max(scores))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
