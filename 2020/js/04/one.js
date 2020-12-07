#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

let requiredFields = [
	'byr',
	'iyr',
	'eyr',
	'hgt',
	'hcl',
	'ecl',
	'pid',
]

let fieldValueRe = /(?<field>\w+):(?<value>[\w#]+)/;
function parsePassport(raw) {
	let passport = {};
	for (let item of raw.split(/\s+/)) {
		if (!item) continue;

		let match = fieldValueRe.exec(item);
		if (!match) {
			console.log('Unable to match:', item);
			continue;
		}
		passport[match.groups.field] = match.groups.value;
	}
	return passport;
}

function validatePassport(passport) {
	for (let field of requiredFields) {
		if (!passport[field]) {
			//console.log(`Passport is missing ${field}`, passport);
			return false;
		}
	}
	return true;
}

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let passports = data.split('\n\n')
			    .map(x => parsePassport(x));
	let numValid = passports.map(p => validatePassport(p))
				.reduce((accumulator, currentValue) => (accumulator + (currentValue ? 1 : 0)));
	console.log(numValid);
});
