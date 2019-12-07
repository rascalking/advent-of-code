#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
from pprint import pprint

def run_program(memory, noun, verb):
    memory[1:3] = [noun, verb]
    pointer = 0
    while pointer < len(memory):
        opcode = memory[pointer]
        if opcode == 1:
            src1, src2, dest = memory[pointer+1:pointer+4]
            memory[dest] = memory[src1] + memory[src2]
        elif opcode == 2:
            src1, src2, dest = memory[pointer+1:pointer+4]
            memory[dest] = memory[src1] * memory[src2]
        elif opcode == 99:
            return memory[0]
        else:
            raise RuntimeError('Got unknown opcode {} at address {}'.format(
                                   opcode, pointer))
        pointer += 4
    raise RuntimeError('Instruction pointer ran off the end of memory')


def main():
    args = parse_args()
    program = list(map(int, args.input.read().split(',')))
    
    for noun in range(1000):
        for verb in range(1000):
            try:
                result = run_program(program.copy(), noun, verb)
            except:
                pass
            else:
                if result == 19690720:
                    print(100 * noun + verb)
                    return 0
    return 1


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
