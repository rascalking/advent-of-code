#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import itertools
import re
from pprint import pprint

import networkx as nx


BASE_JOB_SECS = 60
NUM_WORKERS = 5


class Worker(object):
    def __init__(self, name, graph):
        self.name = name
        self.graph = graph
        self.node = self.done_at = None

    def is_available(self):
        return self.node is None and self.done_at is None

    def start(self, node, tick):
        assert self.is_available()
        assert self.graph.nodes[node]['available'] == True
        assert self.graph.nodes[node]['in_progress'] == False
        assert self.graph.nodes[node]['done'] == False
        self.node = node
        self.done_at = tick + duration(node)
        self.graph.nodes[node]['in_progress'] = True
        
    def tick(self, tick):
        if self.done_at and self.done_at <= tick:
            node = self.node
            self.graph.nodes[node]['in_progress'] = False
            self.graph.nodes[node]['done'] = True
            self.node, self.done_at = None, None
            return node
        return None


def build_graph(args):
    graph = nx.DiGraph()
    for line in args.input:
        match = re.search(
                    'Step (\w+) must be finished before step (\w+) can begin.\n',
                    line)
        if match:
            graph.add_edge(*match.groups())

    for node in graph:
        graph.nodes[node]['available'] = (graph.in_degree(node) == 0)
        graph.nodes[node]['done'] = False
        graph.nodes[node]['in_progress'] = False
    return graph


def duration(node):
    return BASE_JOB_SECS + ord(node) - 64
    # ord('A') is 65, and A should take 61 seconds
    return ord(node)-4


def mark_available(graph):
    import ipdb; ipdb.set_trace
    for node in graph:
        if graph.in_degree(node) == 0:
            graph.nodes[node]['available'] = True
        elif all(graph.nodes[p]['done'] for p in graph.predecessors(node)):
            graph.nodes[node]['available'] = True
        else:
            graph.nodes[node]['available'] = False


def main():
    args = parse_args()
    graph = build_graph(args)
    workers = [Worker(i, graph) for i in range(NUM_WORKERS)]

    """
    for tick in itertools.count():
        get any nodes that just completed, add them to done list
        if set(done) == set(graph.nodes):
            break
        check to see if any successors to those just done nodes are now available
            if so, add them to the available list
        for each available worker:
            pop from the available list, break out of the loop if there are none
            assign the node to the worker
    """
    for tick in itertools.count():
        # get any nodes that just completed, add them to done list
        for worker in workers:
            node = worker.tick(tick)

        # break out if we're all done
        if all(graph.nodes[n]['done'] for n in graph.nodes):
            break

        # check to see if any new nodes should be flagged as available
        mark_available(graph)

        # pass out work to available workers
        for worker in workers:
            if worker.is_available():
                todo = sorted(
                    n for n in graph.nodes if (graph.nodes[n]['available'] and
                                               (not graph.nodes[n]['done']) and
                                               (not graph.nodes[n]['in_progress'])))
                if todo:
                    worker.start(todo[0], tick)
    print(tick)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
