#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

const numRows = 128;
const numColumns = 8;

class BoardingPass {
	constructor(id) {
		this.id = id;

		let rowPart = id.slice(0, 7);
		this.row = parseInt(rowPart.replaceAll('F', '0')
					   .replaceAll('B', '1'),
				    2);
		let columnPart = id.slice(7, 10);
		this.column = parseInt(columnPart.replaceAll('L', '0')
						 .replaceAll('R', '1'),
		                       2);
		this.seatId = (this.row * 8) + this.column;
	}
}

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let passes = data.split('\n')
			 .filter(x => x)
			 .map(x => new BoardingPass(x));
	console.log(passes);
	let max = Math.max(...passes.map(p => p.seatId));
	console.log(max);
});
