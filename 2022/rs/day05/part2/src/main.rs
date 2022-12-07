use std::env;
use std::fs::File;
use std::io::read_to_string;

use regex::Regex;

#[derive(Debug)]
struct Crane {
    num_stacks: u8,
    stacks: Vec<Vec<char>>,
}

impl Crane {
    fn new(num_stacks: u8) -> Self {
        let mut crane = Self {
            num_stacks,
            stacks: Vec::<Vec<char>>::with_capacity(num_stacks as usize),
        };
        for _ in 0..num_stacks {
            crane.stacks.push(Vec::<char>::new());
        }
        crane
    }

    fn move_crates(&mut self, num: u8, from_stack: u8, to_stack: u8) {
        let mut tmp = Vec::<char>::with_capacity(num as usize);
        for _ in 0..num {
            let item = self.stacks[(from_stack-1) as usize].pop().unwrap();
            tmp.push(item);
        }
        for _ in 0..num {
            let item = tmp.pop().unwrap();
            self.stacks[(to_stack-1) as usize].push(item);
        }
    }

    fn tops(&self) -> String {
        self.stacks.iter().map(|s| s[s.len()-1]).collect()
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let sections = contents.splitn(2, "\n\n").collect::<Vec<&str>>();
    
    // parse the initial crane/crate state
    let mut lines = sections[0].lines().rev();
    let num_stacks = lines
        .next()
        .expect("no line with stack numbers")
        .split_whitespace()
        .last()
        .expect("no last stack number")
        .parse::<u8>()
        .expect("unable to parse last stack number");
    let mut crane = Crane::new(num_stacks);
    for line in lines {
        for n in 0..crane.num_stacks {
            let item = line.chars().nth((n*4+1) as usize).unwrap();
            if item != ' ' {
                crane.stacks[n as usize].push(item);
            }
        }
    }
    //println!("{:?}", crane);

    // do the moves
    let re = Regex::new(r"^move (?P<num>\d+) from (?P<from>\d+) to (?P<to>\d+)$").unwrap();
    for line in sections[1].lines() {
        let caps = re.captures(line).unwrap();
        let num = caps["num"].parse::<u8>().unwrap();
        let from_stack = caps["from"].parse::<u8>().unwrap();
        let to_stack = caps["to"].parse::<u8>().unwrap();
        crane.move_crates(num, from_stack, to_stack);
    }
    //println!("{:?}", crane);
    println!("{}", crane.tops());

	Ok(())
}
