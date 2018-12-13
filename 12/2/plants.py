#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint


class System(object):
    NUM_GENERATIONS = 50000000
    PLANT = '#'
    NO_PLANT = '.'

    def __init__(self, initial, rules):
        self.generation = 0
        self.padding = self.NUM_GENERATIONS + 2
        self.initial_length = len(initial)
        self.current = initial + (['.'] * (self.padding * 2))
        self.rules = rules

    def __str__(self):
        return '<Gen {}: ({} - {}): {}>'.format(
            self.generation,
            -self.padding,
            self.initial_length + self.padding,
            ''.join(self.current[-self.padding:])
                + ''.join(self.current[:-self.padding]),
        )

    def plant_pot_number_total(self):
        total = 0
        for i in range(-self.padding, self.initial_length + self.padding):
            if self.current[i] == self.PLANT:
                total += i
        return total

    def tick(self):
        self.generation += 1
        next_ = self.current.copy()
        for i in range(-self.padding + 2,
                       self.initial_length + self.padding - 2):
            pattern = ''.join(self.current[i+n] for n in range(-2,3))
            next_[i] = self.rules.get(pattern, '.')
        self.current = next_
        if self.generation % 1000000 == 0:
            print('Generation', self.generation)


def main():
    args = parse_args()
    system = System(*parse_input(args))
    for _ in range(System.NUM_GENERATIONS):
        system.tick()
    print('plant pot number total', system.plant_pot_number_total())
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


def parse_input(args):
    initial = list(args.input.readline().strip().split(': ')[-1])
    rules = {}
    for line in args.input:
        line = line.strip()
        if not line:
            continue
        pattern, result = line.split(' => ')
        rules[pattern] = result
    return initial, rules


if __name__ == '__main__':
    raise SystemExit(main())
