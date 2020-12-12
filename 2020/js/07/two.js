#!/usr/bin/env node

const fs = require('fs');

let myArgs = process.argv.slice(2);
if (!myArgs[0]) { throw "missing input file"; }
let myBag = myArgs[1];
if (!myBag) { throw "missing my bag"; }

Set.prototype.update = function(elems) {
	for (elem of elems) {
		this.add(elem);
	}
}

class RuleSet {
	constructor(rules) {
		// rules = { kind: { kind: number, ... }, ... } 
		this.rules = rules;
		this.canContain = {}
		this.expandRules();
	}

	expandRules() {
		for (let kind in this.rules) {
			this.canContain[kind] = this._kindCanContain(kind);
		}
	}

	_kindCanContain(kind) {
		let canContain = new Set();
		for (let child in this.rules[kind]) {
			// TODO: avoid lots of duplicated work
			canContain.update(this._kindCanContain(child));
			canContain.add(child);
		}
		return canContain;
	}

  howManyCanContain(kind) {
    let count = 0;
    for (let k in this.canContain) {
      if (this.canContain[k].has(kind)) {
        count++;
      }
    }
    return count;
  }

  totalContainedBags(kind) {
		// this.rules = { kind: { kind: number, ... }, ... } 
    let count = 0;
    for (let child in this.rules[kind]) {
      count += (1 + this.totalContainedBags(child)) * this.rules[kind][child];
    }
    return count;
  }
}

let ruleRe = /^(?<kind>\w+ \w+) bags contain (?<contents>[\w\s,]+)$/;
let contentsRe = /^(?<count>\d+) (?<kind>\w+ \w+) bags?$/;

fs.readFile(myArgs[0], 'utf8', (err, data) => {
	if (err) throw err;

	let rules = {};
	for (let line of data.split('\.\n').filter(x=>x)) {
		let ruleMatch = ruleRe.exec(line);
		if (!ruleMatch) { throw `No rule match for '${line}'`; }
		let kind = ruleMatch.groups.kind;
		let contents = {};
		for (let item of ruleMatch.groups.contents.split(', ').filter(x=>x)) {
			if (item === 'no other bags') { break; }
			let contentMatch = contentsRe.exec(item);
			if (!contentMatch) { throw `No content match for '${item}'`; }
			contents[contentMatch.groups.kind] = parseInt(contentMatch.groups.count, 10);
		}
		rules[kind] = contents;
	}

	let ruleset = new RuleSet(rules);
  console.log(ruleset.totalContainedBags(myBag));
});
