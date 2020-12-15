#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }

let shittyMemoize = [];

// can we memoize this and prune the tree?
function countPaths(index, nodes) {
    let count = 0;
    let reachable = [nodes[index]+1, nodes[index]+2, nodes[index]+3];
    for (let intermediate of reachable) {
        let intermediateIndex = nodes.indexOf(intermediate);
        if (intermediateIndex === nodes.length - 1) {
            count += 1;
        }
        else if (intermediateIndex >= 0) {
            let result = shittyMemoize[intermediateIndex];
            if (result === undefined) {
                result = countPaths(intermediateIndex, nodes);
                shittyMemoize[intermediateIndex] = result;
            }
            count += result;
        }
    }
    return count;
}

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;

    let numbers = data.split('\n').filter(x=>x).map(x=>parseInt(x, 10));
    let device = Math.max(...numbers)+3;
    numbers = numbers.concat(0, device);
    numbers.sort((a,b)=>a-b);

    console.log(countPaths(0, numbers));
});
