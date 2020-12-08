#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

// why is js so fucking broken?!??! had to steal this from 
//   https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Set#Instance_methods
// since Set doesn't have set methods for whatever stupid reason
function intersection(setA, setB) {
	return new Set([...setA].filter(x => setB.has(x)));
}

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let groups = data.split('\n\n');

	let counts = groups.map(group => {
		let people = group.split('\n').filter(x=>x);
		let yesses = people.map(person => new Set(person.split('')));
		let soSayWeAll = yesses.reduce((acc, curr) => intersection(acc, curr));
		return soSayWeAll.size;
	});
	console.log(counts.reduce((acc, curr) => acc + curr));
});
