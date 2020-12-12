#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

function fromPreamble(number, preamble) {
    for (let i=0; i<preamble.length-1; i++) {
        for (let j=i; j<preamble.length; j++) {
            if (number === preamble[i] + preamble[j]) {
                return true;
            }
        }
    }
    return false;
}

let preambleSize = parseInt(myArgs[1], 10);
fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
    let numbers = data.split('\n').filter(x=>x).map(x=>parseInt(x, 10));
    for (let i=preambleSize; i<numbers.length; i++) {
        if (!fromPreamble(numbers[i], numbers.slice(i-preambleSize, i))) {
            console.log(numbers[i]);
            break;
        }
    }
});
