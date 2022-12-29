use std::env;
use std::fs::File;
use std::io::read_to_string;
use std::collections::VecDeque;

#[derive(Debug)]
struct Monkey {
    items: VecDeque<i32>,
    operation: String,
    test: i32,
    if_true: usize,
    if_false: usize,
    num_inspected: u32,
}

impl Monkey {
    fn new(input: &str) -> Self {
        let lines: Vec<_> = input.splitn(6, '\n').collect();
        let items: Vec<_> = lines[1].splitn(2, ": ").collect();
        let items: VecDeque<i32> = items[1].split(", ").map(|i| i.parse().unwrap()).collect();
        let operation: Vec<_> = lines[2].splitn(2, "= ").collect();
        let test: Vec<_> = lines[3].split(' ').collect();
        let test: i32 = test[test.len()-1].parse().unwrap();
        let if_true: Vec<_> = lines[4].split(' ').collect();
        let if_true: usize = if_true[if_true.len()-1].parse().unwrap();
        let if_false: Vec<_> = lines[5].trim().split(' ').collect();
        let if_false: usize = if_false[if_false.len()-1].parse().unwrap();
        Self {
            operation: String::from(operation[1]),
            num_inspected: 0,
            items, test, if_true, if_false,
        }
    }

    fn inspect(&mut self, item: i32) -> i32 {
        // i cannot for the fucking life of me figure out how to
        // generate and store a closure after parsing the operation,
        // so we'll just do it every fucking time.
        let v: Vec<_> = self.operation.splitn(3, ' ').collect();
        let op1: i32 = match v[0] {
            "old" => item,
            _ => v[0].parse().unwrap(),
        };
        let op2: i32 = match v[2] {
            "old" => item,
            _ => v[2].parse().unwrap(),
        };
        self.num_inspected += 1;
        match v[1] {
            "+" => op1 + op2,
            "-" => op1 - op2,
            "*" => op1 * op2,
            "/" => op1 / op2,
            _ => panic!("fuck"),
        }
    }
}

fn main() -> std::io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1]).expect("unable to open input file");
    let contents = read_to_string(file).expect("unable to read from input file");
    let mut monkeys: Vec<Monkey> = contents.split("\n\n").map(Monkey::new).collect();
    for _ in 0..20 {
        for m in 0..monkeys.len() {
            while !monkeys[m].items.is_empty() {
                let mut item = monkeys[m].items.pop_front().unwrap();
                item = monkeys[m].inspect(item);
                item = item / 3;
                let dest: usize;
                if (item % monkeys[m].test) == 0 {
                    dest = monkeys[m].if_true;
                } else {
                    dest = monkeys[m].if_false;
                }
                monkeys[dest].items.push_back(item);
            }
        }
    }
    for m in 0..monkeys.len() {
        println!("{}: {:?}", m, (&monkeys[m]).items);
    }
    let mut scores: Vec<u32> = monkeys.iter().map(|m| m.num_inspected).collect();
    scores.sort();
    println!("{}", scores[scores.len()-2] * scores[scores.len()-1]);
    Ok(())
}
