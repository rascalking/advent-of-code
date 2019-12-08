#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import operator
from enum import IntEnum
from pprint import pprint


class Program:
    def __init__(self):
        self.pointer = 0
        self.memory = []

    def load(self, memory):
        self.memory = memory

    def run(self):
        while self.pointer < len(self.memory):
            instruction = Instruction(self)
            debug(f'>>> {instruction}')
            instruction.execute()
        raise RuntimeError('Instruction pointer ran off the end of memory')

program = Program()


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

    COMPARE_OPERATIONS = {
        Opcode.LESS_THAN: operator.lt,
        Opcode.EQUALS: operator.eq,
    }
    MATH_OPERATIONS = {
        Opcode.ADD: operator.add,
        Opcode.MULTIPLY: operator.mul,
    }
    LENGTHS = {
        Opcode.ADD: 4,
        Opcode.MULTIPLY: 4,
        Opcode.INPUT: 2,
        Opcode.OUTPUT: 2,
        Opcode.JUMP_IF_TRUE: 3,
        Opcode.JUMP_IF_FALSE: 3,
        Opcode.LESS_THAN: 4,
        Opcode.EQUALS: 4,
        Opcode.HALT: 1,
    }

    def __init__(self, program):
        self.program = program
        self.address = program.pointer
        self.opcode = Instruction.Opcode(program.memory[self.address] % 100)
        self.modes = tuple(reversed(tuple(map(lambda x: Instruction.Mode(int(x)),
                                              str(program.memory[self.address]).zfill(5)[:3]))))
        self.length = self.LENGTHS[self.opcode]

    def __str__(self):
        return self.opcode.name

    def execute(self):
        if self.opcode in (Instruction.Opcode.ADD, Instruction.Opcode.MULTIPLY):
            self._execute_math()
        elif self.opcode is Instruction.Opcode.INPUT:
            self._execute_input()
        elif self.opcode is Instruction.Opcode.OUTPUT:
            self._execute_output()
        elif self.opcode is Instruction.Opcode.JUMP_IF_TRUE:
            self._execute_jump_if_true()
        elif self.opcode is Instruction.Opcode.JUMP_IF_FALSE:
            self._execute_jump_if_false()
        elif self.opcode in (Instruction.Opcode.LESS_THAN, Instruction.Opcode.EQUALS):
            self._execute_compare()
        elif self.opcode is Instruction.Opcode.HALT:
            raise Instruction.Halt()
        else:
            raise NotImplemented()

    def _execute_input(self):
        address = self.program.memory[self.program.pointer+1]
        debug(f'>>> ({address})')
        self.program.memory[address] = int(input('input: '))
        self.program.pointer += self.length

    def _execute_jump_if_false(self):
        test = self._read_param(0)
        destination = self._read_param(1)
        debug(f'>>> ({test}, {destination})')
        if test == 0:
            self.program.pointer = destination
        else:
            self.program.pointer += self.length

    def _execute_jump_if_true(self):
        test = self._read_param(0)
        destination = self._read_param(1)
        debug(f'>>> ({test}, {destination})')
        if test != 0:
            self.program.pointer = destination
        else:
            self.program.pointer += self.length

    def _execute_compare(self):
        operand1 = self._read_param(0)
        operand2 = self._read_param(1)
        destination = self._read_writeto_param(2)
        operation = self.COMPARE_OPERATIONS[self.opcode]
        self.program.memory[destination] = operation(operand1, operand2)
        self.program.pointer += self.length

    def _execute_math(self):
        operand1 = self._read_param(0)
        operand2 = self._read_param(1)
        destination = self._read_writeto_param(2)
        debug(f'>>> ({operand1}, {operand2}, {destination})')
        operation = self.MATH_OPERATIONS[self.opcode]
        self.program.memory[destination] = operation(operand1, operand2)
        self.program.pointer += self.length

    def _execute_output(self):
        address = self.program.memory[self.program.pointer+1]
        debug(f'>>> ({address})')
        print(self.program.memory[address])
        self.program.pointer += self.length

    def _read_writeto_param(self, param_num):
        param_address = self.address + 1 + param_num
        return self.program.memory[param_address]

    def _read_param(self, param_num):
        param_address = self.address + 1 + param_num
        if self.modes[param_num] is Instruction.Mode.POSITION:
            address = self.program.memory[param_address]
        else:
            address = param_address
        return self.program.memory[address]


DEBUG = False
def debug(*args, **kwargs):
    if DEBUG:
        print(*args, **kwargs)


def main():
    global DEBUG

    args = parse_args()
    DEBUG = args.debug

    program.load(list(map(int, args.input.read().split(','))))
    try:
        program.run()
    except Instruction.Halt:
        return 0
    return 1


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    parser.add_argument('--debug', action='store_true')
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
