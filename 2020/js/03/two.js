#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }


function countTreeCollisions(grid, slope) {
	let count = 0;
	let x = 0, y = 0;
	while (y < grid.length) {
		//console.log('(', x, ', ', y, ') = ', grid[y][x]);
		if (grid[y][x] === '#') {
			count++;
		}
		x = (x + slope.x) % grid[y].length;
		y = y + slope.y;
	}
	return count;
}


fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
	let grid = data.split('\n')
		       .filter(item => item !== '')
		       .map(x => x.split(''));

	let slopes = [
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	];
	let numTrees = slopes.map(s => countTreeCollisions(grid, s));
	console.log(numTrees);
	let result = numTrees.reduce((accumulator, currentValue) => (accumulator * currentValue));
	console.log(result);
});
