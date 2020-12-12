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
    let notFromPreamble = -1;
    for (let i=preambleSize; i<numbers.length; i++) {
        if (!fromPreamble(numbers[i], numbers.slice(i-preambleSize, i))) {
            notFromPreamble = numbers[i];
            break;
        }
    }
    if (notFromPreamble === -1) { throw 'no number not from preamble found'; }
    for (let i=0; i<numbers.length; i++) {
        let accumulator = 0;
        let range = null;
        for (j=i; j<numbers.length; j++) {
            accumulator += numbers[j];
            if (accumulator === notFromPreamble) {
                range = numbers.slice(i, j+1);
                break;
            }
        }
        if (range) {
            console.log(range);
            console.log(Math.min(...range) + Math.max(...range));
            break;
        }
    }
});
