#!/usr/bin/env python3

import argparse
from collections import defaultdict
from pprint import pprint

START = 'start'
END = 'end'


def find_paths(edges: dict, node: str, path: tuple):
    paths = set()
    for n in edges[node]:
        n_path = path + (n,)
        if n == END:
            paths.add(n_path)
        elif n.islower() and n in path:
            continue
        else:
            paths.update(find_paths(edges, n, n_path))
    return paths


def main():
    args = parse_args()
    edges = defaultdict(list)
    for line in args.input:
        a, b = line.strip().split('-', 1)
        edges[a].append(b)
        edges[b].append(a)
    paths = find_paths(edges, START, (START,))
    pprint(paths)
    print(len(paths))
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
