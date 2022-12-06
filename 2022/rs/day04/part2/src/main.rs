use std::env;
use std::fs::File;
use std::io::read_to_string;
use std::collections::HashSet;

use regex::Regex;

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let re = Regex::new(r"^(\d+)-(\d+),(\d+)-(\d+)$").unwrap();
    let mut overlap = 0;
    for line in contents.lines() {
        let caps = re.captures(line).unwrap();
        let mut vals = Vec::new();
        for i in 1..=4 {
            vals.push(caps.get(i).unwrap().as_str().parse::<u32>().unwrap());
        }
        let first: HashSet<u32> = HashSet::from_iter((vals[0]..=vals[1]).into_iter());
        let second: HashSet<u32> = HashSet::from_iter((vals[2]..=vals[3]).into_iter());
        if !first.is_disjoint(&second) {
            overlap += 1;
        }
    }
    println!("{}", overlap);
    Ok(())
}
