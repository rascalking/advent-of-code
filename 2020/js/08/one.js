#!/usr/bin/env node

const fs = require('fs');


class Program {
    static instructionRe = /^(?<operation>\w+) (?<argument>[+-]{1}\d+)$/;

    constructor(text) {
        this.instructions = [];
        for (let line of text.split('\n').filter(x=>x)) {
            let match = Program.instructionRe.exec(line);
            if (!match) { throw `Unable to parse instruction "{line}"`; }
            this.instructions.push({
                operation: match.groups.operation,
                argument: parseInt(match.groups.argument, 10),
            });
        }
    }

    run() {
        let accumulator = 0;
        let alreadyExecuted = new Set();
        let pc = 0;

        while (true) {
            if (alreadyExecuted.has(pc)) { break; }
            let instruction = this.instructions[pc];
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
                    throw `Unexpected operation {instruction.operation}`;
            }
        }

        return accumulator;
    }
}

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

fs.readFile(myArgs[0], 'utf8', (err, data) => {
    if (err) throw err;
    let program = new Program(data);
    console.log(program.run());
});
