#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import copy
import re
from collections import defaultdict
from pprint import pprint


NUM_GENERATIONS = 20


class System(object):
    PLANT = '#'
    NO_PLANT = '.'

    def __init__(self, initial, rules):
        self.current = initial
        initial_len = len(initial)
        for i in range(1, NUM_GENERATIONS+1):
            self.current[-i] = self.NO_PLANT
            self.current[initial_len+i] = self.NO_PLANT
        self.rules = rules

    def __str__(self):
        return '<({}, {}): {}>'.format(
                   min(self.current.keys()),
                   max(self.current.keys()),
                   ''.join([i[1] for i in sorted(self.current.items())]))

    def sum_of_plants(self):
        return sum(k for k,v in self.current.items() if v == self.PLANT)

    def tick(self):
        next_ = copy.deepcopy(self.current)
        for number in list(self.current.keys()):
            key = ''.join(self.current[number+i] for i in range(-2,3))
            next_[number] = self.rules[key]
        self.current = next_
        print(self)


def main():
    args = parse_args()
    initial, rules = parse_input(args.input)
    system = System(initial, rules)
    print(system)
    for _ in range(NUM_GENERATIONS):
        system.tick()
    print('sum of plants:', system.sum_of_plants())
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


def parse_input(input_):
    initial_str = input_.readline().strip().split(': ')[-1]
    initial = defaultdict(lambda: System.NO_PLANT)
    initial.update((int(n), p) for n,p in enumerate(initial_str))

    rules = defaultdict(lambda: System.NO_PLANT)
    for line in input_:
        line = line.strip()
        if not line:
            continue
        pattern, result = line.split(' => ')
        rules[pattern] = result
    return initial, rules


if __name__ == '__main__':
    raise SystemExit(main())
