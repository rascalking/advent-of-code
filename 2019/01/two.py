#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint


def fuel_for_mass(mass):
    """
    Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

    For example:

        For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
        For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
        For a mass of 1969, the fuel required is 654.
        For a mass of 100756, the fuel required is 33583.
    """
    total_fuel = 0
    current_mass = mass
    while current_mass > 0:
        fuel = max(int(current_mass/3)-2,0)
        total_fuel += fuel
        current_mass = fuel
    return total_fuel


def main():
    args = parse_args()
    masses = list(map(int, args.input))
    total_fuel = sum(map(fuel_for_mass, masses))
    print(total_fuel)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
