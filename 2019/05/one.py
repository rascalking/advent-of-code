#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import operator
from enum import IntEnum
from pprint import pprint


class Instruction:
    class Halt(Exception): pass

    class Mode(IntEnum):
        POSITION = 0
        IMMEDIATE = 1

    class Opcode(IntEnum):
        ADD = 1
        MULTIPLY = 2
        INPUT = 3
        OUTPUT = 4
        JUMP_IF_TRUE = 5
        JUMP_IF_FALSE = 6
        LESS_THAN = 7
        EQUALS = 8
        HALT = 99

    MATH_OPERATIONS = {
        Opcode.ADD: operator.add,
        Opcode.MULTIPLY: operator.mul,
    }

    LENGTHS = {
        Opcode.ADD: 4,
        Opcode.MULTIPLY: 4,
        Opcode.INPUT: 2,
        Opcode.OUTPUT: 2,
        Opcode.HALT: 1,
    }

    def __init__(self, memory, pointer):
        self.address = pointer
        self.opcode = Instruction.Opcode(memory[pointer] % 100)
        self.modes = tuple(reversed(tuple(map(lambda x: Instruction.Mode(int(x)),
                                              str(memory[pointer]).zfill(5)[:3]))))
        self.length = self.LENGTHS[self.opcode]

    def execute(self, memory):
        if self.opcode in (Instruction.Opcode.ADD, Instruction.Opcode.MULTIPLY):
            self._execute_math(memory)
        elif self.opcode is Instruction.Opcode.INPUT:
            self._execute_input(memory)
        elif self.opcode is Instruction.Opcode.OUTPUT:
            self._execute_output(memory)
        elif self.opcode is Instruction.Opcode.HALT:
            raise Instruction.Halt()
        else:
            raise NotImplemented()

    def _execute_input(self, memory):
        destination = memory[self.address+1]
        memory[destination] = int(input('input: '))

    def _execute_math(self, memory):
        if self.modes[0] is Instruction.Mode.POSITION:
            address = memory[self.address+1]
            operand1 = memory[address]
        else:
            operand1 = memory[self.address+1]

        if self.modes[1] is Instruction.Mode.POSITION:
            address = memory[self.address+2]
            operand2 = memory[address]
        else:
            operand2 = memory[self.address+2]

        destination = memory[self.address+3]
        operation = self.MATH_OPERATIONS[self.opcode]
        memory[destination] = operation(operand1, operand2)

    def _execute_output(self, memory):
        address = memory[self.address+1]
        print(memory[address])


def run_program(memory):
    pointer = 0
    while pointer < len(memory):
        instruction = Instruction(memory, pointer)
        instruction.execute(memory)
        pointer += instruction.length
    raise RuntimeError('Instruction pointer ran off the end of memory')


def main():
    args = parse_args()
    program = list(map(int, args.input.read().split(',')))
    try:
        run_program(program)
    except Instruction.Halt:
        return 0
    return 1


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
