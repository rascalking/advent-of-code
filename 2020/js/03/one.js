#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }


function countTreeCollisions(grid, slope) {
	let count = 0;
	let x = 0;
	for (let y=0; y<grid.length; y++) {
		console.log('(', x, ', ', y, ') = ', grid[y][x]);
		if (grid[y][x] === '#') {
			count++;
		}
		x = (x + slope) % grid[y].length;
	}
	return count;
}


fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let grid = data.split('\n')
		       .filter(item => item !== '')
		       .map(x => x.split(''));
	let numTrees = countTreeCollisions(grid, 3);
	console.log(numTrees);
});
