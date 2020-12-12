#!/usr/bin/env node

const fs = require('fs');

const instructionRe = /^(?<operation>\w+) (?<argument>[+-]{1}\d+)$/;
function parseInstructions(text) {
    let instructions = [];
    for (let line of text.split('\n').filter(x=>x)) {
        let match = instructionRe.exec(line);
        if (!match) { throw `Unable to parse instruction "${line}"`; }
        instructions.push({
            operation: match.groups.operation,
            argument: parseInt(match.groups.argument, 10),
        });
    }
    return instructions;
}

class ProgramRepeatError {};
class SegFault {};

class Program {
    constructor(instructions) {
        this.instructions = instructions.slice();
    }

    run() {
        let accumulator = 0;
        let alreadyExecuted = new Set();
        let pc = 0;

        while (true) {
            if (alreadyExecuted.has(pc)) { throw new ProgramRepeatError(); }
            if (pc == this.instructions.length) { break; }

            let instruction = this.instructions[pc];
            if (!instruction) { throw new SegFault(); }
            alreadyExecuted.add(pc);
            switch(instruction.operation) {
                case 'acc':
                    accumulator += instruction.argument;
                    pc += 1;
                    break;

                case 'nop':
                    pc += 1;
                    break;

                case 'jmp':
                    pc += instruction.argument;
                    break;

                default:
                    throw `Unexpected operation ${instruction.operation}`;
            }
        }

        return accumulator;
    }
}

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

fs.readFile(myArgs[0], 'utf8', (err, data) => {
    if (err) throw err;
    let result;

    let instructions = parseInstructions(data);
    for (let i=0; i<instructions.length; i++) {
        let instructions = parseInstructions(data);
        let instruction = instructions[i];
        //console.log(`Patching instruction ${i}: ${instruction.operation} ${instruction.argument}`);
        switch(instruction.operation) {
            case 'acc':
                break;

            case 'jmp':
                instruction.operation = 'nop';
                break;

            case 'nop':
                instruction.operation = 'jmp';
                break;

            default:
                throw `Unexpected operation ${instruction.operation}`;
        }

        let program = new Program(instructions);
        let result;
        try {
            result = program.run();
        }
        catch (e) {
            if (e instanceof ProgramRepeatError) {
                //console.log('\t... caught ProgramRepeatError');
                continue;
            }
            else if (e instanceof SegFault) {
                //console.log('\t... caught SegFault');
                continue;
            }
            else {
                throw e;
            }
        }
        //console.log(program);
        console.log(result);
        break;
    }
});
