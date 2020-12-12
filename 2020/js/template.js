#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;

});
