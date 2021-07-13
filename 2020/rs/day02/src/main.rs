#[macro_use]
extern crate lazy_static;

use std::env;
use std::fs::File;
use std::io::prelude::*;

use regex::Regex;

fn read_lines(filename: &str) -> Result<Vec<String>, std::io::Error> {
    let mut file = File::open(filename)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    let lines: Vec<&str> = contents.split('\n').collect();
    // there has to be a less messy way to do this
    Ok(lines.iter().map(|&line| String::from(line)).collect())
}

fn is_password_valid(entry: &str) -> bool {
    lazy_static! {
        static ref RE: Regex = Regex::new(
            r"(?P<min>\d+)-(?P<max>\d+) (?P<letter>[[:alpha:]]): (?P<password>[[:alpha:]]+)"
        ).unwrap();
    }
    if let Some(caps) = RE.captures(entry) {
        let min: usize = caps.name("min").unwrap().as_str().parse().unwrap();
        let max: usize = caps.name("max").unwrap().as_str().parse().unwrap();
        let letter = caps.name("letter").unwrap().as_str().chars().next().unwrap();
        let password = caps.name("password").unwrap().as_str();
        let count = password.chars().filter(|c| c == &letter).count();
        return min <= count && count <= max;
    }
    false
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let lines = read_lines(&args[1])?;
    let valid = lines.iter().filter(|e| is_password_valid(e)).count();
    println!("{}", valid);
    Ok(())
}
