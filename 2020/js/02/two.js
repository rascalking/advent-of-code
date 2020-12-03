#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }


let entryRE = /(\d+)-(\d+) (\w): (\w+)/;
function parseEntry(line) {
	let match = entryRE.exec(line);
	return { 
		lower: parseInt(match[1], 10),
		upper: parseInt(match[2], 10),
		letter: match[3],
		password: match[4],
	}
}


function checkEntry(entry) {
	let a = entry.password[entry.lower-1];
	let b = entry.password[entry.upper-1];
	let count = (a === entry.letter ? 1 : 0) + (b === entry.letter ? 1 : 0);
	return count === 1;
}


fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let entries = data.split('\n')
			  .filter(item => item !== '')
			  .map(x => parseEntry(x));
	let numValid = entries.map(x => checkEntry(x))
			      .map(x => x ? 1 : 0)
			      .reduce((accumulator, currentValue) => accumulator+currentValue);
	console.log(numValid);
});
