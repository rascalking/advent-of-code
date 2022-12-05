use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io::read_to_string;

//use itertools::Itertools;

fn score(c: char) -> u32 {
    let mut bytes = [0];
    let _ = c.encode_utf8(&mut bytes);
    let score: u8;
    if c.is_ascii_uppercase() {
        // 'A' is 65, and should return 27
        score = bytes[0] - 38;
    } else if c.is_ascii_lowercase() {
        // 'a' is 97, and should return 1
        score = bytes[0] - 96;
    } else {
        score = 0;
    }
    score as u32
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let lines = contents.lines().collect::<Vec<_>>();
    //println!("{:?}", lines);
    let mut total: u32 = 0;
    for group in lines.chunks(3) {
        let first: HashSet<&str> = HashSet::from_iter(group[0].split("").filter(|x| !x.is_empty()));
        let second: HashSet<&str> = HashSet::from_iter(group[1].split("").filter(|x| !x.is_empty()));
        let third: HashSet<&str> = HashSet::from_iter(group[2].split("").filter(|x| !x.is_empty()));
        let mut intersection = &first & &second;
        intersection = &intersection & &third;
        let common = intersection.iter().last().unwrap().chars().last().unwrap();
        let priority = score(common);
        println!(
            "{} (priority {}) is common to {:?} {:?} {:?}",
            common, priority, first, second, third,
        );
        total += priority;
    }
    println!("{}", total);
    Ok(())
}
