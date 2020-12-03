#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let numbers = data.split('\n')
			  .filter(item => item !== '')
			  .map(x => parseInt(x, 10));
	for (let i=0; i<numbers.length-2; i++) {
		for (let j=i+1; j<numbers.length-1; j++) {
			for (let k=j+1; k<numbers.length; k++) {
				const x = numbers[i];
				const y = numbers[j];
				const z = numbers[k];
				if (x + y + z == 2020) {
					console.log(x*y*z);
					return 0;
				}
			}
		}
	}
	console.log('Unable to find a pair that add up to 2020');
});
