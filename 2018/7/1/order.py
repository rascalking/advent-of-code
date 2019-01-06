#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import re
from pprint import pprint

import networkx as nx


def main():
    args = parse_args()

    graph = nx.DiGraph()
    for line in args.input:
        match = re.search(
                    'Step (\w+) must be finished before step (\w+) can begin.\n',
                    line)
        if match:
            graph.add_edge(*match.groups())

    # prime the available list with any roots of the graph
    available = sorted(n for n in graph.nodes if graph.in_degree(n) == 0)
    ordered = []
    while available:
        node = available.pop(0)
        ordered.append(node)
        for successor in graph.successors(node):
            if successor in ordered or successor in available:
                continue
            if set(graph.predecessors(successor)).issubset(ordered):
                available.append(successor)
        available.sort()
    print(''.join(ordered))

    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
