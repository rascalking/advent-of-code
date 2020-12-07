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

let hgtRe = /(?<num>\d+)(?<unit>\w+)/;
let hclRe = /#[\da-f]{6}/;
let eclValidValues = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'];
let pidRe = /\d{9}/;
function validatePassport(passport) {
	for (let field of requiredFields) {
		if (!passport[field]) {
			console.log(`Passport is missing ${field}`, passport);
			return false;
		}
	}

	let byr = parseInt(passport.byr, 10);
	if ((byr < 1920) || (byr > 2002)) {
		console.log(`Passport has invalid byr ${byr}`, passport);
		return false;
	}

	let iyr = parseInt(passport.iyr, 10);
	if ((iyr < 2010) || (iyr > 2020)) {
		console.log(`Passport has invalid iyr ${iyr}`, passport);
		return false;
	}

	let eyr = parseInt(passport.eyr, 10);
	if ((eyr < 2020) || (eyr > 2030)) {
		console.log(`Passport has invalid eyr ${eyr}`, passport);
		return false;
	}

	let hgtMatch = hgtRe.exec(passport.hgt);
	if (!hgtMatch) {
		console.log(`Passport has invalid hgt ${passport.hgt}`, passport);
		return false;
	}
	let height = parseInt(hgtMatch.groups.num, 10);
	if (hgtMatch.groups.unit === 'in') {
		if ((height < 59) || (height > 76)) {
			console.log(`Passport has invalid hgt ${passport.hgt}`, passport);
			return false;
		}
	}
	else if (hgtMatch.groups.unit === 'cm') {
		if ((height < 150) || (height > 193)) {
			console.log(`Passport has invalid hgt ${passport.hgt}`, passport);
			return false;
		}
	}
	else {
		console.log(`Passport has invalid hgt ${passport.hgt}`, passport);
		return false;
	}

	let hclMatch = hclRe.exec(passport.hcl);
	if (!hclMatch) {
		console.log(`Passport has invalid hcl ${passport.hcl}`, passport);
		return false;
	}

	if (!eclValidValues.includes(passport.ecl)) {
		console.log(`Passport has invalid ecl ${passport.ecl}`, passport);
		return false;
	}

	let pidMatch = pidRe.exec(passport.pid);
	if (!pidMatch) {
		console.log(`Passport has invalid pid ${passport.pid}`, passport);
		return false;
	}

	return true;
}

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let passports = data.split('\n\n')
			    .map(x => parsePassport(x));
	valid = passports.filter(validatePassport);
	console.log(valid.length);
	console.log('nominally valid', valid);
});
