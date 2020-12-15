#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;
    let numbers = data.split('\n').filter(x=>x).map(x=>parseInt(x, 10));
    numbers.sort((a,b)=>a-b);
    let ones=0, threes=1; // device always is 3 above the largest adapter
    let last=0;
    for (let current of numbers) {
        switch(current - last) {
            case 1:
                ones++;
                break;

            case 3:
                threes++;
                break;
        }
        last = current;
    }
    console.log(ones*threes);
});
