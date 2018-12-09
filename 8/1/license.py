#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint

import networkx as nx


class Node(object):
    def __init__(self, offset):
        self.offset = offset
        self.metadata = []
        self.length = 0


def add_node(offset, nums, graph, parent):
    node = Node(offset)
    if parent:
        graph.add_edge(parent, node)
    num_children = nums[offset]
    num_metadata = nums[offset+1]
    if num_children == 0:
        node.metadata = nums[offset+2:offset+2+num_metadata]
        node.length = 2 + num_metadata
    else:
        child_offset = offset+2
        for _ in range(num_children):
            child = add_node(child_offset, nums, graph, node)
            child_offset += child.length
        node.metadata = nums[child_offset:child_offset+num_metadata]
        node.length = child_offset + num_metadata - offset
    return node


def main():
    args = parse_args()
    nums = list(map(int, args.input.read().strip().split()))
    graph = nx.DiGraph()
    root = add_node(0, nums, graph, None)
    total = sum(sum(n.metadata) for n in graph.nodes)
    print(total)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
